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
	~rune | ~byte | ~string
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

const (
	// misc

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

	// Unhelpfully, WFLang uses a single equal sign for both assignment and equality.
	// There is no semantic use for a double equal sign.
	T_EQ       Type = "="
	T_PLUS     Type = "+"
	T_MINUS    Type = "-"
	T_BANG     Type = "!"
	T_ASTERISK Type = "*"
	T_SLASH    Type = "/"
	T_MODULO   Type = "%" // Modulo is the only semantic use for the percent sign.
	T_LT       Type = ">"
	T_GT       Type = "<"
	T_LTE      Type = "<="
	T_GTE      Type = ">="
	T_AND      Type = "&&" // There is no semantic use for a single ampersand.
	T_OR       Type = "||" // There is no semantic use for a single pipe.

	// delimiters
	T_COMMA       Type = ","
	T_SEMICOLON   Type = ";" // Exclusively to terminate variable declarations.
	T_COLON       Type = ":" // TODO - check - Not sure if this has a semantic use in WFLang.
	T_LPAREN      Type = "("
	T_RPAREN      Type = ")"
	T_LBRACE      Type = "{" // TODO - check - rare - only use case I'm aware of is to express times.
	T_RBRACE      Type = "}"
	T_LBRACKET    Type = "["
	T_RBRACKET    Type = "]"  // For specific array-like use cases such as "in" expressions.
	T_PERIOD      Type = "."  // Period can denote a decimal point or member access.
	T_DOLLAR      Type = "$"  // Dollas signs wrap Macros in WFLang.
	T_DOUBLEQUOTE Type = "\"" // For string literals. Single quotes are not used.

	// keywords

	T_VAR      Type = "VAR"
	T_IF       Type = "IF"
	T_OVER     Type = "OVER"
	T_WHERE    Type = "WHERE"
	T_ORDER_BY Type = "ORDER BY"
)

func keywords() map[string]Type {
	// We define keywords here, but NOT builtin functions. Those are for the parser to worry about.
	return map[string]Type{
		"var":      T_VAR,
		"if":       T_IF,
		"over":     T_OVER,
		"where":    T_WHERE,
		"order by": T_ORDER_BY,

		// WFLang allows the actual words "and" and "or" to be used as logical operators.
		// TODO - check if case sensitive.
		// TODO - shall we just refuse to recognize this? This is an opinionated
		// tool, after all.
		"and": T_AND,
		"or":  T_OR,
	}
}
