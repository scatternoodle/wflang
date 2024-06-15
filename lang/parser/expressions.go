package parser

import (
	"fmt"
	"strconv"

	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/token"
)

func blankExpression(t token.Token) ast.Expression {
	return ast.BlankExpression{Token: t}
}

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
	if isReserved(val) {
		msg := fmt.Sprintf("token %s is a reserved word, and cannot be used as an identifier", val)
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

// parsesBlockExpression advances through as many VarStatements as needed until
// an expression is met, which must be the end of the Expression.
//
// There can optionally be an unlimited number of VarStatements, but aside from
// that BlockExpression MUST contain ONE Expression, which MUST be at the END.
func (p *Parser) parseBlockExpression() (ast.Expression, error) {
	p.trace.trace("BlockExpression")
	defer p.trace.untrace("BlockExpression")

	blockExp := ast.BlockExpression{Token: p.current, Vars: []ast.VarStatement{}}
	for {
		if p.next.Type == token.T_EOF {
			return nil, newParseErr("EOF reached before BlockStatement end", p.current)
		}

		if p.current.Type == token.T_VAR {
			vStmt, err := p.parseVarStatement()
			if err != nil {
				return nil, fmt.Errorf("varstatment parse error: %w", err)
			}
			blockExp.Vars = append(blockExp.Vars, vStmt)
			p.advance() // past the semicolon.
			continue
		}

		exp, err := p.parseExpression(p_LOWEST)
		if err != nil {
			return nil, fmt.Errorf("value expression parse error: %w", err)
		}
		blockExp.Value = exp
		return blockExp, nil
	}
}

// parseIfExpression - ifExpressions in WFLang look like this:
//
//	if( <Condition>
//	  , <Consequence>
//	  , <Alternative> )
//
// Condition, Consequence, and Alternative are BlockStatements.
func (p *Parser) parseIfExpression() (ast.Expression, error) {
	p.trace.trace("IfExpression")
	defer p.trace.untrace("IfExpression")

	eWrap := func(e error) error {
		return fmt.Errorf("error parsing IfExpression: %w", e)
	}

	exp := ast.IfExpression{Token: p.current}

	// if (...
	if err := p.passIf(token.T_LPAREN); err != nil {
		return nil, eWrap(err)
	}
	// ...Condition,
	cnd, err := p.parseBlockExpression()
	if err != nil {
		return nil, eWrap(err)
	}
	exp.Condition = cnd

	if err := p.passIf(token.T_COMMA); err != nil {
		return nil, eWrap(err)
	}

	// ...Consequence,
	cns, err := p.parseBlockExpression()
	if err != nil {
		return nil, eWrap(err)
	}
	exp.Consequence = cns

	if err := p.passIf(token.T_COMMA); err != nil {
		return nil, eWrap(err)
	}

	// ...Alternative )
	alt, err := p.parseBlockExpression()
	if err != nil {
		return nil, eWrap(err)
	}
	exp.Alternative = alt

	if err := p.passIf(token.T_RPAREN); err != nil {
		return nil, eWrap(err)
	}

	return exp, nil
}

func (p *Parser) parseParenExpression() (ast.Expression, error) {
	p.trace.trace("ParenExpression")
	defer p.trace.untrace("ParenExpression")

	eWrap := func(e error) error {
		return fmt.Errorf("ParenExpression: %w", e)
	}

	parExp := ast.ParenExpression{Token: p.current}
	if p.next.Type == token.T_RPAREN {
		parExp.Inner = blankExpression(p.current)
		parExp.RParen = p.next
	}

	p.advance()
	inner, err := p.parseBlockExpression()
	if err != nil {
		return nil, eWrap(err)
	}
	parExp.Inner = inner

	if err = p.wantPeek(token.T_RPAREN); err != nil {
		return nil, eWrap(err)
	}
	p.advance()

	parExp.RParen = p.current
	return parExp, nil
}

// parseMacroExpression - MacroExpressions in WFLang look like this:
//
//	$<IDENT>([]<MacroParam>)$
//
// TODO - check if you can also do a macro expression without parens - look at GO_LIVE_DATE as example.
func (p *Parser) parseMacroExpression() (ast.Expression, error) {
	p.trace.trace("MacroExpression")
	p.trace.untrace("MacroExpression")

	eWrap := func(e error) error {
		return fmt.Errorf("parseMacroExpression: %s", e)
	}

	// $<IDENT>...
	if err := p.wantPeek(token.T_IDENT); err != nil {
		return nil, eWrap(err)
	}
	p.advance()

	macro := ast.MacroExpression{Token: p.current}
	name, err := p.parseIdent()
	if err != nil {
		return nil, eWrap(err)
	}
	macro.Name = name.(ast.Ident)

	// ...([]<Expression>)$
	if err = p.passIf(token.T_LPAREN); err != nil {
		return nil, eWrap(err)
	}

	macro.Params = []ast.Expression{}
	for {
		param, err := p.parseExpression(p_LOWEST)
		if err != nil {
			return nil, eWrap(err)
		}
		macro.Params = append(macro.Params, param)

		if p.next.Type != token.T_COMMA {
			break
		}
		p.advance()
		p.advance()
	}

	if err = p.wantPeek(token.T_RPAREN); err != nil {
		return nil, eWrap(err)
	}
	p.advance()
	if err = p.wantPeek(token.T_DOLLAR); err != nil {
		return nil, eWrap(err)
	}

	p.advance()
	macro.RDollar = p.current
	return macro, nil
}

// parseFunctionCall - function calls in WFLang always take the same structure:
//
//	Name<Ident>(Args[]<BlockExpression>)
func (p *Parser) parseFunctionCall() (ast.Expression, error) {
	p.trace.trace("FunctionCall")
	p.trace.untrace("FunctionCall")

	wrap := func(e error) error {
		return fmt.Errorf("parseFunctionCall: %w", e)
	}

	funCall := ast.FunctionCall{Token: p.current}

	name := p.current.Literal
	if name == "" {
		return nil, wrap(newParseErr("name cannot be blank", p.current))
	}
	funCall.Name = name

	if err := p.passIf(token.T_LPAREN); err != nil {
		return nil, wrap(err)
	}

	funCall.Args = []ast.Expression{}
	for {
		arg, err := p.parseBlockExpression()
		if err != nil {
			return nil, wrap(err)
		}
		funCall.Args = append(funCall.Args, arg)

		if p.next.Type != token.T_COMMA {
			break
		}
		p.advance()
		p.advance()
	}

	if err := p.wantPeek(token.T_RPAREN); err != nil {
		return nil, wrap(err)
	}
	p.advance()
	funCall.RParen = p.current
	return funCall, nil
}
