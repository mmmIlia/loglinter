package rules

import (
	"go/ast"
	"go/token"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	defaultVariablePattern string = `(?i)(password|secret|token|api_?key|access_?key)`
	defaultStringPattern   string = `(?i)(password|secret|token|api_?key|access_?key)[_.\-\s]*[:=]`
)

type SensitiveDataRule struct {
	variablePattern *regexp.Regexp
	stringPattern   *regexp.Regexp
}

func NewSensitiveDataRule(customPatterns string) *SensitiveDataRule {
	varPattern := defaultVariablePattern
	strPattern := defaultStringPattern

	if customPatterns != "" {
		patterns := strings.Split(customPatterns, ",")
		for i, p := range patterns {
			patterns[i] = strings.TrimSpace(p)
		}

		customRegex := `(?i)(` + strings.Join(patterns, "|") + `)`
		
		varPattern = customRegex
		strPattern = customRegex + `[_.\-\s]*[:=]`
	}

	return &SensitiveDataRule{
		variablePattern: regexp.MustCompile(varPattern),
		stringPattern:   regexp.MustCompile(strPattern),
	}
}

func (r *SensitiveDataRule) Check(pass *analysis.Pass, msgArg ast.Expr) {
	ast.Inspect(msgArg, func(n ast.Node) bool {
		if lit, ok := n.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			val, err := strconv.Unquote(lit.Value)
			if err != nil {
				return true
			}
			
			if r.stringPattern.MatchString(val) {
				pass.Reportf(lit.Pos(), "log message should not contain potential sensitive data")
				return false
			}
		}

		if id, ok := n.(*ast.Ident); ok {
			if r.variablePattern.MatchString(id.Name) {
				pass.Reportf(id.Pos(), "log message should not use variable with potential sensitive data")
				return false
			}
		}

		return true
	})
}