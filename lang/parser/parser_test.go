package parser

import (
	"testing"

	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/lexer"
)

func TestParse(t *testing.T) {
	input := "1"
	_, _ = testRunParser(t, input, 1)
}

func testRunParser(t *testing.T, input string, wantLen int) (*Parser, *ast.AST) {
	l := lexer.New(input)
	prg := New(l)
	AST := prg.Parse()
	checkParseErrors(t, prg)

	if AST == nil {
		t.Fatal("AST is nil")
	}
	if AST.Statements == nil {
		t.Fatal("AST.Statements is nil")
	}
	if len(AST.Statements) != wantLen {
		t.Fatalf("AST statements length have %d, want %d", len(AST.Statements), wantLen)
	}
	return prg, AST
}

func checkParseErrors(t *testing.T, p *Parser) {
	if len(p.errors) != 0 {
		t.Errorf("parser has %d errors:", len(p.errors))
		for _, err := range p.errors {
			t.Error(err.Error())
		}
		t.FailNow()
	}
}
