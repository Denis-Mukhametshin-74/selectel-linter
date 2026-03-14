package analyzer

import (
	"go/ast"
	"go/token"

	"github.com/Denis-Mukhametshin-74/selectel-linter/internal/logcheck"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "test-linter",
	Doc:      "checks that log messages start with a lowercase letter",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		logInfo, ok := logcheck.IsSlogCall(call)
		if !ok {
			return
		}

		msg := logInfo.Message
		pos := logInfo.Pos

		checkLowercase(pass, pos, msg)
		checkEnglish(pass, pos, msg)
		checkSpecialChars(pass, pos, msg)
		checkSensitive(pass, pos, msg)
	})

	return nil, nil
}

func checkSensitive(pass *analysis.Pass, pos token.Pos, msg string) {
	panic("unimplemented")
}

func checkSpecialChars(pass *analysis.Pass, pos token.Pos, msg string) {
	panic("unimplemented")
}

func checkEnglish(pass *analysis.Pass, pos token.Pos, msg string) {
	panic("unimplemented")
}

func checkLowercase(pass *analysis.Pass, pos token.Pos, msg string) {
	panic("unimplemented")
}
