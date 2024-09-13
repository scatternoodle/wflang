package token

import (
	"fmt"
)

// Token represents a word, or "semantic token" in WFLang.
type Token struct {
	Type     Type   // The token type - see "T_" consts in this pkg.
	Literal  string // The token expressed as a string literal.
	StartPos Pos    // The starting position of the token.
	EndPos   Pos    // The ending position of the token.
	Len      int    // The length of the token, in bytes.
}

// Type is a string that represents the type of a token. See the consts beginning
// with "T_" in this package.
type Type string

// Types that may be used to build a Token literal.
type Literal interface {
	rune | byte | ~string
}

// Pos represents the textual position of a token
type Pos struct {
	Line uint // Line number.
	Col  uint // Column number within the current line.
}

func (p Pos) String() string {
	return fmt.Sprintf("[l:%d, c:%d]", p.Line, p.Col)
}

// Positional represents a type that has a start and end position, each
// represented by a Pos struct.
type Positional interface {
	Pos() (start, end Pos)
}
