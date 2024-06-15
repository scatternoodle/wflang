// parser package provides static analysis of WFLang tokens, producing an ast.AST
// as its main output. The completed AST is then used to serve syntax highlighting,
// code completion and errors/warnings via the Language Server.
package parser

import (
	"fmt"
	"strings"

	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/builtins"
	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/token"
)

// Parser is the struct that controls the lexer and produces the AST. It is the
// main entry point for the static analysis process.
type Parser struct {
	l             *lexer.Lexer
	current       token.Token
	next          token.Token
	prefixParsers map[token.Type]prefixParser
	infixParsers  map[token.Type]infixParser
	precedenceMap map[token.Type]int
	errors        []error
	trace         *trace
}

type (
	prefixParser func() (ast.Expression, error)
	infixParser  func(ast.Expression) (ast.Expression, error)
)

// New takes a lexer, creates a new Parser, and advances it into the first token
// within the lexer.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:             l,
		prefixParsers: map[token.Type]prefixParser{},
		infixParsers:  map[token.Type]infixParser{},
		precedenceMap: precedenceMap(),
		errors:        []error{},
		trace:         &trace{0, &strings.Builder{}},
	}

	p.advance()
	p.advance()

	p.prefixParsers[token.T_NUM] = p.parseNumberLiteral
	p.prefixParsers[token.T_IDENT] = p.parseIdent
	p.prefixParsers[token.T_MINUS] = p.parsePrefixExpression
	p.prefixParsers[token.T_BANG] = p.parsePrefixExpression
	p.prefixParsers[token.T_STRING] = p.parseStringLiteral
	p.prefixParsers[token.T_TRUE] = p.parseBooleanLiteral
	p.prefixParsers[token.T_FALSE] = p.parseBooleanLiteral
	p.prefixParsers[token.T_IF] = p.parseIfExpression
	p.prefixParsers[token.T_LPAREN] = p.parseParenExpression
	p.prefixParsers[token.T_DOLLAR] = p.parseMacroExpression
	p.prefixParsers[token.T_BUILTIN] = p.parseFunctionCall

	p.infixParsers[token.T_MINUS] = p.parseInfixExpression
	p.infixParsers[token.T_PLUS] = p.parseInfixExpression
	p.infixParsers[token.T_SLASH] = p.parseInfixExpression
	p.infixParsers[token.T_ASTERISK] = p.parseInfixExpression
	p.infixParsers[token.T_MODULO] = p.parseInfixExpression
	p.infixParsers[token.T_GT] = p.parseInfixExpression
	p.infixParsers[token.T_GTE] = p.parseInfixExpression
	p.infixParsers[token.T_LT] = p.parseInfixExpression
	p.infixParsers[token.T_LTE] = p.parseInfixExpression
	p.infixParsers[token.T_EQ] = p.parseInfixExpression
	p.infixParsers[token.T_NEQ] = p.parseInfixExpression
	p.infixParsers[token.T_AND] = p.parseInfixExpression
	p.infixParsers[token.T_OR] = p.parseInfixExpression

	return p
}

// Parse begins the static analysis process, producing an AST from the token stream
// created by the lexer.
func (p *Parser) Parse() *ast.AST {
	p.trace.trace("AST")
	defer p.trace.untrace("AST")

	AST := &ast.AST{Statements: []ast.Statement{}}

	for p.current.Type != token.T_EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			err = fmt.Errorf("error parsing statement: %w", err)
			p.errors = append(p.errors, err)
			return nil
		}

		AST.Statements = append(AST.Statements, stmt)
		p.advance()
	}

	return AST
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) Trace() string {
	return p.trace.String()
}

// advance moves the parser forward by one token.
func (p *Parser) advance() {
	p.current = p.next
	p.next = p.l.NextToken()
}

// isReserved returns true if string is a reserved language keyword.
func isReserved(s string) bool {
	_, isKeyword := lexer.Keyword(s)
	_, isBuiltin := builtins.Builtins()[s]
	return isKeyword || isBuiltin
}

// wantPeek checks if the next token is of the expected type. If not, returns
// a ParseErr wrapping the next token.
func (p *Parser) wantPeek(want token.Type) error {
	if p.next.Type != want {
		msg := fmt.Sprintf("token type: have %s, want %s", p.next.Type, want)
		err := newParseErr(msg, p.next)
		return err
	}
	return nil
}

// passIf returns error if the next token is not the wanted type, or else returns nil
// and advances the parser twice.
func (p *Parser) passIf(want token.Type) error {
	if err := p.wantPeek(want); err != nil {
		return err
	}
	p.advance()
	p.advance()
	return nil
}
