package rules

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

type TextRule interface {
	Apply(text string) (newText string, violations []string)
}

type NodeRule interface {
	Check(pass *analysis.Pass, expr ast.Expr)
}