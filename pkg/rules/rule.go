package rules

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

type Rule interface {
	Check(pass *analysis.Pass, msgArg ast.Expr)
}