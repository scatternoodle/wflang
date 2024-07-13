package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/parser"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.Fatal(repl.Run())

	input := `/* 1 */
/* 2.1
2.2 */`

	prs := parser.New(lexer.New(input))

	fmt.Println("errors:")
	for _, e := range prs.Errors() {
		fmt.Printf("%+v\n", e)
	}
	fmt.Print("\n\n")

	fmt.Println("tokens:")
	for _, t := range prs.Tokens() {
		fmt.Printf("%+v\n", t)
	}
}
