package cmd

import (
	"fmt"
	"go/token"
	"io"
	"bytes"
	gen "github.com/zemnmez/cardauth/apdu/gen/lib"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/spf13/cobra"
)

// instructionsCmd represents the instructions command
var instructionsCmd = &cobra.Command{
	Use:   "constants",
	Short: "Generates constants and associated metadata from a csv",
	RunE:  instructions,
}

func instructions(cmd *cobra.Command, args []string) (err error) {
	// remove header
	records := inputCSV.Records[1:]

	var tmp bytes.Buffer

	var f *dst.File
	if f, err = makeInstructions(inputGo.File, records); err != nil {
		return
	}

	if err = decorator.Fprint(&tmp, f); err != nil {
		return
	}

	if _, err = io.Copy(outputFile, &tmp); err != nil { return }

	return
}

type infoMetadata struct {
	OriginalName string
	Identifier string
	Info string
}

func makeInstructions(f *dst.File, records [][]string) (out *dst.File, err error) {
	out = f

	var constants = dst.GenDecl{
		Tok: token.CONST,
	}

	var infoMap = dst.CompositeLit{
		Type: &dst.MapType{
			Key:   dst.NewIdent(typeName),
			Value: dst.NewIdent("string"),
		},
	}

	var stringifyMap = dst.CompositeLit{
		Type: &dst.MapType{
			Key:   dst.NewIdent(typeName),
			Value: dst.NewIdent("string"),
		},
	}

	var aliasesMap = dst.CompositeLit{
		Type: &dst.MapType{
			Key:   dst.NewIdent(typeName),
			Value: dst.NewIdent(typeName),
		},
	}

	for _, record := range records {
		byteSq, name, reference := record[0], record[1], record[2]

		var ns []gen.Number
		if ns, err = gen.ParseValues(byteSq); err != nil {
			return
		}

		for i, n := range ns {
			var spec dst.ValueSpec
			myIdent := prefixString + gen.MakeName(name)
			var info bytes.Buffer

			if err = commentTemplate.Execute(&info, infoMetadata {
				OriginalName: name,
				Identifier: myIdent,
				Info: reference,
			}); err != nil { return }

			myComment := fmt.Sprintf("// %s", info.String())

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

			spec.Type = dst.NewIdent(typeName)
			spec.Names = []*dst.Ident{dst.NewIdent(myIdent)}
			spec.Values = []dst.Expr{myValue}
			spec.Decorations().End.Append(myComment)

			infoMap.Elts = append(infoMap.Elts, &dst.KeyValueExpr{
				Key:   dst.NewIdent(myIdent),
				Value: &dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%+q", info.String())},
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
					Names:  []*dst.Ident{dst.NewIdent(infoMapName)},
					Values: []dst.Expr{&infoMap},
				},

				&dst.ValueSpec{
					Names:  []*dst.Ident{dst.NewIdent(reverseStringMapName)},
					Values: []dst.Expr{&reversedStringify},
				},

				&dst.ValueSpec{
					Names:  []*dst.Ident{dst.NewIdent(stringMapName)},
					Values: []dst.Expr{&stringifyMap},
				},

				&dst.ValueSpec{
					Names:  []*dst.Ident{dst.NewIdent(secondariesMapName)},
					Values: []dst.Expr{&aliasesMap},
				},
			},
		},
	)

	return
}

var (
	typeName,
	stringMapName,
	infoMapName,
	prefixString,
	secondariesMapName,
	reverseStringMapName string

	commentTemplate gen.Template
)


func init() {
	rootCmd.AddCommand(instructionsCmd)

	instructionsCmd.Flags().Var(&commentTemplate, "comment", "template for generating the comment; see infoMetadata for available values")
	instructionsCmd.MarkFlagRequired("comment")

	instructionsCmd.Flags().StringVar(&stringMapName, "stringMap", "", "name of mapping from instruction to string")
	instructionsCmd.MarkFlagRequired("stringMap")

	instructionsCmd.Flags().StringVar(&infoMapName, "infoMap", "", "name of mapping from instruction to info")
	instructionsCmd.MarkFlagRequired("infoMap")

	instructionsCmd.Flags().StringVar(&reverseStringMapName, "reverseStringMap", "", "name of mapping from string to instruction")
	instructionsCmd.MarkFlagRequired("reverseStringMap")

	instructionsCmd.Flags().StringVar(&prefixString, "prefix", "", "prefix added to generated consts")
	instructionsCmd.MarkFlagRequired("prefix")

	instructionsCmd.Flags().StringVar(&typeName, "type", "", "the type name of an instruction")
	instructionsCmd.MarkFlagRequired("type")

	instructionsCmd.Flags().StringVar(&secondariesMapName, "secondariesMap", "", "the name of the generated mapping from secondaries to primary values for ranges")
	instructionsCmd.MarkFlagRequired("secondariesMapName")

}
