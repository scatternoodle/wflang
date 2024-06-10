package parser

import (
	"fmt"
	"strconv"

	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/token"
)

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

func (p *Parser) parseStringLiteral() (ast.Expression, error) {
	return ast.StringLiteral{Token: p.current}, nil
}

func (p *Parser) parseBooleanLiteral() (ast.Expression, error) {
	exp := ast.BooleanLiteral{
		Token: p.current,
		Value: p.current.Type == token.T_TRUE,
	}
	return exp, nil
}

func (p *Parser) parseIfExpression() (ast.Expression, error) {
	p.trace.trace("IfExpression")
	defer p.trace.untrace("IfExpression")

	eWrap := func(e error) error {
		return fmt.Errorf("error parsing IfExpression: %w", e)
	}

	exp := ast.IfExpression{Token: p.current}

	// if (...
	if err := p.wantPeek(token.T_LPAREN); err != nil {
		return nil, eWrap(err)
	}
	p.advance()

	// ...ConditionStatement,
	cnd, err := p.parseExpression(p_LOWEST)
	if err != nil {
		return nil, eWrap(err)
	}
	exp.Condition = cnd

	if err := p.wantComma(); err != nil {
		return nil, eWrap(err)
	}

	// ...ConsequenceStatement,
	cns, err := p.parseExpression(p_LOWEST)
	if err != nil {
		return nil, eWrap(err)
	}
	exp.Consequence = cns

	// ...AlternativeStatement )
	alt, err := p.parseExpression(p_LOWEST)
	if err != nil {
		return nil, eWrap(err)
	}
	exp.Alternative = alt

	if err := p.wantComma(); err != nil {
		return nil, eWrap(err)
	}

	return exp, nil
}
