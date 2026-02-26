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

	var firstErrorMsg string
	
	if strings.HasSuffix(val, ".") {
		firstErrorMsg = "log message should not end with punctuation"
	}

	cleanVal := strings.TrimSuffix(val, ".")
	var sb strings.Builder
	
	for _, runeValue := range cleanVal {
		isEmoji := unicode.In(runeValue, unicode.So, unicode.Sk)
		isNoisy := strings.ContainsRune("!?", runeValue)

		if isEmoji || isNoisy {
			if firstErrorMsg == "" {
				if isEmoji {
					firstErrorMsg = "log message should not contain emojis or special symbols"
				} else {
					firstErrorMsg = "log message should not contain exclamation or question marks"
				}
			}
			continue
		}
		sb.WriteRune(runeValue)
	}

	if firstErrorMsg != "" {
		finalStr := sb.String()
		newLit := strconv.Quote(finalStr)

		pass.Report(analysis.Diagnostic{
			Pos:     lit.Pos(),
			End:     lit.End(),
			Message: firstErrorMsg,
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "Clean up log message",
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
}