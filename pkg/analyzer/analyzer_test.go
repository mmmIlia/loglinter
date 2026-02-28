package analyzer_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/mmmIlia/loglinter/pkg/analyzer"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	a := analyzer.NewAnalyzer()
	analysistest.RunWithSuggestedFixes(t, testdata, a, "a")
}

func TestAnalyzer_ConfigDisabled(t *testing.T) {
	testdata := analysistest.TestData()

	a := analyzer.NewAnalyzer()

	if err := a.Flags.Set("disable-lowercase", "true"); err != nil {
		t.Fatal(err)
	}
	if err := a.Flags.Set("disable-english", "true"); err != nil {
		t.Fatal(err)
	}
	if err := a.Flags.Set("disable-special-chars", "true"); err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, a, "config_test")
}

func TestAnalyzer_CustomPatterns(t *testing.T) {
	testdata := analysistest.TestData()

	a := analyzer.NewAnalyzer()

	if err := a.Flags.Set("sensitive-patterns", "email,ssn"); err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, a, "custom_patterns")
}
