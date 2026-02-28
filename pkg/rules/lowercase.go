package rules

import (
	"unicode"
	"unicode/utf8"
)

type LowercaseRule struct{}

func NewLowercaseRule() *LowercaseRule {
	return &LowercaseRule{}
}

func (r *LowercaseRule) Apply(text string) (string, []string) {
	if text == "" {
		return text, nil
	}

	firstRune, _ := utf8.DecodeRuneInString(text)

	if !unicode.IsLetter(firstRune) || unicode.IsLower(firstRune) {
		return text, nil
	}

	runes := []rune(text)
	runes[0] = unicode.ToLower(runes[0])

	return string(runes), []string{"log message should start with a lowercase letter"}
}
