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
		AST := prs.Parse()

		errs := prs.Errors()
		if len(errs) > 0 {
			fmt.Printf("Parser has %d errors:\n", len(errs))
			for i, err := range errs {
				fmt.Printf("\t[%d]: %s\n", i, err.Error())
			}
			return nil
		}

		fmt.Println(AST.String())
	}
}
