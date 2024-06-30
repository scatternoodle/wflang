package server

import (
	"testing"

	"github.com/scatternoodle/wflang/lang/token"
	"github.com/scatternoodle/wflang/lsp"
)

func TestEncodeToken(t *testing.T) {
	tests := []struct {
		name  string
		token token.Token
		want  []lsp.Uint
		found bool
	}{
		{
			name: string(token.T_VAR),
			token: token.Token{
				Type:     token.T_VAR,
				Literal:  "var",
				StartPos: token.Pos{Num: 0, Line: 0, Col: 0},
				EndPos:   token.Pos{Num: 2, Line: 0, Col: 2},
				Len:      3,
			},
			want:  []lsp.Uint{0, 0, 3, lsp.Uint(semIndexKeyword), 0},
			found: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have, found := encodeToken(tt.token)

			if !tt.found && found {
				t.Fatal("found token when expected not to")
			}
			if tt.found && !found {
				t.Fatal("token not found when expected")
			}
			if len(have) != 5 {
				t.Fatalf("array len = %d, want 5", len(have))
			}

			var equal bool
			for i := range 5 {
				equal = have[i] == tt.want[i]
			}
			if !equal {
				t.Fatalf("token encoding: have %v, want %v", have, tt.want)
			}
		})
	}
}
