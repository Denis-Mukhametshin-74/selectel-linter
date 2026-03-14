package main

import (
	"github.com/Denis-Mukhametshin-74/selectel-linter/internal/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
