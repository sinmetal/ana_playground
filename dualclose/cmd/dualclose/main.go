package main

import (
	"github.com/sinmetal/ana_playground/dualclose"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(dualclose.Analyzer) }
