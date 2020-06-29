package ana_playground

import (
	"fmt"
	"go/ast"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
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
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	if err := duimport(pass); err != nil {
		return nil, err
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			if n.Name == "gopher" {
				pass.Reportf(n.Pos(), "identifyer is gopher")
			}
		}
	})

	return nil, nil
}

type DuplicateValue struct {
	Spec         *ast.ImportSpec
	AlreadyError bool
}

func duimport(pass *analysis.Pass) error {
	m := make(map[string]*DuplicateValue)
	for _, f := range pass.Files {
		for _, imp := range f.Imports {
			path, err := strconv.Unquote(imp.Path.Value)
			if err != nil {
				return err
			}
			dv, ok := m[path]
			if ok {
				if !dv.AlreadyError {
					ip, err := buildImportPath(dv.Spec)
					if err != nil {
						return err
					}
					pass.Reportf(dv.Spec.Pos(), "%s is duplicated import", ip)
					dv.AlreadyError = true
				}

				ip, err := buildImportPath(imp)
				if err != nil {
					return err
				}
				pass.Reportf(imp.Pos(), "%s is duplicated import", ip)
			} else {
				m[path] = &DuplicateValue{
					Spec: imp,
				}
			}
		}
	}

	return nil
}

func buildImportPath(imp *ast.ImportSpec) (string, error) {
	path, err := strconv.Unquote(imp.Path.Value)
	if err != nil {
		return "", err
	}
	if imp.Name == nil {
		return path, nil
	}

	return fmt.Sprintf("%s %s", imp.Name.Name, path), nil
}
