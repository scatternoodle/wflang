package wflang

import "github.com/scatternoodle/wflang/internal/lsp"

func SignatureHelp(pos lsp.Position) (info lsp.SignatureInfo, activeParam int, err error) {
	// get node at pos, return if not a callExpression
	// look up function object
	// get params
	// convert them to protocol info
	// determine active param
	// return sig info

	return lsp.SignatureInfo{}, 0, nil
}
