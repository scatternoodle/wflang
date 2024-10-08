package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/scatternoodle/wflang/wflang/ast"
	"github.com/scatternoodle/wflang/wflang/token"
	"github.com/scatternoodle/wflang/wflang/types/wdate"
)

func blankExpression(t token.Token) ast.Expression {
	return ast.BlankExpression{Token: t}
}

// parseExpression, similarly to parseStatement, is mainly a triage function that
// delegates to the appropriate parsing function based on the current token type.
func (p *Parser) parseExpression() (ast.Expression, error) {
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

	right, err := p.parseExpression()
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

	right, err := p.parseExpression()
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
	return ast.NumberLiteral{Token: p.current, Val: val}, nil
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
		if p.current.Type == token.T_VAR {
			vStmt, err := p.parseVarStatement()
			if err != nil {
				return nil, fmt.Errorf("varstatment parse error: %w", err)
			}
			blockExp.Vars = append(blockExp.Vars, vStmt)
			p.advance() // past the semicolon.
			continue
		}

		exp, err := p.parseExpression()
		if err != nil {
			return nil, fmt.Errorf("value expression parse error: %w", err)
		}
		blockExp.Value = exp
		return blockExp, nil
	}
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
	defer p.trace.untrace("MacroExpression")

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
	if err = p.wantPeek(token.T_LPAREN); err != nil {
		return nil, eWrap(err)
	}
	p.advance()
	macro.LPar = p.current
	p.advance()

	macro.Args = []ast.Expression{}
	for {
		param, err := p.parseExpression()
		if err != nil {
			return nil, eWrap(err)
		}
		macro.Args = append(macro.Args, param)

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
	macro.RPar = p.current
	if err = p.wantPeek(token.T_DOLLAR); err != nil {
		return nil, eWrap(err)
	}

	p.advance()
	macro.RDollar = p.current
	return macro, nil
}

// parseBuiltinCall - function calls in WFLang always take the same structure:
//
//	Name<Ident>(Args[]<BlockExpression>)
func (p *Parser) parseBuiltinCall() (ast.Expression, error) {
	p.trace.trace("BuiltinCall")
	defer p.trace.untrace("BuiltinCall")

	wrap := func(e error) error {
		return fmt.Errorf("parseBuiltinCall: %w", e)
	}

	call := ast.BuiltinCall{Token: p.current}

	name := strings.ToLower(p.current.Literal)
	if name == "" {
		return nil, wrap(newParseErr("name cannot be blank", p.current))
	}
	call.Name = name

	if err := p.wantPeek(token.T_LPAREN); err != nil {
		return nil, wrap(err)
	}
	p.advance()
	call.LPar = p.current
	p.advance()
	call.Args = []ast.Expression{}
	if p.current.Type == token.T_RPAREN {
		call.Last = p.current
		return call, nil
	}

	for {
		if p.next.Type == token.T_EOF {
			break
		}
		arg, err := p.parseBlockExpression()
		if err != nil {
			return nil, wrap(err)
		}
		call.Args = append(call.Args, arg)

		if p.next.Type != token.T_COMMA {
			break
		}
		p.advance()
		p.advance()
	}

	p.advance()
	call.Last = p.current
	return call, nil
}

// parseOverExpression - looks like:
//
//	over Context<Expression>
func (p *Parser) parseOverExpression() (ast.Expression, error) {
	p.trace.trace("OverExpression")
	defer p.trace.untrace("OverExpression")

	overExp := ast.OverExpression{Token: p.current}
	p.advance()

	ctx, err := p.parseExpression()
	if err != nil {
		return nil, fmt.Errorf("parseOverExpression: %w", err)
	}
	overExp.Context = ctx

	if p.next.Type != token.T_ALIAS {
		return overExp, nil
	}

	p.advance()
	alias, err := p.parseAliasExpression()
	if err != nil {
		return nil, fmt.Errorf("parseOverExpression: %w", err)
	}
	overExp.HasAlias = true
	overExp.Alias = alias.(ast.AliasExpression)
	return overExp, nil
}

// parseWhereExpression - looks like:
//
//	where Condition<ast.Expression>
func (p *Parser) parseWhereExpression() (ast.Expression, error) {
	p.trace.trace("whereExpression")
	defer p.trace.untrace("whereExpression")

	whereExp := ast.WhereExpression{Token: p.current}
	p.advance()

	cnd, err := p.parseExpression()
	if err != nil {
		return nil, fmt.Errorf("paseOverExpression: %w", err)
	}
	whereExp.Condition = cnd

	return whereExp, nil
}

// parseOrderByExpression = looks like:
//
//	order by Expression<ast.Expression> [asc|desc]
//
// asc / desc is a single token, which is optional.
func (p *Parser) parseOrderByExpression() (ast.Expression, error) {
	p.trace.trace("OrderByExpression")
	defer p.trace.untrace("OrderByExpression")
	wrap := func(e error) error { return fmt.Errorf("parseOrderByExpression: %w", e) }

	orderByExp := ast.OrderByExpression{Token: p.current}
	if err := p.passIf(token.T_BY); err != nil {
		return nil, wrap(err)
	}

	exp, err := p.parseExpression()
	if err != nil {
		return nil, wrap(err)
	}
	orderByExp.Expression = exp

	if p.next.Type == token.T_ASC || p.next.Type == token.T_DESC {
		p.advance()
		asc := p.current
		orderByExp.Asc = &asc
	}
	return orderByExp, nil
}

