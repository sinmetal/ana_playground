package main

import (
	"github.com/sinmetal/ana_playground"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(ana_playground.Analyzer) }
