package parser

import (
	"fmt"
	"log/slog"

	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/object"
)

// eval descends an AST recursively and evaluates the statements, resolving
// object types and looking for more advanced syntax and semantic errors than what
// the lexer / parser can detect.
func (p *Parser) eval(n ast.Node) object.Object {
	// @TEMPLOG
	s, e := n.Pos()
	slog.Debug("node", "string", n.String(), "startPos", s, "endPos", e)

	var obj object.Object
	switch v := n.(type) {
	case *ast.AST:
		for _, stmt := range v.Statements {
			obj = p.eval(stmt)
		}

	case ast.ExpressionStatement:
		obj = p.eval(v.Expression)

	case ast.VarStatement:
		// @TEMPLOG
		slog.Debug("evaluating variable statement")
		slog.Debug("v", "name", v.Name, "statement", v, "name.Value", v.Value)
		slog.Debug(fmt.Sprintf("eval val = %+v", p.eval(v.Value)))

		obj = object.Variable{
			Name:      v.Name.Value,
			Statement: &v,
			Val:       p.eval(v.Value),
		}

		//@TEMPLOG
		slog.Debug(fmt.Sprintf("obj=%+v", obj))

		varObj := obj.(object.Variable)

		//@TEMPLOG
		slog.Debug(fmt.Sprintf("varObj=%+v", varObj))

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
