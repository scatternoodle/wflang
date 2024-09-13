package wflang

import "github.com/scatternoodle/wflang/internal/lsp"

func SignatureHelp(pos lsp.Position) (info lsp.SignatureInfo, activeParam int, err error) {

	return lsp.SignatureInfo{}, 0, nil
}
