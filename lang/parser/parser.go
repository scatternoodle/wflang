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

	p.prefixParsers[token.T_NUM] = p.parseNumberLiteral
	p.prefixParsers[token.T_IDENT] = p.parseIdent
	p.prefixParsers[token.T_MINUS] = p.parsePrefixExpression
	p.prefixParsers[token.T_BANG] = p.parsePrefixExpression

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
	switch p.current.Type {
	case token.T_VAR:
		return p.parseVarStatement()
	case token.T_COMMENT_LINE:
		return p.parseLineCommentStatement(), nil
	case token.T_COMMENT_BLOCK:
		return p.parseBlockCommentStatement(), nil
	default:
		return p.parseExpressionStatement()
	}
}

// parseExpressionStatement attempts to resolve any statement that does not begin
// with a keyword, and is therefore assumed to just be an expression. These are
// commonplace in WFLang, as all formulas are ultimately expressions.
func (p *Parser) parseExpressionStatement() (ast.ExpressionStatement, error) {
	p.trace.trace("ExpressionStatement")
	defer p.trace.untrace("ExpressionStatement")

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
	p.trace.trace("VarStatement")
	defer p.trace.untrace("VarStatement")

	eWrap := func(e error) error {
		return fmt.Errorf("error parsing var statement at token %+v: %w", p.current, e)
	}
	stmt := ast.VarStatement{Token: p.current}

	// var [T_IDENT]...
	if err := p.wantPeek(token.T_IDENT); err != nil {
		return ast.VarStatement{}, eWrap(err)
	}
	p.advance()

	name, err := p.parseIdent()
	if err != nil {
		return ast.VarStatement{}, eWrap(err)
	}
	stmt.Name = name.(ast.Ident)

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

	for p.next.Type != token.T_SEMICOLON && p.next.Type != token.T_EOF /*&& precedence < p.peekPrecedence()*/ {
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

func (p *Parser) parsePrefixExpression() (ast.Expression, error) {
	p.trace.trace("PrefixExpression")
	defer p.trace.untrace("PrefixExpression")

	exp := ast.PrefixExpression{
		Token:  p.current,
		Prefix: p.current.Literal,
	}
	p.advance()

	right, err := p.parseExpression(p_LOWEST)
	if err != nil {
		return nil, fmt.Errorf("error parsing right expression: %w", err)
	}
	exp.Right = right
	return exp, nil
}

func (p *Parser) parseInfixExpression(left ast.Expression) (ast.Expression, error) {
	p.trace.trace("InfixExpression")
	defer p.trace.untrace("InfixExpression")

	if left == nil {
		return nil, newParseErr("left expression is nil", p.current)
	}

	exp := ast.InfixExpression{
		Token: p.current,
		Left:  left,
		Infix: p.current.Literal,
	}
	p.advance()

	right, err := p.parseExpression(p_LOWEST)
	if err != nil {
		return nil, fmt.Errorf("error parsing right expresion: %w", err)
	}
	exp.Right = right
	return exp, nil
}

// parseIdent attempts to resolve an Identifier expression.
func (p *Parser) parseIdent() (ast.Expression, error) {
	p.trace.trace("Ident")
	defer p.trace.untrace("Ident")

	val := p.current.Literal

	if val == "" {
		return ast.Ident{}, newParseErr("blank ident string", p.current)
	}
	if isKeyword(val) {
		msg := fmt.Sprintf("token %s is a reserved keyword, and cannot be used as an identifier", val)
		return ast.Ident{}, newParseErr(msg, p.current)
	}

	ident := ast.Ident{Token: p.current, Value: val}
	return ident, nil
}

// parseNumberLiteral attempts to resolve a numeric literal expression. The expression
// can can be a float or int type, but will resolve to float64, which is how all WFLang
// numbers are represented in Golang.
func (p *Parser) parseNumberLiteral() (ast.Expression, error) {
	p.trace.trace("NumberLiteral")
	defer p.trace.untrace("NumberLiteral")

	lit := p.current.Literal
	val, err := strconv.ParseFloat(lit, 64)
	if err != nil {
		msg := fmt.Sprintf("error parsing number literal: %s", err)
		return nil, newParseErr(msg, p.current)
	}
	return ast.NumberLiteral{Token: p.current, Value: val}, nil
}

// parseLineCommentStatement returns a LineCommentStatement with the current token.
func (p *Parser) parseLineCommentStatement() ast.LineCommentStatement {
	p.trace.trace("LineCommentStatement")
	defer p.trace.untrace("LineCommentStatement")

	return ast.LineCommentStatement{Token: p.current}
}

// parseBlockCommentStatement returns a BlockCommentStatement with the current token.
func (p *Parser) parseBlockCommentStatement() ast.BlockCommentStatement {
	p.trace.trace("BlockCommentStatement")
	defer p.trace.untrace("BlockCommentStatement")

	return ast.BlockCommentStatement{Token: p.current}
}

func isKeyword(s string) bool {
	return token.LookupKeyword(s) != token.T_IDENT
}
