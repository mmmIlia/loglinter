package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"github.com/mmmIlia/loglinter/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.NewAnalyzer())
}