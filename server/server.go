package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/scatternoodle/wflang/jrpc2"
	"github.com/scatternoodle/wflang/lsp"
)

type Server struct {
	initialized bool
}

func (srv *Server) ListenAndServe(r io.Reader, w io.Writer) {
	slog.Info("Scanning for messages...")
	scanner := bufio.NewScanner(r)
	scanner.Split(jrpc2.Split)

	for scanner.Scan() {
		srv.handleMessage(w, scanner.Bytes())
	}
	slog.Info("Server stopped listening")
}

func (srv *Server) handleMessage(w io.Writer, msg []byte) {
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

		if err = respond(w, lsp.Initialize(initReq.ID)); err != nil {
			slog.Error("response failed", "error", err)
			return
		}
		slog.Info("InitializeResponse sent")
	}

}

func respond(w io.Writer, v any) error {
	response, err := jrpc2.EncodeMessage(v)
	if err != nil {
		return fmt.Errorf("encoding error: %w", err)
	}
	if _, err = w.Write([]byte(response)); err != nil {
		return fmt.Errorf("write error: %w", err)
	}
	return nil
}
