package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/scatternoodle/wflang/lang/repl"
	"github.com/scatternoodle/wflang/server/jrpc2"
)

func mainREPL() {
	if err := repl.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	setupLogging(".local/logs/server.log", slog.LevelDebug)

	slog.Info("Language Server started.")
	// log.Fatal(listenAndServe(os.Stdin, os.Stdout))
}

func listenAndServe(r io.Reader, w io.Writer) error {
	slog.Info("Scanning for messages...")
	scanner := bufio.NewScanner(r)
	scanner.Split(jrpc2.Split)

	for scanner.Scan() {
		method, content, err := jrpc2.DecodeMessage(scanner.Bytes())
		_, _, _ = method, content, err
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func setupLogging(logPath string, level slog.Level) {
	slog.SetLogLoggerLevel(level)

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Errorf("error opening logfile: %v", err))
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
