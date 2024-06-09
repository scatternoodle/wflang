package main

import (
	"fmt"

	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/parser"
)

func main() {
	s := `var x = 1;`
	l := lexer.New(s)
	p := parser.New(l)
	ast := p.Parse()

	if ast != nil {
		fmt.Println(ast.String())
	}

	errs := p.Errors()
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err.Error())
		}
	}
}
