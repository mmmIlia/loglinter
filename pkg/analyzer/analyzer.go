package analyzer

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	
	"github.com/mmmIlia/loglinter/pkg/rules"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "loglinter",
		Doc:  "Checks log messages for specific formatting rules",
		Run:  run,
		Requires:[]*analysis.Analyzer{inspect.Analyzer},
	}
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

func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter :=[]ast.Node{
		(*ast.CallExpr)(nil),
	}

	styleRules := []rules.Rule{
		rules.NewLowercaseRule(),
		rules.NewEnglishRule(),
		rules.NewSpecialCharsRule(),
	}

	dataRules := []rules.Rule{
		rules.NewSensitiveDataRule(),
	}

	insp.Preorder(nodeFilter, func(node ast.Node) {
		call := node.(*ast.CallExpr)

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
		for _, rule := range styleRules {
			rule.Check(pass, msgArg)
		}

		for _, arg := range call.Args {
			for _, rule := range dataRules {
				rule.Check(pass, arg)
			}
		}
	})

	return nil, nil
}