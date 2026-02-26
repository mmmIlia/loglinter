package rules

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

type EnglishRule struct{}

func NewEnglishRule() *EnglishRule {
	return &EnglishRule{}
}

func (r *EnglishRule) Check(pass *analysis.Pass, msgArg ast.Expr) {
	lit, ok := msgArg.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return
	}

	val, err := strconv.Unquote(lit.Value)
	if err != nil || len(val) == 0 {
		return
	}

	for _, runeValue := range val {
		if !unicode.IsLetter(runeValue) {
			continue
		}

		if runeValue > unicode.MaxASCII {
			pass.Reportf(lit.Pos(), "log message must be in English")
			return
		}
	}
}