package ast_test

import (
	"reflect"
	"testing"

	"github.com/scatternoodle/wflang/wflang/ast"
	"github.com/scatternoodle/wflang/wflang/lexer"
	"github.com/scatternoodle/wflang/wflang/parser"
	"github.com/scatternoodle/wflang/wflang/token"
)

const nodePosInput = `
var workedHrs = sumTime( over day
		       , hours
		       , where pay_code in set COUNTS_AS_WORKED );
workedHours * 0.5`

func TestNodeAtPos(t *testing.T) {
	tests := []struct {
		name string
		pos  token.Pos
		want reflect.Type
	}{
		{"VarStatement", token.Pos{Line: 1, Col: 4}, reflect.TypeOf(ast.VarStatement{})},
		{"BuiltinCall", token.Pos{Line: 1, Col: 18}, reflect.TypeOf(ast.BuiltinCall{})},
		{"Ident", token.Pos{Line: 1, Col: 31}, reflect.TypeOf(ast.Ident{})},
		{"NumberLiteral", token.Pos{Line: 4, Col: 15}, reflect.TypeOf(ast.NumberLiteral{})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := parser.New(lexer.New(nodePosInput))
			AST, err := prs.AST()
			if err != nil {
				t.Fatal(err)
			}
			node, err := ast.NodeAtPos(AST, tt.pos)
			if err != nil {
				t.Fatal(err)
			}
			if have := reflect.TypeOf(node); have != tt.want {
				t.Fatalf("type: have %s, want %s", have, tt.want)
			}
		})
	}
}

func TestNodesEnclosing(t *testing.T) {
	tests := []struct {
		name string
		pos  token.Pos
		want []reflect.Type
	}{
		{"sumTimeFirstClause", token.Pos{Line: 1, Col: 31},
			[]reflect.Type{
				reflect.TypeOf(ast.VarStatement{}),
				reflect.TypeOf(ast.BuiltinCall{}),
				reflect.TypeOf(ast.BlockExpression{}),
				reflect.TypeOf(ast.OverExpression{}),
				reflect.TypeOf(ast.Ident{}),
			}},
		{"sumTimeWhereClause", token.Pos{Line: 3, Col: 37},
			[]reflect.Type{
				reflect.TypeOf(ast.VarStatement{}),
				reflect.TypeOf(ast.BuiltinCall{}),
				reflect.TypeOf(ast.BlockExpression{}),
				reflect.TypeOf(ast.WhereExpression{}),
				reflect.TypeOf(ast.InExpression{}),
				reflect.TypeOf(ast.SetExpression{}),
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AST, err := parser.New(lexer.New(nodePosInput)).AST()
			if err != nil {
				t.Fatal(err)
			}
			nodes, err := ast.NodesEnclosing(AST, tt.pos)
			if err != nil {
				t.Fatal(err)
			}
			types := make([]reflect.Type, 0, len(nodes))
			for _, n := range nodes {
				types = append(types, reflect.TypeOf(n))
			}
			if !reflect.DeepEqual(types, tt.want) {
				t.Fatalf("have %s want %s", types, tt.want)
			}
		})
	}
}
