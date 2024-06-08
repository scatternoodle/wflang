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

// Lexer is the font of all semantic tokens. Here be words.
type Lexer struct {
	input string // Holds the entire text context of the Lexer
	pos   int    // Current position in input
	next  int    // Next reading position (char after pos)
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
		tokn = newToken(l, token.T_EOF, eof)
		tokn.Literal = ""

	case '=': // single = used for both assignment and equality (yuck)
		tokn = newToken(l, token.T_EQ, '=')
	case '+':
		tokn = newToken(l, token.T_PLUS, '+')
	case '-':
		tokn = newToken(l, token.T_MINUS, '-')

	case '!':
		if l.peek() == '=' {
			l.advance()
			tokn = newToken(l, token.T_NEQ, "!=")
		} else {
			tokn = newToken(l, token.T_BANG, '!')
		}

	case '*':
		tokn = newToken(l, token.T_ASTERISK, '*')

	case '/':
		if l.peek() == '/' {
			tokn = newToken(l, token.T_COMMENT_LINE, l.readLineComment())
		} else if l.peek() == '*' {
			tokn = newToken(l, token.T_COMMENT_BLOCK, l.readBlockComment())
		} else {
			tokn = newToken(l, token.T_SLASH, '/')
		}

	case '%':
		tokn = newToken(l, token.T_MODULO, '%')

	case '>':
		if l.peek() == '=' {
			l.advance()
			tokn = newToken(l, token.T_GTE, ">=")
		} else {
			tokn = newToken(l, token.T_GT, '>')
		}
	case '<':
		if l.peek() == '=' {
			l.advance()
			tokn = newToken(l, token.T_LTE, "<=")
		} else {
			tokn = newToken(l, token.T_LT, '<')
		}

	// TODO: take a closer look at these... delimeters are a bit more complex.
	case ',':
		tokn = newToken(l, token.T_COMMA, ',')
	case ';':
		tokn = newToken(l, token.T_SEMICOLON, ';')
	case ':':
		tokn = newToken(l, token.T_COLON, ':')
	case '(':
		tokn = newToken(l, token.T_LPAREN, '(')
	case ')':
		tokn = newToken(l, token.T_RPAREN, ')')
	case '{':
		tokn = newToken(l, token.T_LBRACE, '{')
	case '}':
		tokn = newToken(l, token.T_RBRACE, '}')
	case '[':
		tokn = newToken(l, token.T_LBRACKET, '[')
	case ']':
		tokn = newToken(l, token.T_RBRACKET, ']')
	case '.':
		tokn = newToken(l, token.T_PERIOD, '.')
	case '$':
		tokn = newToken(l, token.T_DOLLAR, '$')

	case '"':
		tokn = newToken(l, token.T_STRING, l.readString())

	default:
		if util.IsDigit(l.ch) {
			tokn = newToken(l, token.T_NUM, l.readNumber())
			l.advance()
			return tokn
		}

		if util.IsLetter(l.ch) {
			s := l.readIdent()
			t := token.LookupKeyword(s)
			tokn = newToken(l, t, s)
			l.advance()
			return tokn
		}

		tokn = newToken(l, token.T_ILLEGAL, l.ch)
	}

	l.advance()
	return tokn
}

// newToken is a helper wrapper around token.New(), which inserts additional data
// from the Lexer's internal state. literal can be any type that satisfies the
// token.Literal interface.
func newToken[T token.Literal](l *Lexer, tType token.Type, literal T) token.Token {
	return token.New(l.pos, l.line, l.lPos, tType, literal)
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
		if l.ch == '\n' {
			l.line++
		}
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
		l.advance()
		if l.ch != '.' && !util.IsDigit(l.ch) {
			break
		}

		if l.ch == '.' {
			if dots > 0 {
				break
			}
			dots++
		}
	}
	return l.input[start:l.pos]
}

// readIdent is similar to readString in that it reads a string, but this version
// is specifically for non quote-wrapped strings that resolve to idents or WFLang
// keywords. Only recognizes alphanumeric characters.
func (l *Lexer) readIdent() string {
	start := l.pos
	for {
		l.advance()
		if !util.IsLetter(l.ch) && !util.IsDigit(l.ch) {
			break
		}
	}
	return l.input[start:l.pos]
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
