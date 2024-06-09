// AST (Abstract Syntax Tree) package defines the structure of the AST that the
// parser will generate from the token stream produced by the lexer. The AST is
// the last step for WFLang, which does not require an interpreter or compiler.
package ast

import (
	"fmt"
	"strings"

	"github.com/scatternoodle/wflang/lang/token"
)

// Node is the base interface that all parts of the AST must implement, including
// the AST itself.
type Node interface {
	fmt.Stringer
	token.Positional
	TokenLiteral() string
}

// Statement must be implemented by all statement nodes in the AST. Aside from
// conferring the Node interface, it simply provides a blank method to distinguish
// statements from expressions.
type Statement interface {
	Node
	statementNode()
}

// Expression must be implemented by all expression nodes in the AST. Aside from
// conferring the Node interface, it simply provides a blank method to distinguish
// expressions from statements.
type Expression interface {
	Node
	expressionNode()
}

// The AST struct is the root node of the AST. It contains a slice of Statement
// nodes, which are the top-level nodes of the AST. The AST struct itself implements
// the Node interface, but neither the Statement nor Expression interfaces.
type AST struct {
	Statements []Statement
}

func (a *AST) String() string {
	var out strings.Builder
	for _, s := range a.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (a *AST) TokenLiteral() string {
	if len(a.Statements) == 0 {
		return ""
	}
	return a.Statements[0].TokenLiteral()
}

// Pos returns the start position of the first statement and the end position of the
// last statement in the AST. If the AST has no statements, it returns two zero positions.
func (a *AST) Pos() (start, end token.Pos) {
	if len(a.Statements) == 0 {
		return token.Pos{}, token.Pos{}
	}

	start, _ = a.Statements[0].Pos()
	_, end = a.Statements[len(a.Statements)-1].Pos()
	return start, end
}

// ExpressionStatement is a statement that contains a single expression. It is the
// most common type of statement in WFLang, and in fact it could be said that all
// formulas, no matter how long and complex, are simply a single expression.
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (e ExpressionStatement) statementNode()       {}
func (e ExpressionStatement) TokenLiteral() string { return e.Token.Literal }

func (e ExpressionStatement) String() string {
	if e.Expression == nil {
		return ""
	}
	return e.Expression.String()
}

// Pos simply returns the position of the Expression within the Statement, by
// calling Expression.Pos().
func (e ExpressionStatement) Pos() (start, end token.Pos) {
	return e.Expression.Pos()
}

// VarStatement is a statement that declares a variable. It is a simple statement
// that consists of the token.T_VAR token, with an Identifier Expression containing
// the name of the variable.
type VarStatement struct {
	Token token.Token // expected to be token.T_VAR type
	Name  Identifier
	Value Expression
}

func (v VarStatement) statementNode() {}
func (v VarStatement) TokenLiteral() string {
	return v.Token.Literal
}

func (v VarStatement) String() string {
	return fmt.Sprintf("var %s = %s;", v.Name.String(), v.Value.String())
}

// Pos returns the StartPos of the var token, and the EndPos of the Value expression.
func (v VarStatement) Pos() (start, end token.Pos) {
	_, end = v.Value.Pos()
	return v.Token.StartPos, end
}

// Identifier is an expression that represents a variable name. It is used in
// VarStatements to declare variables, and in other expressions to reference them.
type Identifier struct {
	Token token.Token
	Value string
}

func (i Identifier) expressionNode()             {}
func (i Identifier) TokenLiteral() string        { return i.Token.Literal }
func (i Identifier) String() string              { return i.Value }
func (i Identifier) Pos() (start, end token.Pos) { return i.Token.StartPos, i.Token.EndPos }
