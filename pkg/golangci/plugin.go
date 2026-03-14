//go:build plugin
// +build plugin

package golangci

import (
	"github.com/Denis-Mukhametshin-74/selectel-linter/internal/analyzer"
	"golang.org/x/tools/go/analysis"
)

var AnalyzerPlugin plugin

type plugin struct{}

func (plugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}
}
