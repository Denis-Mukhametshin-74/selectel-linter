package analyzer

import (
	"go/ast"
	"go/token"
	"strings"
	"unicode"

	"github.com/Denis-Mukhametshin-74/selectel-linter/internal/logcheck"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "selectel-linter",
	Doc:      "проверяет лог-сообщения на соответствие правилам оформления",
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

func checkLowercase(pass *analysis.Pass, pos token.Pos, msg string) {
	if len(msg) == 0 {
		return
	}
	first := []rune(msg)[0]
	if unicode.IsUpper(first) {
		pass.Reportf(pos, "лог-сообщения должны начинаться со строчной буквы")
	}
}

func checkEnglish(pass *analysis.Pass, pos token.Pos, msg string) {
	for _, r := range msg {
		if r > unicode.MaxASCII {
			pass.Reportf(pos, "лог-сообщения должны быть только на английском языке")
			return
		}
	}
}

func checkSpecialChars(pass *analysis.Pass, pos token.Pos, msg string) {
	for _, r := range msg {
		if !(r >= 'a' && r <= 'z') &&
			!(r >= 'A' && r <= 'Z') &&
			!(r >= '0' && r <= '9') &&
			r != ' ' && r != '-' && r != '\'' {
			pass.Reportf(pos, "лог-сообщения не должны содержать спецсимволы или эмодзи")
			return
		}
	}
}

func checkSensitive(pass *analysis.Pass, pos token.Pos, msg string) {
	sensitiveKeywords := []string{
		"password", "pass", "pwd",
		"token", "api_key", "apikey", "secret",
		"key", "auth", "credential",
	}

	lowerMsg := strings.ToLower(msg)
	for _, keyword := range sensitiveKeywords {
		if strings.Contains(lowerMsg, keyword) {
			pass.Reportf(pos, "лог-сообщения не должны содержать потенциально чувствительные данные")
			return
		}
	}
}
