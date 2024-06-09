package main

import (
	"fmt"

	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/parser"
)

func main() {
	s := `$MYMACRO(day)$`
	l := lexer.New(s)
	p := parser.New(l)
	ast := p.Parse()

	fmt.Println(ast.String())
}
