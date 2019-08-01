package cmd

import (
	"fmt"
	"go/token"
	"io"
	"bytes"
	"os"

	gen "github.com/zemnmez/cardauth/apdu/gen/lib"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/spf13/cobra"
)

// instructionsCmd represents the instructions command
var instructionsCmd = &cobra.Command{
	Use:   "instructions",
	Short: "Generates definitions for APDU instructions",
	RunE:  instructions,
}

func instructions(cmd *cobra.Command, args []string) (err error) {
	var records [][]string
	if records, err = gen.ReadCSV(inputData); err != nil {
		return
	}

	// remove header
	records = records[1:]

	var f *dst.File
	if f, err = gen.Load(outputFile); err != nil {
		return
	}

	var tmp bytes.Buffer

	var out io.Writer = os.Stdout

	if !idempotent {
		var f *os.File
		if f, err = os.OpenFile(outputFile, os.O_TRUNC|os.O_WRONLY, 0700); err != nil {
			return
		}

		defer f.Close()

		out = f
	}


	if f, err = makeInstructions(f, records); err != nil {
		return
	}

	if err = decorator.Fprint(&tmp, f); err != nil {
		return
	}

	if _, err = io.Copy(out, &tmp); err != nil { return }

	return
}

func makeInstructions(f *dst.File, records [][]string) (out *dst.File, err error) {
	out = f

	var constants = dst.GenDecl{
		Tok: token.CONST,
	}

	var infoMap = dst.CompositeLit{
		Type: &dst.MapType{
			Key:   dst.NewIdent(instrTypeName),
			Value: dst.NewIdent("string"),
		},
	}

	var stringifyMap = dst.CompositeLit{
		Type: &dst.MapType{
			Key:   dst.NewIdent(instrTypeName),
			Value: dst.NewIdent("string"),
		},
	}

	var aliasesMap = dst.CompositeLit{
		Type: &dst.MapType{
			Key:   dst.NewIdent(instrTypeName),
			Value: dst.NewIdent(instrTypeName),
		},
	}

	for _, record := range records {
		bytes, name, reference := record[0], record[1], record[2]

		var ns []gen.Number
		if ns, err = gen.ParseValues(bytes); err != nil {
			return
		}

		for i, n := range ns {
			var spec dst.ValueSpec
			myIdent := instrPrefixString + gen.MakeName(name)
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

			spec.Type = dst.NewIdent(instrTypeName)
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

	reversedStringify, err := gen.ReverseMapLit(stringifyMap)
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
					Names:  []*dst.Ident{dst.NewIdent(instrInfoMapName)},
					Values: []dst.Expr{&infoMap},
				},

				&dst.ValueSpec{
					Names:  []*dst.Ident{dst.NewIdent(instrReverseStringMapName)},
					Values: []dst.Expr{&reversedStringify},
				},

				&dst.ValueSpec{
					Names:  []*dst.Ident{dst.NewIdent(instrStringMapName)},
					Values: []dst.Expr{&stringifyMap},
				},

				&dst.ValueSpec{
					Names:  []*dst.Ident{dst.NewIdent(instrSecondariesMapName)},
					Values: []dst.Expr{&aliasesMap},
				},
			},
		},
	)

	return
}

var (
	instrTypeName,
	instrStringMapName,
	instrInfoMapName,
	instrPrefixString,
	instrSecondariesMapName,
	instrReverseStringMapName string
)

func init() {
	rootCmd.AddCommand(instructionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// instructionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	instructionsCmd.Flags().StringVar(&instrStringMapName, "stringMap", "instrStringMap", "name of mapping from instruction to string")
	instructionsCmd.Flags().StringVar(&instrInfoMapName, "infoMap", "instrInfoMap", "name of mapping from instruction to info")
	instructionsCmd.Flags().StringVar(&instrReverseStringMapName, "reverseStringMap", "reverseInstrStringMap", "name of mapping from string to instruction")
	instructionsCmd.Flags().StringVar(&instrPrefixString, "prefix", "Instruction", "prefix added to generated consts")
	instructionsCmd.Flags().StringVar(&instrTypeName, "type", "Instruction", "the type name of an instruction")
	instructionsCmd.Flags().StringVar(&instrSecondariesMapName, "secondariesMap", "instrSecondaries", "the name of the generated mapping from secondaries to primary values for ranges")
}
