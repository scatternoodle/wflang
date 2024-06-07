package lexer

import (
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

// Todo - test func for positional testing
