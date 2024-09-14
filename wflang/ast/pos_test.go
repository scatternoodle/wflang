package ast_test

import (
	"reflect"
	"testing"

	"github.com/scatternoodle/wflang/wflang/ast"
	"github.com/scatternoodle/wflang/wflang/lexer"
	"github.com/scatternoodle/wflang/wflang/parser"
	"github.com/scatternoodle/wflang/wflang/token"
)

func TestNodeAtPos(t *testing.T) {
	const input = `var workedHrs = sumTime( over day
			, hours
			, where pay_code in set COUNTS_AS_WORKED );
workedHours * 0.5`

	tests := []struct {
		name string
		pos  token.Pos
		want reflect.Type
	}{
		{"VarStatement", token.Pos{Line: 0, Col: 4}, reflect.TypeOf(ast.VarStatement{})},
		{"BuiltinCall", token.Pos{Line: 0, Col: 18}, reflect.TypeOf(ast.BuiltinCall{})},
		{"Ident", token.Pos{Line: 0, Col: 31}, reflect.TypeOf(ast.Ident{})},
		{"NumberLiteral", token.Pos{Line: 3, Col: 15}, reflect.TypeOf(ast.NumberLiteral{})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prs := parser.New(lexer.New(input))
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
