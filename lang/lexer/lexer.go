// lexer package provides the lexer for WFLang, which is the first step in the
// static analysis process. The lexer reads the input text and produces a stream
// of semantic tokens, which are then used by the parser to generate the AST.
package lexer

import (
	"github.com/scatternoodle/wflang/lang/token"
	"github.com/scatternoodle/wflang/util"
)

// New creates a new lexer and advances it into the first byte within the input
// string.
func New(input string) *Lexer {
	l := &Lexer{input: input, lines: []int{0}}
	l.advance()
	return l
}

// Lexer is the font of all semantic tokens. Here be words.
type Lexer struct {
	input string // Holds the entire text context of the Lexer
	pos   int    // Current position in input
	next  int    // Next reading position (char after pos)
	ch    byte   // The current character
	line  int    // The current line
	lPos  int    // The position within the current line
	lines []int  // Slice holding the lengths of all lines. Updated whenever lexer advances.
}

const eof byte = 0

// NextToken advances the lexer until a token is completed, and returns that token.
// This is the primary way in which the Parser interfaces with the Lexer.
func (l *Lexer) NextToken() token.Token {
	var tokn token.Token
	l.skipWhiteSpace()

	switch l.ch {
	case eof:
		tokn = newToken(l, token.T_EOF, eof, l.here())
		tokn.Literal = ""

	// this is always a one-shot as = is both assignment and equality. There is no double =.
	case '=':
		tokn = newToken(l, token.T_EQ, '=', l.here())
	case '+':
		tokn = newToken(l, token.T_PLUS, '+', l.here())
	case '-':
		tokn = newToken(l, token.T_MINUS, '-', l.here())

	// Bang can either be ! or !=
	case '!':
		if l.peek() == '=' {
			start := l.here()
			l.advance()
			tokn = newToken(l, token.T_NEQ, "!=", start)
		} else {
			tokn = newToken(l, token.T_BANG, '!', l.here())
		}

	// Can also be part of a block comment terminator (*/) but as the first char in a token,
	// this is always a multiplication infix.
	case '*':
		tokn = newToken(l, token.T_ASTERISK, '*', l.here())

	// / starts comments or division infix
	case '/':
		start := l.here()
		if l.peek() == '/' {
			lit := l.readLineComment()
			tokn = newToken(l, token.T_COMMENT_LINE, lit, start)
		} else if l.peek() == '*' {
			lit := l.readBlockComment()
			tokn = newToken(l, token.T_COMMENT_BLOCK, lit, start)
		} else {
			tokn = newToken(l, token.T_SLASH, '/', start)
		}

	// This will always be a modulo infix
	case '%':
		tokn = newToken(l, token.T_MODULO, '%', l.here())

	// GT or GTE infix
	case '>':
		start := l.here()
		if l.peek() == '=' {
			l.advance()
			tokn = newToken(l, token.T_GTE, ">=", start)
		} else {
			tokn = newToken(l, token.T_GT, '>', start)
		}

	// LT or LTE infix
	case '<':
		start := l.here()
		if l.peek() == '=' {
			l.advance()
			tokn = newToken(l, token.T_LTE, "<=", start)
		} else {
			tokn = newToken(l, token.T_LT, '<', start)
		}

	// TODO: take a closer look at these... delimeters are a bit more complex.
	case ',':
		tokn = newToken(l, token.T_COMMA, ',', l.here())
	case ';':
		tokn = newToken(l, token.T_SEMICOLON, ';', l.here())
	case ':':
		tokn = newToken(l, token.T_COLON, ':', l.here())
	case '(':
		tokn = newToken(l, token.T_LPAREN, '(', l.here())
	case ')':
		tokn = newToken(l, token.T_RPAREN, ')', l.here())
	case '{':
		tokn = newToken(l, token.T_LBRACE, '{', l.here())
	case '}':
		tokn = newToken(l, token.T_RBRACE, '}', l.here())
	case '[':
		tokn = newToken(l, token.T_LBRACKET, '[', l.here())
	case ']':
		tokn = newToken(l, token.T_RBRACKET, ']', l.here())
	case '.':
		tokn = newToken(l, token.T_PERIOD, '.', l.here())
	case '$':
		tokn = newToken(l, token.T_DOLLAR, '$', l.here())

	case '&':
		start := l.here()
		if l.peek() != '&' {
			tokn = newToken(l, token.T_ILLEGAL, '&', start)
		} else {
			l.advance()
			tokn = newToken(l, token.T_AND, "&&", start)
		}
	case '|':
		start := l.here()
		if l.peek() != '|' {
			tokn = newToken(l, token.T_ILLEGAL, '|', start)
		} else {
			l.advance()
			tokn = newToken(l, token.T_OR, "||", start)
		}

	// Always a string literal.
	case '"':
		start := l.here()
		tokn = newToken(l, token.T_STRING, l.readString(), start)

	// Otherwise, will be some sort of num / ident, or illegal. This is baking in some opinions about
	default:
		start := l.here()

		if util.IsDigit(l.ch) {
			tokn = newToken(l, token.T_NUM, l.readNumber(), start)
			l.advance()
			return tokn
		}

		if util.IsLetter(l.ch) {
			lit := l.readIdent()
			tType := token.LookupKeyword(lit)
			tokn = newToken(l, tType, lit, start)
			l.advance()
			return tokn
		}

		tokn = newToken(l, token.T_ILLEGAL, l.ch, start)
	}

	l.advance()
	return tokn
}

