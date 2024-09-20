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

// Valid returns true if the token is a zero-value, in which case, uniquely, the
// token literal will be an empty string.
func (tk Token) Valid() bool {
	return tk.Literal != ""
}

// Type is a string that represents the type of a token. See the consts beginning
// with "T_" in this package.
type Type string

// Types that may be used to build a Token literal.
type Literal interface {
	rune | byte | ~string
}

// Positional represents a type that has a start and end position, each
// represented by a Pos struct.
type Positional interface {
	Pos() (start, end Pos)
}

// Pos represents the textual position of a token
type Pos struct {
	Line uint // Line number.
	Col  uint // Column number within the current line.
}

func (p Pos) String() string {
	return fmt.Sprintf("[l:%d, c:%d]", p.Line, p.Col)
}

// GT returns true if Pos p is greater than Pos x.
func (p Pos) GT(x Pos) bool {
	return (p.Line > x.Line) || (p.Line >= x.Line && p.Col > x.Col)
}

// GT returns true if Pos p is less than Pos x.
func (p Pos) LT(x Pos) bool {
	return (p.Line < x.Line) || (p.Line <= x.Line && p.Col < x.Col)
}

// GT returns true if Pos p is equal to Pos x.
func (p Pos) EQ(x Pos) bool {
	return p.Line == x.Line && p.Col == x.Col
}

// GT returns true if Pos p is greater than or equal to Pos x.
func (p Pos) GTE(x Pos) bool {
	return p.GT(x) || p.EQ(x)
}

// GT returns true if Pos p is less than or equal to Pos x.
func (p Pos) LTE(x Pos) bool {
	return p.LT(x) || p.EQ(x)
}

// InRange returns true if Pos p is within range start>>end.
func (p Pos) InRange(start, end Pos) bool {
	return start.LTE(p) && p.LTE(end)
}

// Right returns the Pos shifted n columns to the right. Assumes n to be positive,
// so if you pass a negative value then it has the same effect as calling Left().
func (p Pos) Right(n int) Pos {
	return Pos{Line: p.Line, Col: p.Col + uint(n)}
}

// Left returns the Pos shifted n columns to the left. Assumes n to be positive,
// so if you pass a negative value then it has the same effect as calling Right().
func (p Pos) Left(n int) Pos {
	return Pos{Line: p.Line, Col: p.Col - uint(n)}
}
