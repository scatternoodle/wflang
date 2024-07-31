package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/parser"
)

func Run() error {
	fmt.Println("Welcome to WFLANG.\nPlease enter formula, or '.q' to quit.")

	scn := bufio.NewScanner(os.Stdin)
	for {
		if !scn.Scan() {
			if err := scn.Err(); err != nil {
				return err
			}
			return nil
		}

		txt := scn.Text()
		if strings.ToLower(txt) == ".q" {
			fmt.Println("Buh-bye.")
			return nil
		}

		prs := parser.New(lexer.New(txt))
		fmt.Println("Tokens:")
		PrintTokens(prs)
		fmt.Println()

		fmt.Println("Statements:")
		PrintStatements(prs)
		fmt.Println()

		fmt.Println("Errors:")
		PrintErrs(prs)
		fmt.Println()
	}
}

func PrintTokens(prs *parser.Parser) {
	for _, token := range prs.Tokens() {
		fmt.Printf("%+v", token)
		fmt.Print("\n")
	}
}

func PrintStatements(prs *parser.Parser) {
	for _, stmt := range prs.Statements() {
		fmt.Println(stmt.String())
	}
}

func PrintErrs(prs *parser.Parser) {
	for _, err := range prs.Errors() {
		fmt.Println(err.Error())
	}
}
