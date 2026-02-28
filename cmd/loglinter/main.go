package main

import (
	"github.com/mmmIlia/loglinter/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.NewAnalyzer())
}
