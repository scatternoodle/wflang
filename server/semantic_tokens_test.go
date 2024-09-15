package server

import (
	"reflect"
	"slices"
	"testing"

	"github.com/scatternoodle/wflang/wflang/lexer"
	"github.com/scatternoodle/wflang/wflang/parser"
)

func TestEncode(t *testing.T) {
	tokTypes := tokenTypes()

	tests := []struct {
		name  string
		input string
		want  []uint
	}{
		{
			name:  "simple",
			input: `var x = 5;`,
			want: []uint{
				0, 0, 3, uint(slices.Index(tokTypes, semKeyword)), 0, // var
				0, 4, 1, uint(slices.Index(tokTypes, semVariable)), 0, // x
				0, 2, 1, uint(slices.Index(tokTypes, semOperator)), 0, // =
				0, 2, 1, uint(slices.Index(tokTypes, semNumber)), 0, // 5
			},
		},
		{
			name:  "string",
			input: `var x = "hello, world!";`,
			want: []uint{
				0, 0, 3, uint(slices.Index(tokTypes, semKeyword)), 0, // var
				0, 4, 1, uint(slices.Index(tokTypes, semVariable)), 0, // x
				0, 2, 1, uint(slices.Index(tokTypes, semOperator)), 0, // =
				0, 2, 15, uint(slices.Index(tokTypes, semString)), 0, // "hello, world!"
			},
		},
		{
			name:  "multiline blockcomment",
			input: "/*1\n2*/",
			want: []uint{
				0, 0, 3, uint(slices.Index(tokTypes, semComment)), 0, // /*1
				1, 0, 3, uint(slices.Index(tokTypes, semComment)), 0, // 2*/
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := parser.New(lexer.New(tt.input))
			encoder := newTokenEncoder(parser.Tokens())

			if !reflect.DeepEqual(encoder.semTokens, tt.want) {
				t.Fatalf("have %v, want %v", encoder.semTokens, tt.want)
			}
		})
	}
}
