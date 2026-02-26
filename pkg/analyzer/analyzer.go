package analyzer

import (
	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "loglinter",
		Doc:  "Checks log messages for specific formatting rules (case, language, symbols, secrets)",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {	
	return nil, nil
}