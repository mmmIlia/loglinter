package rules

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

type SpecialCharsRule struct{}

func NewSpecialCharsRule() *SpecialCharsRule {
	return &SpecialCharsRule{}
}

func (r *SpecialCharsRule) Check(pass *analysis.Pass, msgArg ast.Expr) {
	lit, ok := msgArg.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return
	}

	val, err := strconv.Unquote(lit.Value)
	if err != nil || len(val) == 0 {
		return
	}

	if strings.HasSuffix(val, ".") {
		pass.Reportf(lit.Pos(), "log message should not end with punctuation")
		return
	}

	for _, r := range val {
		if unicode.In(r, unicode.So, unicode.Sk) {
			pass.Reportf(lit.Pos(), "log message should not contain emojis or special symbols")
			return
		}

		if strings.ContainsRune("!?", r) {
			pass.Reportf(lit.Pos(), "log message should not contain exclamation or question marks")
			return
		}
	}
}