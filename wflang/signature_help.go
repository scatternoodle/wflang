package wflang

import "github.com/scatternoodle/wflang/internal/lsp"

func SignatureHelp(pos lsp.Position) (info lsp.SignatureInfo, activeParam int, err error) {
	// CURRENT
	// get all nodes within range
	// find callExpression in those nodes
	// look up function object
	// get params
	// convert them to protocol info
	// determine active param
	// return sig info

	return lsp.SignatureInfo{}, 0, nil
}
