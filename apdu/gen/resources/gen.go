package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

var (
	instructions       string
	instructionType    string
	instructionStrings string
	target             string
	idempotent         bool
)

func init() {
	flag.StringVar(&target, "target", "", "file to add to")
	flag.StringVar(&instructions, "ins.file", "", "instructions.csv")
	flag.StringVar(&instructionType, "ins.type", "Instruction", "the type of the generated instructions block")
	flag.BoolVar(&idempotent, "x", false, "print instead of overwriting")
}

func main() {
	if err := do(); err != nil {
		panic(err)
	}
}

func do() (err error) {
	flag.Parse()
	if target == "" {
		return errors.New("specify target")
	}

	var inFile []byte
	if inFile, err = ioutil.ReadFile(target); err != nil {
		return
	}

	var f *dst.File
	if f, err = decorator.Parse(inFile); err != nil {
		return
	}

	var newDecls []dst.Decl
	for _, decl := range f.Decls {
		comments := decl.Decorations().Start.All()
		if len(comments) < 1 {
			continue
		}

		if !strings.Contains(comments[0], "//go:generate") {
			continue
		}

		newDecls = append(newDecls, decl)
	}

	f.Decls = newDecls

	if instructions != "" {
		var csvFile []byte
		if csvFile, err = ioutil.ReadFile(instructions); err != nil {
			return
		}

		var records [][]string
		csvReader := csv.NewReader(bytes.NewReader(csvFile))
		csvReader.TrimLeadingSpace = true
		if records, err = csvReader.ReadAll(); err != nil {
			return
		}

		if err = instructionsSwitch(f, records[1:]); err != nil {
			return
		}
	}

	var b bytes.Buffer
	if err = decorator.Fprint(&b, f); err != nil {
		return
	}

	var out io.Writer = os.Stdout

	if !idempotent {
		out, err = os.OpenFile(target, os.O_TRUNC|os.O_WRONLY, 0700)
	}

	if _, err = io.Copy(out, &b); err != nil {
		return
	}

	return
}

func reverseMapLit(lit dst.CompositeLit) (reversed dst.CompositeLit, err error) {
	reversed.Type = &dst.MapType{
		Key:   dst.Clone(lit.Type.(*dst.MapType).Value).(dst.Expr),
		Value: dst.Clone(lit.Type.(*dst.MapType).Key).(dst.Expr),
	}

	reversed.Elts = make([]dst.Expr, len(lit.Elts))
	for i, f := range lit.Elts {
		var kv *dst.KeyValueExpr
		var ok bool
		if kv, ok = f.(*dst.KeyValueExpr); !ok {
			err = fmt.Errorf("reverseMapLit: %s not dst.KeyValueExpr", reflect.TypeOf(f))
		}

		reversed.Elts[i] = &dst.KeyValueExpr{
			Key:   dst.Clone(kv.Value).(dst.Expr),
			Value: dst.Clone(kv.Key).(dst.Expr),
		}
	}

	return
}

