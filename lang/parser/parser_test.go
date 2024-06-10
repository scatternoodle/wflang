package parser

import (
	"strconv"
	"testing"

	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/token"
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

func TestInfixExpression(t *testing.T) {
	tests := []struct {
		input string
		left  any
		op    string
		right any
	}{
		{"1 + 2", float64(1), "+", float64(2)},
		{"1 - 2", float64(1), "-", float64(2)},
		{"1 * 2", float64(1), "*", float64(2)},
		{"1 / 2", float64(1), "/", float64(2)},
		{"1 > 2", float64(1), ">", float64(2)},
		{"1 < 2", float64(1), "<", float64(2)},
		{"1 = 2", float64(1), "=", float64(2)},
		{"1 != 2", float64(1), "!=", float64(2)},
		{"1 >= 2", float64(1), ">=", float64(2)},
		{"1 <= 2", float64(1), "<=", float64(2)},
		{"1 and 2", float64(1), "and", float64(2)},
		{"1 or 2", float64(1), "or", float64(2)},
		{"1 % 2", float64(1), "%", float64(2)},
		{"1 && 2", float64(1), "&&", float64(2)},
		{"1 || 2", float64(1), "||", float64(2)},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, AST := testRunParser(t, tt.input, 1, false)

			exp, ok := AST.Statements[0].(ast.ExpressionStatement)
			if !ok {
				t.Fatalf("statement type: have %T, want ast.ExpressionStatement", AST.Statements[0])
			}
			if !testInfix(t, exp.Expression, tt.op, tt.left, tt.right) {
				return
			}
		})
	}
}

func TestParseLineComment(t *testing.T) {
	input := `// this is a comment
// and another
var x = 1; // comment at end of line`

	tests := []struct {
		stmtIndex int
		literal   string
		start     token.Pos
		end       token.Pos
	}{
		{0, "// this is a comment", token.Pos{Num: 0, Line: 0, Col: 0}, token.Pos{Num: 19, Line: 0, Col: 19}},
		{1, "// and another", token.Pos{Num: 21, Line: 1, Col: 0}, token.Pos{Num: 34, Line: 1, Col: 13}},
		{3, "// comment at end of line", token.Pos{Num: 47, Line: 2, Col: 11}, token.Pos{Num: 71, Line: 2, Col: 35}},
	}

	_, AST := testRunParser(t, input, 4, false)
	for i, tt := range tests {
		stmt := AST.Statements[tt.stmtIndex]

		cStmt, ok := stmt.(ast.LineCommentStatement)
		if !ok {
			t.Fatalf("tests[%d] statement type: have %T, want ast.LineCommentStatement", i, stmt)
		}
		if cStmt.TokenLiteral() != tt.literal {
			t.Fatalf("tests[%d] literal: have %s, want %s", i, stmt.TokenLiteral(), tt.literal)
		}
		if cStmt.Token.StartPos != tt.start {
			t.Fatalf("tests[%d] start position: have %s, want %s", i, cStmt.Token.StartPos.String(), tt.start.String())
		}
		if cStmt.Token.EndPos != tt.end {
			t.Fatalf("tests[%d] end position: have %s, want %s", i, cStmt.Token.EndPos.String(), tt.end.String())
		}
	}
}

func TestParseBlockComment(t *testing.T) {
	input := `/* 1 */
/* 2.1
2.2 */`

	tests := []struct {
		stmtIndex int
		literal   string
		start     token.Pos
		end       token.Pos
	}{
		{0, "/* 1 */", token.Pos{Num: 0, Line: 0, Col: 0}, token.Pos{Num: 6, Line: 0, Col: 6}},
		{1, "/* 2.1\n2.2 */", token.Pos{Num: 8, Line: 1, Col: 0}, token.Pos{Num: 20, Line: 2, Col: 5}},
	}

	_, AST := testRunParser(t, input, 2, false)
	for i, tt := range tests {
		stmt := AST.Statements[tt.stmtIndex]

		cStmt, ok := stmt.(ast.BlockCommentStatement)
		if !ok {
			t.Fatalf("tests[%d] statement type: have %T, want ast.BlockCommentStatement", i, stmt)
		}
		if cStmt.TokenLiteral() != tt.literal {
			t.Fatalf("tests[%d] literal: have %s, want %s", i, stmt.TokenLiteral(), tt.literal)
		}
		if cStmt.Token.StartPos != tt.start {
			t.Fatalf("tests[%d] start position: have %s, want %s", i, cStmt.Token.StartPos.String(), tt.start.String())
		}
		if cStmt.Token.EndPos != tt.end {
			t.Fatalf("tests[%d] end position: have %s, want %s", i, cStmt.Token.EndPos.String(), tt.end.String())
		}
	}
}

func TestStringLiteral(t *testing.T) {
	input := `"hello"`
	_, AST := testRunParser(t, input, 1, false)

	exp, ok := AST.Statements[0].(ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement type: have %T, want ast.ExpressionStatement", AST.Statements[0])
	}
	if !testLiteral(t, exp.Expression, "hello") {
		return
	}
}

func TestBooleanLiteral(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"true", true},
		{"True", true},
		{"TRUE", true},
		{"false", false},
		{"False", false},
		{"FALSE", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, AST := testRunParser(t, tt.input, 1, false)
			exp, ok := AST.Statements[0].(ast.ExpressionStatement)
			if !ok {
				t.Fatalf("statement type: have %T, want ast.ExpressionStatement", AST.Statements[0])
			}
			if !testLiteral(t, exp.Expression, tt.want) {
				return
			}
		})
	}
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
	case string:
		return testStringLiteral(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("unhandled expression type %T", exp)
	return false
}

func testStringLiteral(t testhelper.TH, exp ast.Expression, want string) bool {
	sStmt, ok := exp.(ast.StringLiteral)
	if !ok {
		t.Errorf("expression type: have %T, want ast.StringLiteral", exp)
		return false
	}

	if sStmt.TokenLiteral() != want {
		t.Errorf("literal: have %s, want %s", sStmt.TokenLiteral(), want)
		return false
	}
	return true
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

func testBooleanLiteral(t testhelper.TH, exp ast.Expression, want bool) bool {
	bstmt, ok := exp.(ast.BooleanLiteral)
	if !ok {
		t.Errorf("expression type: have %T, want ast.BooleanLiteral", exp)
		return false
	}

	if bstmt.Value != want {
		t.Errorf("value: have %t, want %t", bstmt.Value, want)
		return false
	}
	return true
}

func testInfix(t testhelper.TH, exp ast.Expression, operator string, left, right any) bool {
	infix, ok := exp.(ast.InfixExpression)
	if !ok {
		t.Errorf("expression type want ast.InfixExpression, got %T", exp)
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
