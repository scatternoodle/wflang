package main

import (
	"log"
	"log/slog"

	"github.com/scatternoodle/wflang/wflang/repl"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err := repl.Run(); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// 	slog.SetLogLoggerLevel(slog.LevelDebug)

// 	input := `sumTime( over day alias x, hours, where pay_code in set BAMUKI_GEN_COUNTS_AS_WORKED )`
// 	prs := parser.New(lexer.New(input))
// 	repl.PrintTokens(prs)
// }
