package ast

import (
	"fmt"

	"github.com/scatternoodle/wflang/wflang/token"
)

// NodeAtPos walks the AST starting at root and returns a pointer to the Node at
// pos. Only the deepest node enclosing pos will be assigned, with any ancestors
// ignored.
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
		rootStart, rootEnd := root.Pos()
		return nil, fmt.Errorf("invalid position %s for root %s", pos,
			"start="+rootStart.String()+" end="+rootEnd.String())
	}
	return node, nil
}

// NodesEnclosing walks the AST starting at root and returns a slice of nodes
// enclosing the given Pos in order of furthest to closest ancestor.
func NodesEnclosing(root Node, pos token.Pos) (nodes []Node, err error) {
	rootStart, rootEnd := root.Pos()
	if !pos.InRange(rootStart, rootEnd) {
		return nil, fmt.Errorf("invalid position %s for root %s", pos,
			"start="+rootStart.String()+" end="+rootEnd.String())
	}
	var visit inspector = func(n Node) bool {
		nodeStart, nodeEnd := n.Pos()
		if !pos.InRange(nodeStart, nodeEnd) {
			return false
		}
		if _, isAST := n.(*AST); isAST {
			return true // then carry on walking, we don't need to include the AST itself
		}
		nodes = append(nodes, n)
		return true
	}
	Walk(visit, root)
	return nodes, nil
}
