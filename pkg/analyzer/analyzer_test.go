package analyzer_test

import (
	"testing"

	"github.com/mmmIlia/loglinter/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, analyzer.NewAnalyzer(), "a")
}

func TestAnalyzer_ConfigDisabled(t *testing.T) {
	testdata := analysistest.TestData()

	a := analyzer.NewAnalyzer()

	a.Flags.Set("disable-lowercase", "true")
	a.Flags.Set("disable-english", "true")
	a.Flags.Set("disable-special-chars", "true")

	analysistest.Run(t, testdata, a, "config_test")
}

func TestAnalyzer_CustomPatterns(t *testing.T) {
	testdata := analysistest.TestData()

	a := analyzer.NewAnalyzer()
	a.Flags.Set("sensitive-patterns", "email,ssn")

	analysistest.Run(t, testdata, a, "custom_patterns")
}
