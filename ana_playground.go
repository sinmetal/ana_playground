package ana_playground

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "ana_playground is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "ana_playground",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		fmt.Println(f.Name.String())
		fset := token.NewFileSet()
		expr, err := parser.ParseExprFrom(fset, f.Name.String(), nil, 0)
		if err != nil { /* エラー処理 */
		}
		ast.Print(fset, expr)
	}

	return nil, nil
}
