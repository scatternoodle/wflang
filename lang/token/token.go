package token

import "fmt"

type Type string

// New creates a new token with the given type and position data. Takes a single
// byte, which is converted to a string and set as the Literal, with Len set to 1.
// This is to be used within the Lexer, either for single-char literals or as a
// "Starter" for longer tokens.
func New(num, line, col int, t Type, ch byte) Token {
	return Token{
		Type:    t,
		Literal: string(ch),
		Pos:     Pos{num, line, col},
		Len:     1,
	}
}

type Token struct {
	Type    Type   // The token type - see "T_" consts in this pkg.
	Literal string // The token expressed as a string literal.
	Pos     Pos    // The starting position of the token.
	Len     int    // The length of the token, in bytes.
}

// Pos represents the textual position of a token
type Pos struct {
	Num  int // Char position within entire input string.
	Line int // Line number.
	Col  int // Column number within the current line.
}

func (p Pos) String() string {
	return fmt.Sprintf("[n:%d, l:%d, c:%d]", p.Num, p.Line, p.Col)
}

const (
	T_ILLEGAL Type = "ILLEGAL"
	T_EOF     Type = "EOF"

	// literals
	T_IDENT  Type = "IDENT"
	T_INT    Type = "INT"
	T_FLOAT  Type = "FLOAT"
	T_STRING Type = "STRING"

	// comments
	T_LINE_COMMENT  = "COMMENT_LINE"
	T_BLOCK_COMMENT = "COMMENT_BLOCK"

	// operators
	T_EQ Type = "EQUALSIGN" // Unhelpfully, WFLang uses a single equal sign for both assignment and equality
)
