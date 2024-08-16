// package jrpc2 is an extremely lightweight JSON-RPC 2.0 implementation, giving me
// the bear minimum needed to make the WFLang LSP function. This is, therefore, NOT
// a general-purpose JRPC2 library!
package jrpc2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

const (
	version = "2.0" // JSON-RPC protocol version

	contentSeparator = "\r\n\r\n"
	fieldSeparator   = "\r\n"
)

// NewRequest returns a JRPC2 request object with the given ID and method. To omit
// ID, pass nil for the ID argument. The JSONRPC version field is automatically set
// to "2.0".
func NewRequest(id *int, method string) Request {
	return Request{JRPC: version, ID: id, Method: method}
}

type Request struct {
	JRPC   string `json:"jsonrpc"`
	ID     *int   `json:"id,omitempty"`
	Method string `json:"method"`
}

// NewResponse returns a JRPC2 response object with the given ID and error. To omit
// ID or ResponseError, pass nil for the respective argument. The JSONRPC version
// field is automatically set to "2.0".
func NewResponse(id *int, respErr *ResponseError) Response {
	return Response{JRPC: version, ID: id, Error: respErr}
}

type Response struct {
	JRPC  string         `json:"jsonrpc"`
	ID    *int           `json:"id"`
	Error *ResponseError `json:"error,omitempty"`
	// Result: defined in enveloping types
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("%s: %+v", e.Message, e.Data)
}

const (
	ERRCODE_PARSE_ERROR      int = -32700
	ERRCODE_INVALID_REQUEST  int = -32600
	ERRCODE_METHOD_NOT_FOUND int = -32601
	ERRCODE_INVALID_PARAMS   int = -32602
	ERRCODE_INTERNAL_ERROR   int = -32603
)

// NewNotification returns a JRPC2 notification object with the given method. The JSONRPC
// version field is automatically set to "2.0".
func NewNotification(method string) Notification {
	return Notification{JRPC: version, Method: method}
}

type Notification struct {
	JRPC   string `json:"jsonrpc"`
	Method string `json:"method"`
	// Params: defined in enveloping types
}

// EncodeMessage takes an interface, marshals it to JSON and wraps it in a message
// string that conforms to JRPC2 spec.
func EncodeMessage(v any) ([]byte, error) {
	msg, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	// msgStr := "Content-Length: " + fmt.Sprint(len(msg)) + contentSeparator + string(msg)
	// return msgStr, nil
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("Content-Length: %d%s", len(msg), contentSeparator))
	if _, err := out.Write(msg); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// DecodeMessage takes a byte slice containing a JRPC2 message and extracts the
// method and contents. Errors if the message is malformed.
//
// Does not properly handle header fields currently - just looks for content-length
// and will have undefined behaviour if content-type or other headers are included.
func DecodeMessage(msg []byte) (method string, content []byte, err error) {
	header, content, ok := bytes.Cut(msg, []byte(contentSeparator))
	if !ok {
		return "", nil, errors.New("missing content separator")
	}

	if !bytes.Contains(header, []byte("Content-Length: ")) {
		return "", nil, errors.New("missing content length")
	}
	cntLenBytes := header[len("Content-Length: "):]
	cntLen, err := strconv.Atoi(string(cntLenBytes))
	if err != nil {
		return "", nil, fmt.Errorf("error parsing content length: %w", err)
	}
	if cntLen != len(content) {
		return "", nil, fmt.Errorf("content length mismatch - header = %d, actual = %d", cntLen, len(content))
	}

	var mthStrct struct {
		Method string `json:"method"`
	}
	if err := json.Unmarshal(content, &mthStrct); err != nil {
		return "", nil, fmt.Errorf("error reading method: %w", err)
	}
	if mthStrct.Method == "" {
		return "", nil, errors.New("missing method field")
	}
	return mthStrct.Method, content, nil
}

// Split is a bufio.SplitFunc that scans tokens for JRPC2 messages
func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, ok := bytes.Cut(data, []byte(contentSeparator))
	if !ok {
		// this is fine, we're just not ready yet
		return 0, nil, nil
	}

	// Content-Length: <number>
	cntLenBytes := header[len("Content-Length: "):]
	cntLen, err := strconv.Atoi(string(cntLenBytes))
	if err != nil {
		return 0, nil, fmt.Errorf("content-length %s is not a number: %v", string(cntLenBytes), err)
	}

	if len(content) < cntLen { // also fine, we just haven't read enough yet
		return 0, nil, nil
	}

	totalLen := len(header) + len(contentSeparator) + cntLen
	return totalLen, data[:totalLen], nil
}
