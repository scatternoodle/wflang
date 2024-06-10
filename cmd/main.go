package main

import (
	"fmt"

	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/parser"
)

func main() {
	s := `// hi there`
	l := lexer.New(s)
	p := parser.New(l)
	ast := p.Parse()

	fmt.Print("\n\nAST:\n\n")
	if ast != nil {
		fmt.Println(ast.String())
	}

	fmt.Print("\n\nErrors:\n\n")
	errs := p.Errors()
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err.Error())
		}
	}

	fmt.Print("\n\nTrace:\n\n")
	fmt.Println(p.Trace())
}
