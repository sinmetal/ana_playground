package main

import (
	"github.com/sinmetal/ana_playground"

	"github.com/gostaticanalysis/forcetypeassert"
	"github.com/gostaticanalysis/nofmt"
	"github.com/gostaticanalysis/notest"
	"github.com/gostaticanalysis/vetgen/analyzers"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(append(
		analyzers.Govet(), // go vetと同じもの
		nofmt.Analyzer,
		notest.Analyzer,
		forcetypeassert.Analyzer,
		ana_playground.Analyzer,
	)...)

}
