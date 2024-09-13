package ast

import (
	"fmt"

	"github.com/scatternoodle/wflang/wflang/token"
)

// NodeAtPos walks the AST starting at root and returns a pointer to the Node at
// pos.
func NodeAtPos(root Node, pos token.Pos) (node Node, err error) {
	var visit inspector = func(n Node) bool {
		nodeStart, nodeEnd := n.Pos()
		if !pos.InRange(nodeStart, nodeEnd) {
			return false
		}
		node = n
		return true
	}
	Walk(visit, root)
	if node == nil {
		return nil, fmt.Errorf(
			"invalid position %s for root %s",
			pos.String(),
			func() string {
				s, e := root.Pos()
				return fmt.Sprintf("start=%s end=%s", s, e)
			}(),
		)
	}
	return node, nil
}
