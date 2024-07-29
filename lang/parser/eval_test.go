package parser

import (
	"testing"

	"github.com/scatternoodle/wflang/lang/object"
	"github.com/scatternoodle/wflang/testhelp"
)

func TestEval(t *testing.T) {
	input := "5"
	parser, AST := testRunParser(t, input, 1, false)
	obj := parser.Eval(AST)
	num := testhelp.AssertType[object.Number](t, obj)
	if num.Val != 5 {
		t.Fatalf("number: have %f, want 5", num.Val)
	}
}
