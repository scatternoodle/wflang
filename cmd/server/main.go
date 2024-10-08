package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path"

	"github.com/scatternoodle/wflang/server"
)

func main() {
	if len(os.Args) < 2 {
		panic("missing arg: logfile path")
	}
	logPath := os.Args[1]
	setupLogging(logPath, true)

	slog.Info("Language Server started.")

	srv := server.New(nil, nil, true) // TODO this should be a program arg
	srv.ListenAndServe(os.Stdin, os.Stdout)
}

func setupLogging(logPath string, debug bool) {
	var level slog.Level
	if debug {
		level = slog.LevelDebug
	} else {
		level = slog.LevelInfo
	}

	logPath = path.Clean(logPath)
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err == os.ErrNotExist { // That's OK, we'll just create it.
		logFile, err = os.Create(logPath)
		if err != nil {
			panic(fmt.Errorf("error creating new logfile: %w", err))
		}
	} else if err != nil {
		panic(fmt.Errorf("error opening logfile: %w", err))
	}

	// slog.SetDefault(slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: level, AddSource: true})))
	slog.SetLogLoggerLevel(level)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)
}
