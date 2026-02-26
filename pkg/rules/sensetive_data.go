package rules

import (
	"go/ast"
	"go/token"
	"regexp"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

const (
	variablePattern string = `(?i)(password|secret|token|api_?key|access_?key)`
	stringPattern string = `(?i)(password|secret|token|api_?key|access_?key)[_.\-\s]*[:=]`
)

type SensitiveDataRule struct {
	variablePattern *regexp.Regexp
	stringPattern *regexp.Regexp
}

func NewSensitiveDataRule() *SensitiveDataRule {
	return &SensitiveDataRule{
		variablePattern: regexp.MustCompile(variablePattern),
		stringPattern: regexp.MustCompile(stringPattern),
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