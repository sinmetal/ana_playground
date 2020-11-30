package dualclose

import (
	"fmt"
	"golang.org/x/tools/go/ssa"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

const doc = "dualclose is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "dualclose",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	m := make(map[ssa.Value]bool)

	s := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	for _, f := range s.SrcFuncs {
		fmt.Println(f)
		for _, b := range f.Blocks {
			fmt.Printf("\tBlock %d\n", b.Index)
			for _, instr := range b.Instrs {
				fmt.Printf("\t\t%[1]T\t%[1]v(%[1]p)\n", instr)
				x, _ := instr.(*ssa.Call)
				if x == nil {
					continue
				}

				y, _ := x.Common().Value.(*ssa.Builtin)
				if y == nil {
					continue
				}
				fmt.Printf("Builtin func name = %s\n", y.Name())
				if y.Name() != "close" || len(x.Common().Args) != 1 {
					continue
				}
				arg0 := x.Common().Args[0]
				if m[arg0] {
					pass.Reportf(x.Pos(), "close dual chan")
					continue
				}
				m[arg0] = true

				//for _, v := range instr.Operands(nil) {
				//	if v != nil {
				//		fmt.Printf("\t\t\t%[1]T\t%[1]v(%[1]p)\n", *v)
				//	}
				//}
			}
		}
	}
	return nil, nil
}
