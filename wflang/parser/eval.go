package parser

import (
	"github.com/scatternoodle/wflang/wflang/ast"
	"github.com/scatternoodle/wflang/wflang/object"
)

// eval descends an AST recursively and evaluates the statements, resolving
// object types and looking for more advanced syntax and semantic errors than what
// the lexer / parser can detect.
func (p *Parser) eval(n ast.Node) object.Object {
	var obj object.Object
	switch v := n.(type) {
	case *ast.AST:
		for _, stmt := range v.Statements {
			obj = p.eval(stmt)
		}

	case ast.ExpressionStatement:
		obj = p.eval(v.Expression)

	case ast.VarStatement:
		obj = object.Variable{
			Name:      v.Name.Value,
			Statement: &v,
			Val:       p.eval(v.Value),
		}
		p.vars = append(p.vars, obj.(object.Variable))

	case ast.NumberLiteral:
		obj = object.Number{Val: v.Val, Static: true}
	case ast.StringLiteral:
		obj = object.String{Val: v.Literal, Static: true}

	default:
		obj = object.Undefined{Val: v}
	}

	return obj
}
