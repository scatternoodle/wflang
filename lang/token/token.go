package token

import (
	"fmt"
	"strings"
)

// LookupKeyword returns the Type of s if it is a keyword, or T_IDENT if not.
func LookupKeyword(s string) Type {
	s = strings.ToLower(s)
	if t, ok := keywords()[s]; ok {
		return t
	}
	return T_IDENT
}

// Type is a string that represents the type of a token. See the consts beginning
// with "T_" in this package.
type Type string

// Types that may be used to build a Token literal.
type Literal interface {
	rune | byte | ~string
}

// New creates a new token with the given type and position data. Takes a single
// byte, which is converted to a string and set as the Literal, with Len set to 1.
// This is to be used within the Lexer, either for single-char literals or as a
// "Starter" for longer tokens.
func New[T Literal](num, line, col int, t Type, lit T) Token {
	return Token{
		Type:    t,
		Literal: string(lit),
		Pos:     Pos{num, line, col},
		Len:     1,
	}
}

// Token represents a word, or "semantic token" in WFLang.
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
