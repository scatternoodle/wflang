package lexer

import (
	"fmt"
	"strings"
	"testing"

	"github.com/scatternoodle/wflang/lang/token"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		input string
		want  token.Type
	}{
		// keywords / operators / literals
		{`=`, token.T_EQ},
		{`+`, token.T_PLUS},
		{`-`, token.T_MINUS},
		{`!`, token.T_BANG},
		{`*`, token.T_ASTERISK},
		{`/`, token.T_SLASH},
		{`%`, token.T_MODULO},
		{`>`, token.T_GT},
		{`>=`, token.T_GTE},
		{`<`, token.T_LT},
		{`<=`, token.T_LTE},
		{`!=`, token.T_NEQ},
		{`||`, token.T_OR},
		{`&&`, token.T_AND},
		{`or`, token.T_OR},
		{`OR`, token.T_OR},
		{`and`, token.T_AND},
		{`AND`, token.T_AND},
		{`"hello world"`, token.T_STRING},
		{`42`, token.T_NUM},
		{`45.5`, token.T_NUM},
		{`Aardvark`, token.T_IDENT},
		{`var`, token.T_VAR},
		{`Var`, token.T_VAR},
		{`over`, token.T_OVER},
		{`where`, token.T_WHERE},
		{`order`, token.T_ORDER},
		{`by`, token.T_BY},
		{`// comment line`, token.T_COMMENT_LINE},
		{`/* comment/*\nblock */`, token.T_COMMENT_BLOCK},

		// builtins...
		// ...math
		{`if`, token.T_BUILTIN},
		{`min`, token.T_BUILTIN},
		{`max`, token.T_BUILTIN},

		// ...general summary
		{`sum`, token.T_BUILTIN},
		{`count`, token.T_BUILTIN},

		// ...time summary
		{`countTime`, token.T_BUILTIN},
		{`sumTime`, token.T_BUILTIN},

		// ...schedule summary
		{`sumSchedule`, token.T_BUILTIN},
		{`countSchedule`, token.T_BUILTIN},

		// ...exceptions
		{`countException`, token.T_BUILTIN},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			l := New(tt.input)
			tk := l.NextToken()

			if tk.Type != tt.want {
				t.Errorf("type = %s, want %s", tk.Type, tt.want)
			}
		})
	}
}

func TestBuiltins(t *testing.T) {
	tests := []string{
		"if",
		"min",
		"max",
		"sum",
		"count",
		"sumTime",
		"countTime",
		"sumSchedule",
		"countSchedule",
		"countException",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			l := New(tt)
			tk := l.NextToken()

			if tk.Type != token.T_BUILTIN {
				t.Fatalf("type = %s, want %s", tk.Type, token.T_BUILTIN)
			}
			if !strings.EqualFold(tk.Literal, tt) {
				t.Fatalf("literal = %s, want %s", tk.Literal, tt)
			}
		})
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
