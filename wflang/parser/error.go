package parser

import (
	"github.com/scatternoodle/wflang/wflang/token"
)

func newParseErr(msg string, tok token.Token) ParseErr {
	return ParseErr{msg, tok}
}

// ParseErr is a struct that represents an error that occurred during the parsing
// process. It implements both expression and statement interfaces in addition to the
// error interface, so that it can also traverse the AST in place of a valid node.
type ParseErr struct {
	Msg   string
	Token token.Token
}

func (p ParseErr) Error() string {
	return p.Msg
}

func (p ParseErr) String() string {
	return p.Error()
}

func (p ParseErr) TokenLiteral() string {
	return p.Token.Literal
}
