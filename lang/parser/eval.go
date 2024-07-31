package parser

import (
	"fmt"

	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/object"
)

// Eval descends an AST recursively and evaluates the statements, resolving
// object types and looking for more advanced syntax and semantic errors than what
// the lexer / parser can detect.
func (p *Parser) Eval(n ast.Node) object.Object {
	fmt.Printf("node: %T\n", n)

	var obj object.Object
	switch v := n.(type) {
	case *ast.AST:
		for _, stmt := range v.Statements {
			obj = p.Eval(stmt)
		}

	case ast.ExpressionStatement:
		obj = p.Eval(v.Expression)

	case ast.VarStatement:
		obj = object.Variable{
			Name: v.Name.Value,
			Val:  p.Eval(v.Value),
			Pos:  v.StartPos,
		}
		p.vars = append(p.vars, obj.(object.Variable))

	case ast.NumberLiteral:
		obj = object.Number{Val: v.Val, Static: true}
	case ast.StringLiteral:
		obj = object.String{Val: v.Literal, Static: true}
	}

	return obj
}
