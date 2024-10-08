// parser package provides static analysis of WFLang tokens, producing an ast.AST
// as its main output. The completed AST is then used to serve syntax highlighting,
// code completion and errors/warnings via the Language Server.
package parser

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/scatternoodle/wflang/wflang/ast"
	"github.com/scatternoodle/wflang/wflang/lexer"
	"github.com/scatternoodle/wflang/wflang/object"
	"github.com/scatternoodle/wflang/wflang/token"
)

// Parser is the struct that controls the lexer and produces the AST. It is the
// main entry point for the static analysis process.
type Parser struct {
	l             *lexer.Lexer
	current       token.Token
	next          token.Token
	prefixParsers map[token.Type]prefixParser
	infixParsers  map[token.Type]infixParser
	tokens        []token.Token
	ast           *ast.AST
	errors        []error
	trace         *trace
	vars          []object.Variable
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
		errors:        []error{},
		trace:         &trace{0, &strings.Builder{}},
		vars:          []object.Variable{},
	}

	p.advance()
	p.advance()

	p.prefixParsers[token.T_INT] = p.parseNumberLiteral
	p.prefixParsers[token.T_FLOAT] = p.parseNumberLiteral
	p.prefixParsers[token.T_IDENT] = p.parseIdent
	p.prefixParsers[token.T_MINUS] = p.parsePrefixExpression
	p.prefixParsers[token.T_BANG] = p.parsePrefixExpression
	p.prefixParsers[token.T_STRING] = p.parseStringLiteral
	p.prefixParsers[token.T_TRUE] = p.parseBooleanLiteral
	p.prefixParsers[token.T_FALSE] = p.parseBooleanLiteral
	p.prefixParsers[token.T_LPAREN] = p.parseParenExpression
	p.prefixParsers[token.T_DOLLAR] = p.parseMacroExpression
	p.prefixParsers[token.T_BUILTIN] = p.parseBuiltinCall
	p.prefixParsers[token.T_OVER] = p.parseOverExpression
	p.prefixParsers[token.T_ALIAS] = p.parseAliasExpression
	p.prefixParsers[token.T_WHERE] = p.parseWhereExpression
	p.prefixParsers[token.T_ORDER] = p.parseOrderByExpression
	p.prefixParsers[token.T_DATE] = p.parseDateLiteral
	p.prefixParsers[token.T_TIME] = p.parseTimeLiteral

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
	p.infixParsers[token.T_IN] = p.parseInExpression

	p.ast = p.parse()
	if p.ast != nil {
		p.eval(p.ast)
	}
	return p
}

func (p *Parser) AST() (*ast.AST, error) {
	if p.ast == nil {
		p.ast = p.parse()
		if p.ast == nil {
			return nil, errors.New("AST is nil")
		}
	}
	return p.ast, nil
}

func (p *Parser) Errors() []error             { return p.errors }
func (p *Parser) Tokens() []token.Token       { return p.tokens }
func (p *Parser) Vars() []object.Variable     { return p.vars }
func (p *Parser) Statements() []ast.Statement { return p.ast.Statements }

// parse begins the static analysis process, producing an AST from the token stream
// created by the lexer.
func (p *Parser) parse() *ast.AST {
	slog.Debug("starting new parser run")

	p.trace.trace("AST")
	defer p.trace.untrace("AST")

	AST := &ast.AST{Statements: []ast.Statement{}}

	for p.current.Type != token.T_EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			err = fmt.Errorf("error parsing statement: %w", err)
			p.errors = append(p.errors, err)
			return &ast.AST{}
		}

		AST.Statements = append(AST.Statements, stmt)
		p.advance()
	}

	if len(p.errors) > 0 {
		// logging these errors is still a debug not error level because syntax errors are expected to be ubiquitous while
		// users type code. These errors do not represent a failure state for the parser or the language server.
		slog.Debug("parser run finished with the following errors")
		for _, e := range p.errors {
			slog.Debug("error", "data", e)
		}
	}

	return AST
}

func (p *Parser) Trace() string {
	return p.trace.String()
}

// advance moves the parser forward by one token.
func (p *Parser) advance() {
	p.current = p.next
	if p.next.Literal != "" {
		p.tokens = append(p.tokens, p.current)
	}
	p.next = p.l.NextToken()
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

// isReserved returns true if string is a reserved language keyword.
func isReserved(s string) bool {
	_, isKeyword := lexer.Keyword(s)
	_, isBuiltin := object.Builtins()[s]
	return isKeyword || isBuiltin
}
