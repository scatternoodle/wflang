package parser

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/token"
	testhelp "github.com/scatternoodle/wflang/testhelper"
)

var testParseInput = `var x = 1;`

func TestParse(t *testing.T) {
	_, _ = testRunParser(t, testParseInput, 1, false)
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

			exp := testExpressionStatement(t, AST.Statements[0])
			prefix := testhelp.AssertType[ast.PrefixExpression](t, exp)

			if prefix.Prefix != tt.op {
				t.Errorf("operator: have %s, want %s", prefix.Prefix, tt.op)
			}
			if !testLiteral(t, prefix.Right, tt.want) {
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

			exp := testExpressionStatement(t, AST.Statements[0])
			if !testInfix(t, exp, tt.op, tt.left, tt.right) {
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

		cStmt := testhelp.AssertType[ast.LineCommentStatement](t, stmt)
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

func TestParseFunctionCall(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		argLen int
		err    bool
	}{
		{"min", "min(2, workedHours)", 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, AST := testRunParser(t, tt.input, 1, tt.err)

			exp := testExpressionStatement(t, AST.Statements[0])
			fCall := testhelp.AssertType[ast.FunctionCall](t, exp)

			if len(fCall.Args) != tt.argLen {
				t.Fatalf("have %d arguments, want %d", len(fCall.Args), tt.argLen)
			}
		})
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

		cStmt := testhelp.AssertType[ast.BlockCommentStatement](t, stmt)
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

	exp := testExpressionStatement(t, AST.Statements[0])
	if !testLiteral(t, exp, "hello") {
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
			exp := testExpressionStatement(t, AST.Statements[0])
			if !testLiteral(t, exp, tt.want) {
				return
			}
		})
	}
}

func TestParenExpression(t *testing.T) {
	input := `(
	var x = "foo";
	var y = "bar";
	x + y
)`
	_, AST := testRunParser(t, input, 1, false)

	exp := testExpressionStatement(t, AST.Statements[0])
	parStmt := testhelp.AssertType[ast.ParenExpression](t, exp)

	vars := []tVar{
		{"x", "foo"},
		{"y", "bar"},
	}
	if !testBlockExpression(t, parStmt.Inner, vars) {
		return
	}
}

type tVar struct {
	name string
	val  any
}

func TestMacroExpression(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		ident  string
		params []any
	}{
		{"good, just literal params", `$TEST(42, "foo", true)$`, "TEST", []any{42, "foo", true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, AST := testRunParser(t, tt.input, 1, false)
			exp := testExpressionStatement(t, AST.Statements[0])

			macro := testhelp.AssertType[ast.MacroExpression](t, exp)
			if macro.Name.String() != tt.ident {
				t.Errorf("macro.Name: have %s, want %s", macro.Name.Literal, tt.ident)
			}
			if len(macro.Params) != len(tt.params) {
				t.Fatalf("macro.Params lenght: have %d, want %d", len(macro.Params), len(tt.params))
			}

			for i, param := range macro.Params {
				if !testLiteral(t, param, tt.params[i]) {
					t.Errorf("caught on macro.Params[%d]", i)
					continue
				}
			}
		})
	}
}

func TestOverExpression(t *testing.T) {
	input := `over day`
	_, AST := testRunParser(t, input, 1, false)

	exp := testExpressionStatement(t, AST.Statements[0])
	oExp := testhelp.AssertType[ast.OverExpression](t, exp)

	if oExp.String() != input {
		t.Fatalf("String() = %s, want %s", oExp.String(), input)
	}
}

func TestWhereExpression(t *testing.T) {
	input := `where x > y`
	_, AST := testRunParser(t, input, 1, false)

	exp := testExpressionStatement(t, AST.Statements[0])
	_ = testhelp.AssertType[ast.WhereExpression](t, exp)
}

func TestOrderByExpression(t *testing.T) {
	input := `
order by 1
order by pay_code asc
order by hours desc`

	_, AST := testRunParser(t, input, 3, false)

	for i, stmt := range AST.Statements {
		t.Run(fmt.Sprint(i), func(t *testing.T) {

			exp := testExpressionStatement(t, stmt)
			if !testOrderByExpression(t, exp) {
				return
			}
		})
	}

}

func testOrderByExpression(t *testing.T, exp ast.Expression) bool {
	obExp := testhelp.AssertType[ast.OrderByExpression](t, exp)

	if obExp.Asc != nil {
		ascT := obExp.Asc.Type
		if ascT != token.T_ASC && ascT != token.T_DESC {
			t.Fatalf("Asc token type = %s, want one of [token.T_ASC, token.T_DESC]", ascT)
			return false
		}
	}
	return true
}

func testBlockExpression(t testhelp.TH, exp ast.Expression, vars []tVar) bool {
	blockExp := testhelp.AssertType[ast.BlockExpression](t, exp)

	if len(blockExp.Vars) != len(vars) {
		t.Fatalf("exp.Vars length: have %d, want %d", len(blockExp.Vars), len(vars))
		return false
	}

	for i, v := range vars {
		have := blockExp.Vars[i]
		if have.Name.Value != v.name {
			t.Fatalf("Vars[%d].Name: have %s, want %s", i, have.Name.Value, v.name)
			return false
		}
	}
	return true
}

func testVarStatement(t testhelp.TH, stmt ast.Statement, name string, val any) bool {
	if stmt.TokenLiteral() != "var" {
		t.Errorf(`TokenLiteral(): have %s, want "var"`, stmt.TokenLiteral())
		return false
	}

	vstmt := testhelp.AssertType[ast.VarStatement](t, stmt)

	if vstmt.Name.Value != name {
		t.Errorf("name: have %s, want %s", vstmt.Name.Value, name)
		return false
	}

	if !testLiteral(t, vstmt.Value, val) {
		return false
	}
	return true
}

func testLiteral(t testhelp.TH, exp ast.Expression, want any) bool {
	switch v := want.(type) {
	case int:
		return testNumberLiteral(t, exp, float64(v))
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

func testStringLiteral(t testhelp.TH, exp ast.Expression, want string) bool {
	sStmt := testhelp.AssertType[ast.StringLiteral](t, exp)
	if sStmt.TokenLiteral() != want {
		t.Errorf("literal: have %s, want %s", sStmt.TokenLiteral(), want)
		return false
	}
	return true
}

func testNumberLiteral(t testhelp.TH, exp ast.Expression, want float64) bool {
	nstmt := testhelp.AssertType[ast.NumberLiteral](t, exp)

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

func testBooleanLiteral(t testhelp.TH, exp ast.Expression, want bool) bool {
	bstmt := testhelp.AssertType[ast.BooleanLiteral](t, exp)
	if bstmt.Value != want {
		t.Errorf("value: have %t, want %t", bstmt.Value, want)
		return false
	}
	return true
}

func testInfix(t testhelp.TH, exp ast.Expression, operator string, left, right any) bool {
	infix := testhelp.AssertType[ast.InfixExpression](t, exp)
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

func testRunParser(t testhelp.TH, input string, wantLen int, errOk bool) (*Parser, *ast.AST) {
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

func checkParseErrors(t testhelp.TH, p *Parser) {
	if len(p.errors) != 0 {
		t.Errorf("parser has %d errors:", len(p.errors))
		for _, err := range p.errors {
			t.Error(err.Error())
		}
		t.FailNow()
	}
}

func testExpressionStatement(t testhelp.TH, stmt ast.Statement) ast.Expression {
	eStmt := testhelp.AssertType[ast.ExpressionStatement](t, stmt)
	if eStmt.Expression == nil {
		t.Fatal("expression is nil")
	}

	return eStmt.Expression
}
