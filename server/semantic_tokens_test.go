package server

import (
	"reflect"
	"testing"

	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/parser"
)

func TestEncode(t *testing.T) {
	input :=
		`var x = 5;
var y = 10;
var z = x / y;
z * x`

	parser := parser.New(lexer.New(input))
	encoder := newTokenEncoder(parser.Tokens())

	// { deltaLine: 2, deltaStartChar: 5, length: 3, tokenType: 0, tokenModifiers: 3 }
	want := []uint{
		0, 0, 3, 0, 0,
		1, 0, 3, 0, 0,
		1, 0, 3, 0, 0,
	}

	if !reflect.DeepEqual(encoder.semTokens, want) {
		t.Fatalf("have %v, want %v", encoder.semTokens, want)
	}
}
