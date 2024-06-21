package main

import (
	"log"

	"github.com/scatternoodle/wflang/lang/repl"
)

func main() {
	log.Fatal(repl.Run())
}
