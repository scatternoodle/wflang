package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"path"

	"github.com/scatternoodle/wflang/server/jrpc2"
	"github.com/scatternoodle/wflang/server/lsp"
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
	slog.Info("Server stopped listening")
}

func handleMessage(w io.Writer, msg []byte) {
	method, content, err := jrpc2.DecodeMessage(msg)
	if err != nil {
		slog.Error("Unable to decode", "error", err, "method", method, "message", msg)
		return
	}
	slog.Info("Recieved", "method", method)
	slog.Debug(fmt.Sprintf("content: %s", string(content)))

	switch method {
	case lsp.MethodInitialize:
		var initReq lsp.InitializeRequest
		if err := json.Unmarshal(content, &initReq); err != nil {
			slog.Error("can't marshal request", "error", err)
			return
		}
		if initReq.ID == nil {
			slog.Error("request ID is nil")
			return
		}

		if err = sendResponse(w, lsp.Initialize(initReq.ID)); err != nil {
			slog.Error("response failed", "error", err)
			return
		}
		slog.Info("InitializeResponse sent")
	}

}

func sendResponse(w io.Writer, v any) error {
	response, err := jrpc2.EncodeMessage(v)
	if err != nil {
		return fmt.Errorf("encoding error: %w", err)
	}
	if _, err = w.Write([]byte(response)); err != nil {
		return fmt.Errorf("write error: %w", err)
	}
	return nil
}

func setupLogging(logPath string, level slog.Level) {
	slog.SetLogLoggerLevel(level)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

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

	log.SetOutput(logFile)
}