func instructionsSwitch(f *dst.File, records [][]string) (err error) {

	var (
		stringMapIdent        = dst.NewIdent("instrStringMap")
		infoMapIdent          = dst.NewIdent("instrInfoMap")
		reverseStringMapIdent = dst.NewIdent("reverseInstrStringMap")
		secondariesIdent      = dst.NewIdent("instrSecondariesMap")
	)

	var constants = dst.GenDecl{
		Tok: token.CONST,
	}

	var infoMap = dst.CompositeLit{
		Type: &dst.MapType{
			Key:   dst.NewIdent("Instruction"),
			Value: dst.NewIdent("string"),
		},
	}

	var stringifyMap = dst.CompositeLit{
		Type: &dst.MapType{
			Key:   dst.NewIdent("Instruction"),
			Value: dst.NewIdent("string"),
		},
	}

	var aliasesMap = dst.CompositeLit{
		Type: &dst.MapType{
			Key:   dst.NewIdent("Instruction"),
			Value: dst.NewIdent("Instruction"),
		},
	}

	for _, record := range records {
		bytes, name, reference := record[0], record[1], record[2]

		var ns []number
		if ns, err = parseValues(bytes); err != nil {
			return
		}

		for i, n := range ns {
			var spec dst.ValueSpec
			myIdent := "Instruction" + makeName(name)
			fullRef := fmt.Sprintf("ISO/IEC 7816-4:2005(E) %s", reference)
			myComment := fmt.Sprintf("// '%s': see %s", name, fullRef)
			var myValue dst.Expr = &dst.BasicLit{Kind: token.INT, Value: n.String()}
			var isAlias = i > 0

			if isAlias {
				newIdent := myIdent + fmt.Sprintf("%d", i+1)
				aliasesMap.Elts = append(aliasesMap.Elts, &dst.KeyValueExpr{
					Key:   dst.NewIdent(newIdent),
					Value: dst.NewIdent(myIdent),
				})

				myComment = "// alias"
				//myValue = dst.NewIdent(myIdent)

				myIdent = newIdent

			}

			spec.Type = dst.NewIdent("Instruction")
			spec.Names = []*dst.Ident{dst.NewIdent(myIdent)}
			spec.Values = []dst.Expr{myValue}
			spec.Decorations().End.Append(myComment)

			infoMap.Elts = append(infoMap.Elts, &dst.KeyValueExpr{
				Key:   dst.NewIdent(myIdent),
				Value: &dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%+q", fmt.Sprintf("'%s': see %s", name, fullRef))},
			})

			stringifyMap.Elts = append(stringifyMap.Elts, &dst.KeyValueExpr{
				Key:   dst.NewIdent(myIdent),
				Value: &dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%+q", myIdent)},
			})

			constants.Specs = append(constants.Specs, &spec)
		}
	}

	constants.Specs = append(constants.Specs, &dst.ValueSpec{
		Names:  []*dst.Ident{dst.NewIdent("_")},
		Values: []dst.Expr{&dst.BasicLit{Kind: token.INT, Value: "0"}},
	})

	reversedStringify, err := reverseMapLit(stringifyMap)
	if err != nil {
		return
	}

	f.Decls = append(
		f.Decls,
		&constants,
		&dst.GenDecl{
			Tok: token.VAR,
			Specs: []dst.Spec{
				&dst.ValueSpec{
					Names:  []*dst.Ident{infoMapIdent},
					Values: []dst.Expr{&infoMap},
				},

				&dst.ValueSpec{
					Names:  []*dst.Ident{reverseStringMapIdent},
					Values: []dst.Expr{&reversedStringify},
				},

				&dst.ValueSpec{
					Names:  []*dst.Ident{stringMapIdent},
					Values: []dst.Expr{&stringifyMap},
				},

				&dst.ValueSpec{
					Names:  []*dst.Ident{secondariesIdent},
					Values: []dst.Expr{&aliasesMap},
				},
			},
		},
	)

	return
}

func makeName(s string) string {
	var letters = make([]rune, 0, len(s))
	var nextTitle bool
	for i, r := range s {
		if unicode.IsSpace(r) {
			nextTitle = true
		}

		switch {
		case !unicode.IsLetter(r) && !unicode.IsNumber(r):
			continue
		}

		switch {
		case i == 0 || nextTitle:
			r = unicode.ToTitle(r)
			nextTitle = false
		default:
			r = unicode.ToLower(r)
		}

		letters = append(letters, r)
	}

	return string(letters)
}

type number interface {
	Int() int
	String() string
	Add(i int) number
}

type hex int

func (h hex) String() string   { return fmt.Sprintf("%#x", int(h)) }
func (h hex) Int() int         { return int(h) }
func (h hex) Add(i int) number { return h + hex(i) }

type dec int

func (d dec) String() string   { return fmt.Sprintf("%d", int(d)) }
func (d dec) Int() int         { return int(d) }
func (d dec) Add(i int) number { return d + dec(i) }

func parseInt(i string) (n number, err error) {
	var parsed int64
	if parsed, err = strconv.ParseInt(i, 0, 64); err != nil {
		return
	}

	switch {
	case strings.HasPrefix(i, "0x"):
		n = hex(parsed)
	default:
		n = dec(parsed)
	}

	return
}

func parseValues(r string) (ns []number, err error) {
	for _, s := range strings.Split(r, ";") {
		var nums []number
		if nums, err = parseRange(strings.TrimSpace(s)); err != nil {
			return
		}

		ns = append(ns, nums...)
	}

	return
}

func parseRange(r string) (ns []number, err error) {
	splits := strings.Split(r, "..")

	ns = make([]number, len(splits))

	for i, v := range splits {
		if ns[i], err = parseInt(strings.TrimSpace(v)); err != nil {
			return
		}
	}

	if len(splits) < 2 {
		return
	}

	if len(splits) != 2 {
		err = errors.New("expecting only one '..' in range decl")
		return
	}

	low, high := ns[0], ns[1]

	if low.Int() > high.Int() {
		err = fmt.Errorf("range invalid: %s should be before %s", low, high)
	}

	ns = make([]number, high.Int()-low.Int())
	for i := 0; (i + low.Int()) <= high.Int(); i++ {
		ns[i] = low.Add(i)
	}

	return
}
