package lexer

import (
	"fmt"
	"testing"

	"github.com/scatternoodle/wflang/lang/token"
)

func TestNextToken(t *testing.T) {
	input := `
= + - ! * / % > >= < <=
"hello world"
42
45.5
Aardvark
var
Var
if over where order by
// comment line
/* comment/*
block */
!=
`
	l := New(input)

	tests := []struct {
		wantType    token.Type
		wantLiteral string
	}{
		{token.T_EQ, "="},                                 // 0
		{token.T_PLUS, "+"},                               // 1
		{token.T_MINUS, "-"},                              // 2
		{token.T_BANG, "!"},                               // 3
		{token.T_ASTERISK, "*"},                           // 4
		{token.T_SLASH, "/"},                              // 5
		{token.T_MODULO, "%"},                             // 6
		{token.T_GT, ">"},                                 // 7
		{token.T_GTE, ">="},                               // 8
		{token.T_LT, "<"},                                 // 9
		{token.T_LTE, "<="},                               // 10
		{token.T_STRING, "hello world"},                   // 11
		{token.T_NUM, "42"},                               // 12
		{token.T_NUM, "45.5"},                             // 13
		{token.T_IDENT, "Aardvark"},                       // 14
		{token.T_VAR, "var"},                              // 15
		{token.T_VAR, "Var"},                              // 16
		{token.T_IF, "if"},                                // 17
		{token.T_OVER, "over"},                            // 18
		{token.T_WHERE, "where"},                          // 19
		{token.T_ORDER, "order"},                          // 20
		{token.T_BY, "by"},                                // 21
		{token.T_COMMENT_LINE, "// comment line\n"},       // 22
		{token.T_COMMENT_BLOCK, "/* comment/*\nblock */"}, // 23
		{token.T_NEQ, "!="},                               // 24
		{token.T_EOF, ""},                                 // last
	}

	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.wantType {
			t.Errorf("tests[%d] type = %s, want %s", i, tk.Type, tt.wantType)
		}
		if tk.Literal != tt.wantLiteral {
			t.Errorf("tests[%d] literal = %s, want %s", i, tk.Literal, tt.wantLiteral)
		}
	}
}

func TestPositionInfo(t *testing.T) {
	input := `var x = 1;
x * 42
"so long and thanks for all the fish"`
	l := New(input)

	tests := []struct {
		n     int
		start token.Pos
		end   token.Pos
	}{
		{
			0,
			token.Pos{Num: 0, Line: 0, Col: 0},
			token.Pos{Num: 2, Line: 0, Col: 2},
		}, // var
		{
			1,
			token.Pos{Num: 4, Line: 0, Col: 4},
			token.Pos{Num: 4, Line: 0, Col: 4},
		}, // x
		{
			2,
			token.Pos{Num: 6, Line: 0, Col: 6},
			token.Pos{Num: 6, Line: 0, Col: 6},
		}, // =
		{
			3,
			token.Pos{Num: 8, Line: 0, Col: 8},
			token.Pos{Num: 8, Line: 0, Col: 8},
		}, // 1
		{
			4,
			token.Pos{Num: 9, Line: 0, Col: 9},
			token.Pos{Num: 9, Line: 0, Col: 9},
		}, // ;
		{
			5,
			token.Pos{Num: 11, Line: 1, Col: 0},
			token.Pos{Num: 11, Line: 1, Col: 0},
		}, // x
		{
			6,
			token.Pos{Num: 13, Line: 1, Col: 2},
			token.Pos{Num: 13, Line: 1, Col: 2},
		}, // *
		{
			7,
			token.Pos{Num: 15, Line: 1, Col: 4},
			token.Pos{Num: 16, Line: 1, Col: 5},
		}, // 42
		{
			8,
			token.Pos{Num: 18, Line: 2, Col: 0},
			token.Pos{Num: 54, Line: 2, Col: 36},
		}, // "so long and thanks for all the fish"
		{
			9,
			token.Pos{Num: 55, Line: 2, Col: 36},
			token.Pos{Num: 55, Line: 2, Col: 36},
		}, // EOF
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			tk := l.NextToken()

			if tk.StartPos != tt.start {
				t.Fatalf("start = %v, want %v", tk.StartPos, tt.start)
			}
			if tk.EndPos != tt.end {
				t.Fatalf("end = %v, want %v", tk.EndPos, tt.end)
			}
		})
	}

}
