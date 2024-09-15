package parser

import (
	"testing"

	"github.com/scatternoodle/wflang/wflang/object"
	"github.com/scatternoodle/wflang/wflang/types"
	"github.com/scatternoodle/wflang/testhelp"
)

func TestEvalLiterals(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		tp     types.Type
		static bool
		val    any
	}{
		{
			name:  "number literal",
			input: "42",
			tp:    types.T_NUMBER,
			val:   float64(42),
		},
		{
			name:  "string literal",
			input: `"42"`,
			tp:    types.T_STRING,
			val:   `"42"`,
		},
		{
			name:  "variable",
			input: "var x = 1;",
			tp:    types.T_NUMBER,
			val:   float64(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := testRunEval(t, tt.input, -1, false)
			_ = testObjectVal(t, obj, tt.val)
		})
	}
}

func testRunEval(t testhelp.TH, input string, wantLen int, errOK bool) object.Object {
	parser, AST := testRunParser(t, input, wantLen, errOK)
	return parser.eval(AST)
}

func testObjectVal(t testhelp.TH, obj object.Object, want any) bool {
	var val any
	var ok bool

	if val, ok = obj.Value(); !ok {
		t.Error("object has no value")
		return false
	}

	if val != want {
		t.Errorf("object.Value() = %+v, want %+v", val, want)
	}

	return false
}
