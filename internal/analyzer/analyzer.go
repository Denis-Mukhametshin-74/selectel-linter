package analyzer

import (
	"go/ast"
	"go/token"
	"regexp"
	"strings"
	"unicode"

	"github.com/Denis-Mukhametshin-74/selectel-linter/internal/logcheck"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "analyzer",
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
	sensitivePatterns := []struct {
		pattern string
		desc    string
	}{
		{`password\s*[:=]`, "password"},
		{`pass\s*[:=]`, "pass"},
		{`pwd\s*[:=]`, "pwd"},
		{`token\s*[:=]`, "token"},
		{`api[_-]?key\s*[:=]`, "API key"},
		{`secret\s*[:=]`, "secret"},
		{`key\s*[:=]`, "key"},
		{`auth\s*[:=]`, "auth"},
		{`credential\s*[:=]`, "credential"},
	}

	lowerMsg := strings.ToLower(msg)

	for _, sp := range sensitivePatterns {
		matched, _ := regexp.MatchString(sp.pattern, lowerMsg)
		if matched {
			pass.Reportf(pos, "лог-сообщения не должны содержать потенциально чувствительные данные")
			return
		}
	}

	words := strings.Fields(lowerMsg)
	sensitiveWords := map[string]bool{
		"password": true, "pass": true, "pwd": true,
		"token": true, "apikey": true, "api_key": true,
		"secret": true, "key": true, "auth": true, "credential": true,
	}

	for _, word := range words {
		cleanWord := strings.Trim(word, ".,!?:;()[]{}")
		if sensitiveWords[cleanWord] {
			if cleanWord == "key" && strings.Contains(lowerMsg, "monkey") {
				continue
			}
			if cleanWord == "auth" && strings.Contains(lowerMsg, "authenticated") {
				continue
			}
			if cleanWord == "token" && strings.Contains(lowerMsg, "token validated") {
				continue
			}
			pass.Reportf(pos, "лог-сообщения не должны содержать потенциально чувствительные данные")
			return
		}
	}
}
