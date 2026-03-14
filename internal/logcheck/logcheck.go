package logcheck

import (
	"go/ast"
	"go/token"
	"strings"
)

type Info struct {
	Message string
	Pos     token.Pos
}

func IsSlogCall(call *ast.CallExpr) (*Info, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, false
	}

	pkg, ok := sel.X.(*ast.Ident)
	if !ok || pkg.Name != "slog" {
		return nil, false
	}

	switch sel.Sel.Name {
	case "Info", "Error", "Warn", "Debug", "Infof", "Errorf", "Warnf", "Debugf":
		// ok
	default:
		return nil, false
	}

	if len(call.Args) == 0 {
		return nil, false
	}

	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return nil, false
	}

	return &Info{
		Message: strings.Trim(lit.Value, "\""),
		Pos:     call.Pos(),
	}, true
}
