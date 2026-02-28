package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/mmmIlia/loglinter/pkg/rules"
)

type Config struct {
	DisableLowercase    bool
	DisableEnglish      bool
	DisableSpecialChars bool
	DisableSensitive    bool
	SensitivePatterns   string
}

var cfg Config

func NewAnalyzer() *analysis.Analyzer {
	a := &analysis.Analyzer{
		Name: "loglinter",
		Doc: `Checks for common issues in log messages.

This linter analyzes calls to slog and zap to enforce several conventions:
1. Log messages must start with a lowercase letter.
2. Log messages must be in English.
3. Log messages should not contain emojis or noisy punctuation.
4. Log messages should not contain sensitive data like passwords or tokens.`,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	a.Flags.BoolVar(&cfg.DisableLowercase, "disable-lowercase", false, "Disable checking for lowercase start")
	a.Flags.BoolVar(&cfg.DisableEnglish, "disable-english", false, "Disable checking for English language")
	a.Flags.BoolVar(&cfg.DisableSpecialChars, "disable-special-chars", false, "Disable checking for emojis and noisy punctuation")
	a.Flags.BoolVar(&cfg.DisableSensitive, "disable-sensitive", false, "Disable checking for sensitive data")
	a.Flags.StringVar(&cfg.SensitivePatterns, "sensitive-patterns", "", "Comma-separated list of custom regex patterns for sensitive data")

	a.Run = func(pass *analysis.Pass) (any, error) {
		var textRules []rules.TextRule
		if !cfg.DisableLowercase {
			textRules = append(textRules, rules.NewLowercaseRule())
		}
		if !cfg.DisableEnglish {
			textRules = append(textRules, rules.NewEnglishRule())
		}
		if !cfg.DisableSpecialChars {
			textRules = append(textRules, rules.NewSpecialCharsRule())
		}

		var dataRules []rules.NodeRule
		if !cfg.DisableSensitive {
			dataRules = append(dataRules, rules.NewSensitiveDataRule(cfg.SensitivePatterns))
		}

		return runLogic(pass, textRules, dataRules)
	}

	return a
}

var targetLoggers = map[string]map[string]int{
	"log/slog": {
		"Info":  0,
		"Error": 0,
		"Warn":  0,
		"Debug": 0,
	},
	"go.uber.org/zap": {
		"Info":  0,
		"Error": 0,
		"Warn":  0,
		"Debug": 0,
		"Fatal": 0,
	},
}

func runLogic(pass *analysis.Pass, textRules []rules.TextRule, dataRules []rules.NodeRule) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}

	insp.Preorder(nodeFilter, func(node ast.Node) {
		processLogCall(pass, node, textRules, dataRules)
	})

	return nil, nil
}

func processLogCall(pass *analysis.Pass, node ast.Node, textRules []rules.TextRule, dataRules []rules.NodeRule) {
	call, ok := node.(*ast.CallExpr)
	if !ok {
		return
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	obj := pass.TypesInfo.Uses[sel.Sel]
	if obj == nil {
		return
	}

	funcObj, ok := obj.(*types.Func)
	if !ok {
		return
	}

	pkg := funcObj.Pkg()
	if pkg == nil {
		return
	}

	methods, isTargetPkg := targetLoggers[pkg.Path()]
	if !isTargetPkg {
		return
	}

	msgIndex, isTargetMethod := methods[funcObj.Name()]
	if !isTargetMethod {
		return
	}

	if len(call.Args) <= msgIndex {
		return
	}

	msgArg := call.Args[msgIndex]

	if lit, ok := msgArg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
		originalVal, err := strconv.Unquote(lit.Value)
		if err == nil {
			currentText := originalVal
			var allViolations []string

			for _, rule := range textRules {
				newText, violations := rule.Apply(currentText)
				if len(violations) > 0 {
					allViolations = append(allViolations, violations...)
				}
				currentText = newText
			}

			if len(allViolations) > 0 {
				var fixes []analysis.SuggestedFix

				if currentText != originalVal {
					newLit := strconv.Quote(currentText)
					fixes = []analysis.SuggestedFix{{
						Message: "Apply all formatting fixes",
						TextEdits: []analysis.TextEdit{{
							Pos:     lit.Pos(),
							End:     lit.End(),
							NewText: []byte(newLit),
						}},
					}}
				}

				for _, violation := range allViolations {
					pass.Report(analysis.Diagnostic{
						Pos:            lit.Pos(),
						End:            lit.End(),
						Message:        violation,
						SuggestedFixes: fixes,
					})
				}
			}
		}
	}

	for _, arg := range call.Args {
		for _, rule := range dataRules {
			rule.Check(pass, arg)
		}
	}
}
