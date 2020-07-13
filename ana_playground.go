package ana_playground

import (
	"go/ast"
	"strconv"

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
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			switch v := decl.(type) {
			case *ast.FuncDecl:
				for _, body := range v.Body.List {
					switch b := body.(type) {
					case *ast.ReturnStmt:
						for _, result := range b.Results {
							switch r := result.(type) {
							case *ast.BasicLit:
								rawValue, err := strconv.Unquote(r.Value)
								if err != nil {
									panic(err)
								}
								if len(rawValue) < 1 {
									pass.Reportf(r.Pos(), "空文字返してるところ")
								}
							}
						}
					}
				}
			default:
			}
		}
	}

	return nil, nil
}
