package main

import (
	"fmt"

	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/token"
)

func main() {
	s := `$MYMACRO(day)$`
	l := lexer.New(s)

	fmt.Println()
	for t := l.NextToken(); t.Type != token.T_EOF; t = l.NextToken() {
		fmt.Print(t.Literal)
	}
	fmt.Println()
}
