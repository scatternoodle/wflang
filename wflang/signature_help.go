package wflang

import (
	"fmt"
	"log/slog"

	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/server/docstring"
	"github.com/scatternoodle/wflang/wflang/ast"
	"github.com/scatternoodle/wflang/wflang/object"
	"github.com/scatternoodle/wflang/wflang/token"
)

func SignatureHelp(root ast.Node, pos token.Pos) (*lsp.SignatureInfo, int, error) {
	// we consider the requested pos to be for the character BEHIND the cursor,
	// and need to adjust for this.
	pos = pos.Left(1)

	// find the callable at pos, if existing
	nodes, err := ast.NodesEnclosing(root, pos)
	if err != nil {
		slog.Error(err.Error())
		return &lsp.SignatureInfo{}, 0, err
	}
	if len(nodes) == 0 {
		// this is a valid state, we don't error but simply return empty.
		return &lsp.SignatureInfo{}, 0, nil
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
		return &lsp.SignatureInfo{}, 0, nil
	}

	// resolve a builtin function, method or macro call from the retrieved callable
	var function object.Function
	switch callable.(type) {
	case ast.BuiltinCall:
		var ok bool
		function, ok = object.Builtin(callable.FName())
		if !ok {
			err = fmt.Errorf("no builtin function found for FName()=%s", callable.FName())
			slog.Error(err.Error())
			return &lsp.SignatureInfo{}, 0, nil
		}
	case ast.MacroExpression:
		// TODO: once macros properly implemented. For now, return empty
		return &lsp.SignatureInfo{}, 0, nil
	default:
		err = fmt.Errorf("cannot evaluate node of type %T val %+v as an ast.CallExpression", callable, callable)
		slog.Error(err.Error())
		return &lsp.SignatureInfo{}, 0, err
	}

	// get the function documentation and start building the signature info
	doc, ok := docstring.BuiltinDoc(function.Name)
	if !ok {
		err = fmt.Errorf("FunctionDoc missing for builtin %s", function.Name)
		slog.Error(err.Error())
		return &lsp.SignatureInfo{}, 0, err
	}
	paramInfos := make([]lsp.ParamInfo, len(doc.Params))
	for i, pDoc := range doc.Params {
		paramInfos[i] = lsp.ParamInfo{
			Label: pDoc.Label,
			Documentation: &lsp.MarkupContent{
				Kind:  lsp.MarkupKindMarkdown,
				Value: pDoc.String(),
			},
		}
	}
	if !function.Variadic() && len(callable.Params()) > len(paramInfos) {
		err = fmt.Errorf("callable param len %d exceeds stored param info len %d for the function",
			len(callable.Params()), len(paramInfos))
		slog.Error(err.Error())
		return &lsp.SignatureInfo{}, 0, err
	}
	info := &lsp.SignatureInfo{
		Label:         doc.Signature,
		Params:        paramInfos,
		Documentation: &lsp.MarkupContent{Kind: lsp.MarkupKindMarkdown, Value: doc.String()},
		ActiveParam:   getActiveParam(pos, callable.Params(), len(paramInfos)),
	}
	return info, info.ActiveParam, nil
}

func getActiveParam(pos token.Pos, params []ast.Expression, max int) int {
	if len(params) == 0 {
		return 0
	}

	var i int
	var param ast.Expression
	for {
		if i >= len(params) {
			break
		}
		if i == max-1 {
			return i
		}
		param = params[i]
		_, end := param.Pos()
		if i == max-1 {
			end = end.Right(1) // the comma after the param
		}
		if pos.LTE(end) {
			return i
		}
		i++
	}
	return i
}
