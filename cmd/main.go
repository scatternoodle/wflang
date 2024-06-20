package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"path"

	"github.com/scatternoodle/wflang/server/jrpc2"
)

// func main() {
// 	if err := repl.Run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	if len(os.Args) < 2 {
		panic("missing arg: logfile path")
	}
	logPath := os.Args[1]

	fmt.Println("started, setting up logging")
	setupLogging(logPath, slog.LevelDebug)

	slog.Info("Language Server started.")
	listenAndServe(os.Stdin, os.Stdout)
}

func listenAndServe(r io.Reader, w io.Writer) {
	slog.Info("Scanning for messages...")
	scanner := bufio.NewScanner(r)
	scanner.Split(jrpc2.Split)

	for scanner.Scan() {
		handleMessage(w, scanner.Bytes())
	}
}

func handleMessage(w io.Writer, msg []byte) {
	method, content, err := jrpc2.DecodeMessage(msg)
	if err != nil {
		slog.Error("Unable to decode", "error", err, "method", method, "message", msg)
		return
	}

	slog.Info("Recieved", "method", method, "content", content)
}

func setupLogging(logPath string, level slog.Level) {
	slog.SetLogLoggerLevel(level)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	logPath = path.Clean(logPath)
	fmt.Printf("Clean path: %s", logPath)
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err == os.ErrNotExist { // That's OK, we'll just create it.
		fmt.Println("logfile does not yet exist, creating...")
		logFile, err = os.Create(logPath)
		if err != nil {
			panic(fmt.Errorf("error creating new logfile: %w", err))
		}

	} else if err != nil {
		panic(fmt.Errorf("error opening logfile: %w", err))
	}

	log.SetOutput(logFile)
}
