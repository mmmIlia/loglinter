package plugin

import (
	"golang.org/x/tools/go/analysis"

	"github.com/mmmIlia/loglinter/pkg/analyzer"
)

const Name = "loglinter"

func New(conf any) ([]*analysis.Analyzer, error) {
	a := analyzer.NewAnalyzer()

	if settings, ok := conf.(map[string]any); ok {
		if val, ok := settings["disable-lowercase"].(bool); ok && val {
			a.Flags.Set("disable-lowercase", "true")
		}
		if val, ok := settings["disable-english"].(bool); ok && val {
			a.Flags.Set("disable-english", "true")
		}
		if val, ok := settings["disable-special-chars"].(bool); ok && val {
			a.Flags.Set("disable-special-chars", "true")
		}
		if val, ok := settings["disable-sensitive"].(bool); ok && val {
			a.Flags.Set("disable-sensitive", "true")
		}

		if val, ok := settings["sensitive-patterns"].(string); ok {
			a.Flags.Set("sensitive-patterns", val)
		}
	}

	return []*analysis.Analyzer{a}, nil
}
