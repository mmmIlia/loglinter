package rules

import (
	"unicode"
)

type EnglishRule struct{}

func NewEnglishRule() *EnglishRule {
	return &EnglishRule{}
}

func (r *EnglishRule) Apply(text string) (string, []string) {
	for _, runeValue := range text {
		if !unicode.IsLetter(runeValue) {
			continue
		}

		if runeValue > unicode.MaxASCII {
			return text, []string{"log message must be in English"}
		}
	}
	return text, nil
}