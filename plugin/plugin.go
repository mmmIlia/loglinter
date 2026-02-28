package plugin

import (
	"fmt"

	"golang.org/x/tools/go/analysis"

	"github.com/mmmIlia/loglinter/pkg/analyzer"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	a := analyzer.NewAnalyzer()

	if conf == nil {
		return []*analysis.Analyzer{a}, nil
	}

	settings, ok := conf.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("loglinter: configuration must be a map, got %T", conf)
	}

	for key, value := range settings {
		strValue := fmt.Sprintf("%v", value)

		if err := a.Flags.Set(key, strValue); err != nil {
			return nil, fmt.Errorf("loglinter: invalid setting %q: %w", key, err)
		}
	}

	return []*analysis.Analyzer{a}, nil
}
