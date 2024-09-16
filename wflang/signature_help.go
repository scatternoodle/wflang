package wflang

import (
	"fmt"

	"github.com/scatternoodle/wflang/internal/lsp"
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
		var ok bool
		fnct, ok = object.Builtins()[callable.FName()]
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

	fName := callable.FName()
	info = lsp.SignatureInfo{
		Label:         fName, // TODO: needs to be implemented (full signature not just function name)
		ActiveParam:   0,
		Params:        []lsp.ParamInfo{},
		Documentation: object.DocMarkdown(fName),
	}

	for _, fParam := range fnct.Params {
		_ = fParam // we're still working out how to convert this one.
		info.Params = append(info.Params, lsp.ParamInfo{
			Label: [2]int{0, 0},
			// TODO: not yet implemented - doscstring has it but will need to be split out
			Documentation: &lsp.MarkupContent{Kind: lsp.MarkupKindMarkdown, Value: ""},
		})
	}

	// TODO: determine active param

	return info, 0, nil
}