// newToken creates a new token.Token, drawing Lexer's internal state. literal can
// be any type that satisfies the token.Literal interface.
func newToken[T token.Literal](l *Lexer, tType token.Type, lit T, start token.Pos) token.Token {
	var tLen int
	if s, ok := any(lit).(string); ok {
		tLen = len(s)
	} else {
		tLen = 1
	}

	return token.Token{
		Type:     tType,
		Literal:  string(lit),
		Len:      tLen,
		StartPos: start,
		EndPos: token.Pos{
			Num:  l.pos,
			Line: l.line,
			Col:  l.lPos,
		},
	}
}

// advance safely advances the Lexer further into its input string, correctly handling
// EOF and internal state updates. This is the only way you should advance the Lexer.
func (l *Lexer) advance() {
	if l.next >= len(l.input) {
		l.ch = eof
	} else {
		l.ch = l.input[l.next]
	}
	l.pos = l.next
	l.next++

	// next, we're going to update line state, but not if we're EOF - that would
	// result in overflowing the input string later, so we're just done.
	if l.ch == eof {
		return
	}

	// we increment length even if newline, because that's still technically a char on
	// the current line. We may need to change this depending on how this impacts
	// behaviour in the Language server / text editor
	l.lines[l.line]++

	if l.pos > 0 {
		if l.input[l.pos-1] == '\n' {
			l.lPos = 0
		} else {
			l.lPos++
		}
	}

	// TODO - think about whether we need to handle any other bytes here?
	if l.ch == '\n' {
		l.line++
		l.lines = append(l.lines, 0)
	}
}

// peek returns the character in the next position without advancing the Lexer.
func (l *Lexer) peek() byte {
	if l.next >= len(l.input) {
		return eof
	}
	return l.input[l.next]
}

// skipWhiteSpace advances until a non-whitespace character is met. Increments the
// Lexer's line number whenever \n is encountered.
func (l *Lexer) skipWhiteSpace() {
	for util.IsWhitespace(l.ch) {
		l.advance()
	}
}

// readString returns the string literal between two \" characters. Call this at
// the first \"
func (l *Lexer) readString() string {
	start := l.pos + 1 // +1 because this is called from first '"' in the token.
	for {
		l.advance()
		if l.ch == '"' || l.ch == eof {
			break
		}
	}
	return l.input[start:l.pos]
}

// readNumber is called on a digit, and reads until it reaches the end of the number.
// A "number" can be a an int or a float, and apart from digits, up to one decimal point
// is recognized. Returns the number as a string.
func (l *Lexer) readNumber() string {
	start := l.pos
	dots := 0 // we allow up to one decimal place

	for {
		if l.peek() != '.' && !util.IsDigit(l.peek()) {
			break
		}

		if l.peek() == '.' {
			if dots > 0 {
				break
			}
			dots++
		}
		l.advance()
	}
	return l.input[start : l.pos+1]
}

// readIdent is similar to readString in that it reads a string, but this version
// is specifically for non quote-wrapped strings that resolve to idents or WFLang
// keywords. Only recognizes alphanumeric characters.
func (l *Lexer) readIdent() string {
	start := l.pos
	for {
		if !util.IsLetter(l.peek()) && !util.IsDigit(l.peek()) {
			break
		}
		l.advance()
	}
	return l.input[start : l.pos+1]
}

// readLineComment advances from a '//' symbol until it reaches a newline or EOF.
// Returns the full comment string, including the '//' symbol.
func (l *Lexer) readLineComment() string {
	start := l.pos
	for {
		l.advance()
		if l.ch == '\n' || l.ch == eof {
			break
		}
	}
	return l.input[start : l.pos+1]
}

// readBlockComment advances from a '/*' opening symbol until it reaches the closing
// '*/' symbol or EOF. Returns the full comment string, including the opening and
// closing symbols.
func (l *Lexer) readBlockComment() string {
	start := l.pos
	for {
		l.advance()
		if l.ch == eof {
			break
		}
		if l.ch == '*' {
			if l.peek() == '/' {
				break
			}
		}
	}
	l.advance()
	return l.input[start : l.pos+1]
}

// here generates a token.Pos object for the current location of the lexer. This
// is used as a helper for generating start and end positions for new tokens.
func (l *Lexer) here() token.Pos {
	return token.Pos{
		Num:  l.pos,
		Line: l.line,
		Col:  l.lPos,
	}
}
