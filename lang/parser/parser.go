// parser package provides static analysis of WFLang tokens, producing an ast.AST
// as its main output. The completed AST is then used to serve syntax highlighting,
// code completion and errors/warnings via the Language Server.
package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/scatternoodle/wflang/lang/ast"
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

	// register parser functions
	p.prefixParsers[token.T_NUM] = p.parseNumericExpression

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

type (
	prefixParser func() (ast.Expression, error)
	infixParser  func(ast.Expression) (ast.Expression, error)
)

// advance moves the parser forward by one token.
func (p *Parser) advance() {
	p.current = p.next
	p.next = p.l.NextToken()
}

// wantPeek checks if the current token is of the expected type. If not, returns
// a ParseErr wrapping the current token.
func (p *Parser) wantPeek(want token.Type) error {
	if p.next.Type != want {
		msg := fmt.Sprintf("token type: have %s, want %s", p.next.Type, want)
		err := newParseErr(msg, p.next)
		return err
	}
	return nil
}

// --- STATEMENTS ---

// parseStatement is the triage function for the parser. It delegates to the
// appropriate parsing function based on the current token type.
func (p *Parser) parseStatement() (ast.Statement, error) {
	p.trace.trace("STATEMENT")
	defer p.trace.untrace("STATEMENT")

	switch p.current.Type {
	case token.T_VAR:
		return p.parseVarStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseExpressionStatement attempts to resolve any statement that does not begin
// with a keyword, and is therefore assumed to just be an expression. These are
// commonplace in WFLang, as all formulas are ultimately expressions.
func (p *Parser) parseExpressionStatement() (ast.ExpressionStatement, error) {
	stmt := ast.ExpressionStatement{Token: p.current}
	exp, err := p.parseExpression(p_LOWEST)
	if err != nil {
		err = fmt.Errorf("error parsing expression: %w", err)
		return ast.ExpressionStatement{}, err
	}
	stmt.Expression = exp
	return stmt, nil
}

// parseVarStatement resolves a statement following this format:
//
//	var [T_IDENT] = [Expression];
func (p *Parser) parseVarStatement() (ast.VarStatement, error) {
	eWrap := func(e error) error {
		return fmt.Errorf("error parsing var statement at token %+v: %w", p.current, e)
	}
	stmt := ast.VarStatement{Token: p.current}

	// var [T_IDENT]...
	if err := p.wantPeek(token.T_IDENT); err != nil {
		return ast.VarStatement{}, eWrap(err)
	}
	p.advance()

	name, err := p.parseIdentifier()
	if err != nil {
		return ast.VarStatement{}, eWrap(err)
	}
	stmt.Name = name

	// ... = ...
	if err := p.wantPeek(token.T_EQ); err != nil {
		return ast.VarStatement{}, eWrap(err)
	}
	p.advance()
	p.advance()

	// ...[Expression]
	exp, err := p.parseExpression(p_LOWEST)
	if err != nil {
		return ast.VarStatement{}, eWrap(err)
	}
	stmt.Value = exp

	// ...;
	if err := p.wantPeek(token.T_SEMICOLON); err != nil {
		return ast.VarStatement{}, eWrap(err)
	}
	p.advance()
	return stmt, nil
}

// --- EXPRESSIONS ---

// parseExpression, similarly to parseStatement, is mainly a triage function that
// delegates to the appropriate parsing function based on the current token type.
func (p *Parser) parseExpression(precedence int) (ast.Expression, error) {
	prefix, ok := p.prefixParsers[p.current.Type]
	if !ok {
		errMsg := fmt.Sprintf("no prefix parser mapped for token type %s", p.current.Type)
		return nil, newParseErr(errMsg, p.current)
	}

	leftExp, err := prefix()
	if err != nil {
		return nil, fmt.Errorf("error parsing prefix: %w", err)
	}

	for p.next.Type != token.T_SEMICOLON && p.next.Type != token.T_EOF && precedence < p.peekPrecedence() {
		infix, ok := p.infixParsers[p.next.Type]
		if !ok {
			return leftExp, nil
		}

		p.advance()
		leftExp, err = infix(leftExp)
		if err != nil {
			return nil, fmt.Errorf("error parsing infix: %w", err)
		}
	}
	return leftExp, nil
}

// parseIdentifier attempts to resolve an Identifier expression.
func (p *Parser) parseIdentifier() (ast.Identifier, error) {
	val := p.current.Literal

	if val == "" {
		return ast.Identifier{}, newParseErr("blank ident string", p.current)
	}
	if isKeyword(val) {
		msg := fmt.Sprintf("token %s is a reserved keyword, and cannot be used as an identifier", val)
		return ast.Identifier{}, newParseErr(msg, p.current)
	}

	ident := ast.Identifier{Token: p.current, Value: val}
	return ident, nil
}

func (p *Parser) parseNumericExpression() (ast.Expression, error) {
	lit := p.current.Literal
	val, err := strconv.ParseFloat(lit, 64)
	if err != nil {
		msg := fmt.Sprintf("error parsing number literal: %s", err)
		return nil, newParseErr(msg, p.current)
	}
	return ast.NumberLiteral{Token: p.current, Value: val}, nil
}

func isKeyword(s string) bool {
	return token.LookupKeyword(s) != token.T_IDENT
}
