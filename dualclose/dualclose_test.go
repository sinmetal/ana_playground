package dualclose_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/sinmetal/ana_playground/dualclose"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, dualclose.Analyzer, "a")
}
