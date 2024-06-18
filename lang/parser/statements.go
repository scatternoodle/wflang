package parser

import (
	"fmt"

	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/token"
)

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
	exp, err := p.parseExpression()
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
	exp, err := p.parseExpression()
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

func (p *Parser) parseLineCommentStatement() ast.LineCommentStatement {
	p.trace.trace("LineCommentStatement")
	defer p.trace.untrace("LineCommentStatement")

	return ast.LineCommentStatement{Token: p.current}
}

func (p *Parser) parseBlockCommentStatement() ast.BlockCommentStatement {
	p.trace.trace("BlockCommentStatement")
	defer p.trace.untrace("BlockCommentStatement")

	return ast.BlockCommentStatement{Token: p.current}
}
