package lsp

const (
	ERRCODE_SERVER_NOT_INITIALIZED int32 = -32002
	ERRCODE_UNKNOWN_ERROR_CODE     int32 = -32001
	ERRCODE_REQUEST_FAILED         int32 = -32803
	ERRCODE_SERVER_CANCELLED       int32 = -32802
)

type Error interface {
	error
	Code() int32
	Message() string
	Data() any
}