// parseAliasExpression - looks like:
//
//	alias <Ident>
func (p *Parser) parseAliasExpression() (ast.Expression, error) {
	p.trace.trace("AliasExpression")
	defer p.trace.untrace("AliasExpression")
	wrap := func(e error) error { return fmt.Errorf("parseAliasExpression: %w", e) }

	aliasExp := ast.AliasExpression{Token: p.current}
	if err := p.wantPeek(token.T_IDENT); err != nil {
		return nil, wrap(err)
	}
	p.advance()

	ident, err := p.parseIdent()
	if err != nil {
		return nil, wrap(err)
	}
	aliasExp.Alias = ident.(ast.Ident)
	return aliasExp, nil
}

// parseSetExpression - looks like:
//
//	set <Ident>
func (p *Parser) parseSetExpression() (ast.Expression, error) {
	p.trace.trace("SetExpression")
	defer p.trace.untrace("SetExpression")
	wrap := func(e error) error { return fmt.Errorf("parseSetExpression: %w", e) }

	setExp := ast.SetExpression{Token: p.current}
	if err := p.wantPeek(token.T_IDENT); err != nil {
		return nil, wrap(err)
	}

	p.advance()
	name, err := p.parseIdent()
	if err != nil {
		return nil, wrap(err)
	}
	setExp.Name = name.(ast.Ident)
	return setExp, nil
}

// parseListLiteral - looks like:
// [ <StringLiteral>, <StringLiteral>, ... ]
//
// the list can also be empty as just "[]"
func (p *Parser) parseListLiteral() (ast.Expression, error) {
	p.trace.trace("ListLiteral")
	defer p.trace.untrace("ListLiteral")
	wrap := func(e error) error { return fmt.Errorf("parseListLiteral: %w", e) }

	if p.next.Type != token.T_STRING && p.next.Type != token.T_RBRACKET {
		msg := fmt.Sprintf("next token type: have %s, want %s or %s", p.next.Type, token.T_STRING, token.T_RBRACKET)
		return nil, wrap(newParseErr(msg, p.next))
	}

	listLit := ast.ListLiteral{
		Token:   p.current,
		Strings: []ast.StringLiteral{},
	}

	for p.next.Type == token.T_STRING {
		p.advance()

		str, err := p.parseStringLiteral()
		if err != nil {
			return nil, wrap(err)
		}
		listLit.Strings = append(listLit.Strings, str.(ast.StringLiteral))

		if p.next.Type != token.T_COMMA {
			break
		}
		p.advance()
	}

	if err := p.wantPeek(token.T_RBRACKET); err != nil {
		return nil, wrap(err)
	}

	p.advance()
	listLit.RBracket = p.current
	return listLit, nil
}

// parseInExpression - looks like:
//
//	left<Expression> in <ListLiteral> | <SetExpression>
func (p *Parser) parseInExpression(left ast.Expression) (ast.Expression, error) {
	p.trace.trace("InExpression")
	defer p.trace.untrace("InExpression")
	wrap := func(e error) error { return fmt.Errorf("parseInExpression: %w", e) }

	if left == nil {
		return nil, wrap(newParseErr("left expression is nil", p.current))
	}

	inExpression := ast.InExpression{
		Token: p.current,
		Left:  left,
	}
	var list ast.Expression
	var err error

	if p.next.Type == token.T_LBRACKET {
		p.advance()
		list, err = p.parseListLiteral()

	} else if p.next.Type == token.T_SET {
		p.advance()
		list, err = p.parseSetExpression()

	} else {
		msg := fmt.Sprintf("next token type: have %s, want %s or %s", p.next.Type, token.T_LBRACKET, token.T_SET)
		err = newParseErr(msg, p.next)
	}
	if err != nil {
		return nil, wrap(err)
	}
	inExpression.List = list
	return inExpression, nil
}

func (p *Parser) parseDateLiteral() (ast.Expression, error) {
	p.trace.trace("DateLiteral")
	defer p.trace.untrace("DateLiteral")

	date, err := wdate.ParseDate(p.current.Literal)
	if err != nil {
		return nil, fmt.Errorf("parseDateLiteral: error parsing date value: %w", err)
	}
	return ast.DateLiteral{Token: p.current, Time: date}, nil
}

func (p *Parser) parseTimeLiteral() (ast.Expression, error) {
	p.trace.trace("TimeLiteral")
	defer p.trace.untrace("TimeLiteral")

	tVal, err := wdate.ParseTime(p.current.Literal)
	if err != nil {
		return nil, fmt.Errorf("parseTimeLiteral: error parsing time value: %w", err)
	}
	return ast.TimeLiteral{Token: p.current, Time: tVal}, nil
}
