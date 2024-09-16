package wflang

import (
	"fmt"

	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/server/docstring"
	"github.com/scatternoodle/wflang/wflang/ast"
	"github.com/scatternoodle/wflang/wflang/object"
	"github.com/scatternoodle/wflang/wflang/token"
)

func SignatureHelp(root ast.Node, pos token.Pos) (info lsp.SignatureInfo, activeParam int, err error) {
	nodes, err := ast.NodesEnclosing(root, token.Pos(pos))
	if err != nil {
		return lsp.SignatureInfo{}, 0, err
	}
	if len(nodes) == 0 {
		// this is a valid state, we don't error but simply return empty.
		return lsp.SignatureInfo{}, 0, nil
	}
	var callable ast.CallExpression = nil
	for _, node := range nodes {
		switch node.(type) {
		case ast.CallExpression:
			callable = node.(ast.CallExpression)
		}
	}
	if callable == nil {
		// again, no error, this is perfectly valid
		return lsp.SignatureInfo{}, 0, nil
	}

	var fnct object.Function
	switch callable.(type) {
	case ast.BuiltinCall:
		fnct, ok := object.Builtins()[callable.FName()]
		if !ok {
			err = fmt.Errorf("no builtin function found for FName()=%s", callable.FName())
			return lsp.SignatureInfo{}, 0, nil
		}
	case ast.MacroExpression:
		// TODO: once macros properly implemented. For now, return empty
		return lsp.SignatureInfo{}, 0, nil
	default:
		err = fmt.Errorf("cannot evaluate node of type %T val %+v as an ast.CallExpression", callable, callable)
		return lsp.SignatureInfo{}, 0, err
	}

	info = lsp.SignatureInfo{
		Label:       callable.FName(),
		ActiveParam: 0,
		Params:      []lsp.ParamInfo{},
	}
	// CURRENT - get docstring

	// get params
	// convert them to protocol info
	// determine active param
	// return sig info

	return lsp.SignatureInfo{}, 0, nil
}
