package main

import (
	"log"

	"github.com/scatternoodle/wflang/lang/repl"
)

func main() {
	if err := repl.Run(); err != nil {
		log.Fatal(err)
	}
}
