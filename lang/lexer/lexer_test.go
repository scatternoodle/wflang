package lexer

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/scatternoodle/wflang/lang/token"
)

func TestNextToken(t *testing.T) {
	s := `
	`
	l := New(s)

	tests := []struct {
		wantType    token.Type
		wantLiteral string
	}{
		{token.T_EOF, ""}, // 0
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
	s := "\n"
	l := New(s)

	// First, check if token is in the right position
	tk := l.NextToken()
	sBytes := []byte(s)
	wantLen := len(sBytes)
	wantLine := bytes.Count(sBytes, []byte("\n"))
	if l.pos != wantLen {
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
		t.Fatalf("token.Pos = %s, have %s", tk.Pos.String(), wantPos.String())
	}
}
