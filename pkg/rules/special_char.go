package rules

import (
	"strings"
	"unicode"
)

type SpecialCharsRule struct{}

func NewSpecialCharsRule() *SpecialCharsRule {
	return &SpecialCharsRule{}
}

func (r *SpecialCharsRule) Apply(text string) (string, []string) {
	var violations []string

	hasBadSuffix := strings.HasSuffix(text, ".")
	if hasBadSuffix {
		violations = append(violations, "log message should not end with punctuation")
	}

	cleanText := strings.TrimRight(text, ".")

	var sb strings.Builder
	var hasEmoji, hasNoise bool

	for _, runeValue := range cleanText {
		isEmoji := unicode.In(runeValue, unicode.So, unicode.Sk)
		isNoisy := strings.ContainsRune("!?", runeValue)

		if isEmoji || isNoisy {
			if isEmoji && !hasEmoji {
				violations = append(violations, "log message should not contain emojis")
				hasEmoji = true
			}
			if isNoisy && !hasNoise {
				violations = append(violations, "log message should not contain exclamation or question marks")
				hasNoise = true
			}
			continue
		}
		sb.WriteRune(runeValue)
	}

	return sb.String(), violations
}
