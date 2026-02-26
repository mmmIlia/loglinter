package rules

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

type LowercaseRule struct{}

func NewLowercaseRule() *LowercaseRule {
	return &LowercaseRule{}
}

func (r *LowercaseRule) Check(pass *analysis.Pass, msgArg ast.Expr) {
	lit, ok := msgArg.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return
	}

	val, err := strconv.Unquote(lit.Value)
	if err != nil || len(val) == 0 {
		return
	}

	firstRune, _ := utf8.DecodeRuneInString(val)

	if !unicode.IsLetter(firstRune) {
		return
	}

	if unicode.IsLower(firstRune) {
		return
	}

	runes := []rune(val)
	runes[0] = unicode.ToLower(runes[0])
	newStr := string(runes)

	newLit := strconv.Quote(newStr)

	pass.Report(analysis.Diagnostic{
		Pos:     lit.Pos(),
		End:     lit.End(),
		Message: "log message should start with a lowercase letter",
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: "Convert to lowercase",
				TextEdits: []analysis.TextEdit{
					{
						Pos:     lit.Pos(),
						End:     lit.End(),
						NewText: []byte(newLit),
					},
				},
			},
		},
	})
}