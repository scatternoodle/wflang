package lexer

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/scatternoodle/wflang/lang/token"
)

func TestNextToken(t *testing.T) {
	s := `
	= + - ! * / % > >= < <=
	"hello world"
	`
	l := New(s)

	tests := []struct {
		wantType    token.Type
		wantLiteral string
	}{
		{token.T_EQ, "="},               // 0
		{token.T_PLUS, "+"},             // 1
		{token.T_MINUS, "-"},            // 2
		{token.T_BANG, "!"},             // 3
		{token.T_ASTERISK, "*"},         // 4
		{token.T_SLASH, "/"},            // 5
		{token.T_MODULO, "%"},           // 6
		{token.T_GT, ">"},               // 7
		{token.T_GTE, ">="},             // 8
		{token.T_LT, "<"},               // 9
		{token.T_LTE, "<="},             // 10
		{token.T_STRING, "hello world"}, // 11
		{token.T_EOF, ""},               // last
	}

	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.wantType {
			t.Fatalf("tests[%d] type = %s, want %s", i, tk.Type, tt.wantType)
		}
		if tk.Literal != tt.wantLiteral {
			t.Fatalf("tests[%d] literal = %s, want %s", i, tk.Literal, tt.wantLiteral)
		}
	}
}

func TestPositionInfo(t *testing.T) {
	s := ">\n"
	l := New(s)

	sBytes := []byte(s)
	wantLen := len(sBytes)
	wantLine := bytes.Count(sBytes, []byte("\n"))
	tk := l.NextToken()

	// First, check if token is in the right position
	// lexer advances before returning token so at EOF, pos is 1 greater than input length.
	if l.pos != wantLen+1 {
		t.Fatalf("l.pos = %d, want %d", l.pos, wantLen)
	}
	if l.line != wantLine {
		t.Fatalf("l.line = %d, want %d", l.line, wantLine)
	}

	// Then check the same for the token
	wantPos := token.Pos{
		Num:  wantLen,
		Line: wantLine,
		Col:  0,
	}
	if !reflect.DeepEqual(tk.Pos, wantPos) {
		t.Fatalf("token.Pos = %s, want %s", tk.Pos.String(), wantPos.String())
	}
}
