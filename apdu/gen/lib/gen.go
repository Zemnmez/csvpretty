// Package gen exposes utility functions for code generation
package gen

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/dave/dst"
	decorator "github.com/dave/dst/decorator"
)

// Load loads a file for overwriting. All entities not annotated
// with a go:generate comment are stripped
func Load(name string) (f *dst.File, err error) {
	var inFile []byte
	if inFile, err = ioutil.ReadFile(name); err != nil {
		return
	}

	if f, err = decorator.Parse(inFile); err != nil {
		return
	}

	var newDecls []dst.Decl
	for _, Decl := range f.Decls {
		comments := Decl.Decorations().Start.All()
		if len(comments) < 1 {
			continue
		}

		if !strings.Contains(comments[0], "//go:generate") {
			continue
		}

		newDecls = append(newDecls, Decl)
	}

	f.Decls = newDecls

	return
}

// ReadCSV reads a CSV file by name, trimming any leading space from
// records.
func ReadCSV(name string) (records [][]string, err error) {

	var csvFile []byte
	if csvFile, err = ioutil.ReadFile(name); err != nil {
		return
	}

	csvReader := csv.NewReader(bytes.NewReader(csvFile))
	csvReader.TrimLeadingSpace = true
	if records, err = csvReader.ReadAll(); err != nil {
		return
	}

	return

}

// ReverseMapLit returns a dst.CompositeLit which is the exact reverse of its
// input.
func ReverseMapLit(lit dst.CompositeLit) (reversed dst.CompositeLit, err error) {
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

// MakeName takes a fairly arbitrary string and returns a transformed one suitable
// for a Go identifier. MakeName strips all non letter characters and spaces.
// Words are capitalized in CamelCase.
func MakeName(s string) string {
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

// Number represents an `int` with a specific string representation
// (i.e. Hex or Decimal).
type Number interface {
	Int() int
	String() string
	Add(i int) Number
}

// Hex represents an `int` that should be represented as HexaDecimal.
// It imlements the `Number` interface.
type Hex int

func (h Hex) String() string   { return fmt.Sprintf("%#x", int(h)) }
func (h Hex) Int() int         { return int(h) }
func (h Hex) Add(i int) Number { return h + Hex(i) }

// Dec represents an `int` that should be represented as Decimal.
// It implements the `Number` interface.
type Dec int

func (d Dec) String() string   { return fmt.Sprintf("%d", int(d)) }
func (d Dec) Int() int         { return int(d) }
func (d Dec) Add(i int) Number { return d + Dec(i) }

// ParseInt parses a Hex or Decimal integer from a string, returning
// a `Number` which is an instance of `Hex` or `Dec` correspondingly.
func ParseInt(i string) (n Number, err error) {
	var parsed int64
	if parsed, err = strconv.ParseInt(i, 0, 64); err != nil {
		return
	}

	switch {
	case strings.HasPrefix(i, "0x"):
		n = Hex(parsed)
	default:
		n = Dec(parsed)
	}

	return
}

// ParseValues parses a set of values in a standard syntax. This form
// represents a list and / or range of integers.
//
// Semicolon (";") denotes a list of Numbers, for example 1; 2; 3
// represents the Numbers 1, 2 and 3 as `Dec` values.
//
// Double dot ("..") denotes a range of Numbers, for example 0x00..0x01
// represents the Numbers 0 and 1, to be rendered in HexaDecimal.
//
// These forms can be combined. For example 1..3; 4..5 represents the sequence
// 1, 2, 3, 4, 5.
func ParseValues(r string) (ns []Number, err error) {
	for _, s := range strings.Split(r, ";") {
		var nums []Number
		if nums, err = ParseRange(strings.TrimSpace(s)); err != nil {
			return
		}

		ns = append(ns, nums...)
	}

	return
}

// ParseRange parses a range of values in a standard syntax. See `ParseValues`
// double dot syntax for more information.
func ParseRange(r string) (ns []Number, err error) {
	splits := strings.Split(r, "..")

	ns = make([]Number, len(splits))

	for i, v := range splits {
		if ns[i], err = ParseInt(strings.TrimSpace(v)); err != nil {
			return
		}
	}

	if len(splits) < 2 {
		return
	}

	if len(splits) != 2 {
		err = errors.New("expecting only one '..' in range Decl")
		return
	}

	low, high := ns[0], ns[1]

	if low.Int() > high.Int() {
		err = fmt.Errorf("range invalid: %s should be before %s", low, high)
	}

	ns = make([]Number, high.Int()-low.Int())
	for i := 0; (i + low.Int()) <= high.Int(); i++ {
		ns[i] = low.Add(i)
	}

	return
}
