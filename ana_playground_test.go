package ana_playground_test

import (
	"testing"

	"github.com/sinmetal/ana_playground"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, ana_playground.Analyzer, "a")
}
