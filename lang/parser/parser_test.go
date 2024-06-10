package parser

import (
	"strconv"
	"testing"

	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/testhelper"
)

var testParseInput = `var x = 1;`

func TestParse(t *testing.T) {
	_, _ = testRunParser(t, testParseInput, 1, false)
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = testRunParser(b, testParseInput, 1, false)
	}
}

func TestParseVarStatement(t *testing.T) {
	tests := []struct {
		input   string
		vName   string
		want    any
		wantErr bool
	}{
		{"var x = 1;", "x", float64(1), false},
		{"var y = 1", "", nil, true},
		{"var x 1;", "", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			prg, AST := testRunParser(t, tt.input, 1, tt.wantErr)

			if tt.wantErr {
				if len(prg.errors) != 0 {
					return
				}
				t.Fatal("did not error when expected")
			}
			if !tt.wantErr && len(prg.errors) > 0 {
				t.Errorf("have %d unexpected error(s):", len(prg.errors))
				for _, err := range prg.errors {
					t.Error(err.Error())
				}
				t.FailNow()
			}

			if !testVarStatement(t, AST.Statements[0], tt.vName, tt.want) {
				return
			}
		})
	}
}

func TestPrefixExpression(t *testing.T) {
	tests := []struct {
		input string
		op    string
		want  any
	}{
		{"!2", "!", float64(2)},
		{"-15.5", "-", float64(15.5)},
		// {"!true;", "!", true}, // TODO - enable once bool parsing implemented.
		// {"!false;", "!", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, AST := testRunParser(t, tt.input, 1, false)

			exp, ok := AST.Statements[0].(ast.ExpressionStatement)
			if !ok {
				t.Fatalf("statement type: have %T, want ast.ExpressionStatement", AST.Statements[0])
			}
			pre, ok := exp.Expression.(ast.PrefixExpression)
			if !ok {
				t.Fatalf("expression type: have %T, want ast.PrefixExpression", exp.Expression)
			}

			if pre.Prefix != tt.op {
				t.Errorf("operator: have %s, want %s", pre.Prefix, tt.op)
			}
			if !testLiteral(t, pre.Right, tt.want) {
				return
			}
		})
	}
}

func testInfixExpression(t *testing.T) {
	// TODO
}

func testVarStatement(t testhelper.TH, stmt ast.Statement, name string, val any) bool {
	if stmt.TokenLiteral() != "var" {
		t.Errorf(`TokenLiteral(): have %s, want "var"`, stmt.TokenLiteral())
		return false
	}

	vstmt, ok := stmt.(ast.VarStatement)
	if !ok {
		t.Errorf("statement type: have %T, want ast.VarStatement", stmt)
		return false
	}

	if vstmt.Name.Value != name {
		t.Errorf("name: have %s, want %s", vstmt.Name.Value, name)
		return false
	}

	if !testLiteral(t, vstmt.Value, val) {
		return false
	}
	return true
}

func testLiteral(t testhelper.TH, exp ast.Expression, want any) bool {
	switch v := want.(type) {
	case float64:
		return testNumberLiteral(t, exp, v)
	}
	t.Errorf("unhandled expression type %T", exp)
	return false
}

func testNumberLiteral(t testhelper.TH, exp ast.Expression, want float64) bool {
	nstmt, ok := exp.(ast.NumberLiteral)
	if !ok {
		t.Errorf("expression type: have %T, want ast.NumberLiteral", exp)
		return false
	}

	if nstmt.Value != want {
		t.Errorf("value: have %f, want %f", nstmt.Value, want)
		return false
	}

	litNum, err := strconv.ParseFloat(nstmt.Token.Literal, 64)
	if err != nil {
		t.Errorf("error parsing literal %s, err: %v", nstmt.Token.Literal, err)
		return false
	}
	if litNum != want {
		t.Errorf("literal: have %f, want %f", litNum, want)
		return false
	}
	return true
}

func testInfix(t testhelper.TH, exp ast.Expression, operator string, left, right any) bool {
	infix, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("expression type want *ast.InfixExpression, got %T", exp)
		return false
	}
	if !testLiteral(t, infix.Left, left) {
		return false
	}
	if infix.Infix != operator {
		t.Errorf("operator want %s, got %s", operator, infix.Infix)
		return false
	}
	if !testLiteral(t, infix.Right, right) {
		return false
	}
	return true
}

func testRunParser(t testhelper.TH, input string, wantLen int, errOk bool) (*Parser, *ast.AST) {
	l := lexer.New(input)
	prg := New(l)
	AST := prg.Parse()

	if !errOk {
		checkParseErrors(t, prg)
	} else {
		return prg, AST
	}

	if AST == nil {
		t.Fatal("AST is nil")
		return nil, nil
	}
	if AST.Statements == nil {
		t.Fatal("AST.Statements is nil")
	}
	if len(AST.Statements) != wantLen {
		t.Fatalf("AST statements length have %d, want %d", len(AST.Statements), wantLen)
	}
	return prg, AST
}

func checkParseErrors(t testhelper.TH, p *Parser) {
	if len(p.errors) != 0 {
		t.Errorf("parser has %d errors:", len(p.errors))
		for _, err := range p.errors {
			t.Error(err.Error())
		}
		t.FailNow()
	}
}
