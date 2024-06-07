package lexer

import (
	"github.com/scatternoodle/wflang/lang/token"
	"github.com/scatternoodle/wflang/util"
)

// New creates a new lexer and advances it into the first byte within the input
// string.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.advance()
	return l
}

type Lexer struct {
	input string // Holds the entire text context of the Lexer
	curr  int    // Current position in input
	next  int    // Next reading position (char after pos.Num)
	ch    byte   // The current character
	line  int    // The current line
	lPos  int    // The position within the current line
}

const eof byte = 0

// NextToken advances the lexer until a token is completed, and returns that token.
// This is the primary way in which the Parser interfaces with the Lexer.
func (l *Lexer) NextToken() token.Token {
	var tokn token.Token
	l.skipWhiteSpace()

	switch l.ch {
	case eof:
		tokn = l.newToken(token.T_EOF, eof)
		tokn.Literal = ""
	}

	return tokn
}

func (l *Lexer) advance() {
	if l.next >= len(l.input) {
		l.ch = eof
	} else {
		l.ch = l.input[l.next]
	}
	l.curr = l.next
	l.next++
}

func (l *Lexer) skipWhiteSpace() {
	for util.IsWhitespace(l.ch) {
		if l.ch == '\n' {
			l.line++
		}
		l.advance()
	}
}

func (l *Lexer) newToken(t token.Type, ch byte) token.Token {
	return token.New(l.curr, l.line, l.lPos, t, ch)
}
