package analyzer_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"github.com/mmmIlia/loglinter/pkg/analyzer"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, analyzer.NewAnalyzer(), "a")
}