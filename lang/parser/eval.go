package parser

import (
	"fmt"

	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/object"
)

func (p *Parser) Eval(n ast.Node) object.Object {
	fmt.Printf("node: %T\n", n)
	switch v := n.(type) {

	case *ast.AST:
		if len(v.Statements) > 0 {
			return p.Eval(v.Statements[len(v.Statements)-1])
		}

	case ast.ExpressionStatement:
		return p.Eval(v.Expression)

	case ast.NumberLiteral:
		return object.Number{Val: v.Val}
	}

	return nil
}
