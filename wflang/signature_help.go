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

func SignatureHelp(root ast.Node, pos token.Pos) (info lsp.SignatureInfo, activeParam int, err error) {
	slog.Debug("called with", "root", root, "pos", pos)
	// we consider the requested pos to be for the character BEHIND the cursor,
	// and need to adjust for this.
	if pos.Col > 0 {
		pos.Col--
	}

	// find the callable at pos, if existing
	nodes, err := ast.NodesEnclosing(root, pos)
	slog.Debug("found enclosing nodes", "nodes", nodes, "err", err)
	if err != nil {
		return lsp.SignatureInfo{}, 0, err
	}
	if len(nodes) == 0 {
		// this is a valid state, we don't error but simply return empty.
		return lsp.SignatureInfo{}, 0, nil
	}
	var callable ast.CallExpression = nil
	slog.Debug("searching nodes for callable")
	for i, node := range nodes {
		slog.Debug(fmt.Sprintf("index %d, node: %+v", i, node))
		switch node.(type) {
		case ast.CallExpression:
			callable = node.(ast.CallExpression)
		}
	}
	if callable == nil {
		slog.Debug("no callable found, returning empty")
		// again, no error, this is perfectly valid
		return lsp.SignatureInfo{}, 0, nil
	}
	slog.Debug("callable found", "callable", callable.FName())

	// resolve a builtin function, method or macro call from the retrieved callable
	var function object.Function
	switch callable.(type) {
	case ast.BuiltinCall:
		var ok bool
		function, ok = object.Builtin(callable.FName())
		if !ok {
			err = fmt.Errorf("no builtin function found for FName()=%s", callable.FName())
			slog.Debug(err.Error())
			return lsp.SignatureInfo{}, 0, nil
		}
	case ast.MacroExpression:
		// TODO: once macros properly implemented. For now, return empty
		slog.Debug("it's a macro, so returning empty")
		return lsp.SignatureInfo{}, 0, nil
	default:
		err = fmt.Errorf("cannot evaluate node of type %T val %+v as an ast.CallExpression", callable, callable)
		slog.Error(err.Error())
		return lsp.SignatureInfo{}, 0, err
	}

	// get the function documentation and start building the signature info
	doc, ok := docstring.BuiltinDoc(function.Name)
	if !ok {
		err = fmt.Errorf("FunctionDoc missing for builtin %s", function.Name)
		slog.Error(err.Error())
		return lsp.SignatureInfo{}, 0, err
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

	// determine the active parameter
	slog.Debug("determining active param", "params length", len(callable.Params()))
	var posInRange bool
	if len(callable.Params()) == 0 {
		posInRange = true
	}
	if len(callable.Params()) > len(paramInfos) {
		err = fmt.Errorf("callable param len %d exceeds stored param info len %d for the function",
			len(callable.Params()), len(paramInfos))
		return lsp.SignatureInfo{}, 0, err
	}
	for i, pExp := range callable.Params() {
		pStart, pEnd := pExp.Pos()
		_, callEnd := callable.(ast.Node).Pos()
		slog.Debug("param", "i", i, "str", pExp.String(), "start", pStart, "end", pEnd)
		if pos.LT(pStart) || callEnd.LT(pos) {
			slog.Debug("pos not in range", "pos", pos)
			continue
		}
		if i > 0 {
			_, prvEnd := callable.Params()[i-1].Pos()
			if pos.LT(prvEnd.Right(1)) { // the comma after each param
				continue
			}
		}
		posInRange = true
		if pos.LT(pEnd) {
			activeParam = i
		} else if i < len(callable.Params()) { // it's the next one after i, that the user hasn't typed yet
			activeParam = i + 1
		} else {
			posInRange = false // something really weird is happening...
		}
	}
	if !posInRange {
		err = fmt.Errorf("pos %+v not within range of callable parameters for %+v", pos, callable)
		return lsp.SignatureInfo{}, 0, err
	}

	info = lsp.SignatureInfo{
		Label:         doc.Signature,
		ActiveParam:   activeParam,
		Params:        paramInfos,
		Documentation: &lsp.MarkupContent{Kind: lsp.MarkupKindMarkdown, Value: doc.String()},
	}
	return info, activeParam, nil
}
