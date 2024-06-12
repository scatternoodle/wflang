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
	StatementNode()
}

// Expression must be implemented by all expression nodes in the AST. Aside from
// conferring the Node interface, it simply provides a blank method to distinguish
// expressions from statements.
type Expression interface {
	Node
	ExpressionNode()
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
		out.WriteString(s.String() + "\n")
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
	token.Token // the first token of the expression
	Expression  Expression
}

func (e ExpressionStatement) StatementNode()       {}
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
	token.Token // expected to be token.T_VAR type
	Name        Ident
	Value       Expression
}

func (v VarStatement) StatementNode()       {}
func (v VarStatement) TokenLiteral() string { return v.Token.Literal }
func (v VarStatement) String() string {
	return fmt.Sprintf("var %s = %s;", v.Name.String(), v.Value.String())
}

// Pos returns the StartPos of the var token, and the EndPos of the Value expression.
func (v VarStatement) Pos() (start, end token.Pos) {
	_, end = v.Value.Pos()
	return v.Token.StartPos, end
}

// LineCommentStatement is a statement that represents a single line comment, where
// the token literal is the comment itself (including the leading "//").
type LineCommentStatement struct {
	token.Token
}

func (l LineCommentStatement) StatementNode()       {}
func (l LineCommentStatement) TokenLiteral() string { return l.Token.Literal }
func (l LineCommentStatement) String() string       { return l.Token.Literal }
func (l LineCommentStatement) Pos() (start, end token.Pos) {
	return l.Token.StartPos, l.Token.EndPos
}

// BlockCommentStatement is a statement that represents a block comment, where the
// token literal is the comment itself (including the leading "/*" and trailing "*/").
type BlockCommentStatement struct {
	token.Token
}

func (b BlockCommentStatement) StatementNode()       {}
func (b BlockCommentStatement) TokenLiteral() string { return b.Token.Literal }
func (b BlockCommentStatement) String() string       { return b.Token.Literal }
func (b BlockCommentStatement) Pos() (start, end token.Pos) {
	return b.Token.StartPos, b.Token.EndPos
}

// Ident is an expression that represents a variable name. It is used in
// VarStatements to declare variables, and in other expressions to reference them.
type Ident struct {
	token.Token
	Value string
}

func (i Ident) ExpressionNode()             {}
func (i Ident) TokenLiteral() string        { return i.Token.Literal }
func (i Ident) String() string              { return i.Value }
func (i Ident) Pos() (start, end token.Pos) { return i.Token.StartPos, i.Token.EndPos }

// NumberLiteral is an expression that represents a number literal. The value
// is stored as a float64 but can also hold an integer type (WFLang has one number
// type).
type NumberLiteral struct {
	token.Token
	Value float64
}

func (n NumberLiteral) ExpressionNode()      {}
func (n NumberLiteral) TokenLiteral() string { return n.Token.Literal }
func (n NumberLiteral) String() string       { return fmt.Sprint(n.Value) }
func (n NumberLiteral) Pos() (start, end token.Pos) {
	return n.Token.StartPos, n.Token.EndPos
}

// PrefixExpression is an expression prefixed by an operator. It stores both the operator
// token and the Expression it prefixes.
type PrefixExpression struct {
	Token  token.Token
	Prefix string
	Right  Expression
}

func (p PrefixExpression) ExpressionNode()      {}
func (p PrefixExpression) TokenLiteral() string { return p.Token.Literal }

func (p PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", p.Prefix, p.Right.String())
}

// Pos returns the StartPos of the prefix token, and the EndPos of the Right expression.
func (p PrefixExpression) Pos() (start, end token.Pos) {
	start = p.Token.StartPos
	if p.Right == nil {
		end = start
		return start, end
	}
	_, end = p.Right.Pos()
	return start, end
}

// InfixExpression is an expression that contains an operator between two expressions.
// It stores both the left and right expressions, as well as the operator token.
type InfixExpression struct {
	token.Token
	Left  Expression
	Infix string
	Right Expression
}

func (i InfixExpression) ExpressionNode()      {}
func (i InfixExpression) TokenLiteral() string { return i.Token.Literal }

func (i InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", i.Left.String(), i.Infix, i.Right.String())
}

// Pos returns the StartPos of the Left expression, and the EndPos of the Right expression.
// If either left or right are nil, returns the start and end of the infix token.
func (i InfixExpression) Pos() (start, end token.Pos) {
	if i.Left == nil || i.Right == nil {
		return i.Token.StartPos, i.Token.EndPos // best we can reasonably do - something has gone pretty wrong.
	}

	start, _ = i.Left.Pos()
	_, end = i.Right.Pos()
	return start, end
}

// StringLiteral is an expression that represents a string literal. The string
// itself is stored in the token literal.
type StringLiteral struct {
	token.Token
}

func (s StringLiteral) ExpressionNode()      {}
func (s StringLiteral) TokenLiteral() string { return s.Token.Literal }
func (s StringLiteral) String() string       { return `"` + s.Token.Literal + `"` }
func (s StringLiteral) Pos() (start, end token.Pos) {
	return s.Token.StartPos, s.Token.EndPos
}

// BooleanLiteral is an expression that represents a boolean literal.
type BooleanLiteral struct {
	token.Token
	Value bool
}

func (b BooleanLiteral) ExpressionNode()      {}
func (b BooleanLiteral) TokenLiteral() string { return b.Token.Literal }
func (b BooleanLiteral) String() string       { return fmt.Sprint(b.Value) }
func (b BooleanLiteral) Pos() (start, end token.Pos) {
	return b.Token.StartPos, b.Token.EndPos
}

// IfExpression comprises a three-subexpression structure::
//
//	if( <condition>
//	  , <consequence>
//	  , <alternative> )
//
// All three expressions are mandatory.
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence Expression
	Alternative Expression
}

func (i IfExpression) ExpressionNode()      {}
func (i IfExpression) TokenLiteral() string { return i.Token.Literal }

func (i IfExpression) String() string {
	var out strings.Builder
	out.WriteString("if( ")
	out.WriteString(i.Condition.String() + ", ")
	out.WriteString(i.Consequence.String() + ", ")
	out.WriteString(i.Alternative.String() + " )")
	return out.String()
}

// Pos returns the StartPos of the if token, and the EndPos of the Alternative expression.
func (i IfExpression) Pos() (start, end token.Pos) {
	start = i.Token.StartPos
	_, end = i.Alternative.Pos()
	return start, end
}

// BlockExpression is a group of 0-n number of VarStatements followed
// by a single Expression. The VarStatements are optional but the Expression is
// mandatory, there can only be one, and it must be the last member of the group.
type BlockExpression struct {
	token.Token
	Vars  []VarStatement
	Value Expression
}

func (b BlockExpression) ExpressionNode()      {}
func (b BlockExpression) TokenLiteral() string { return b.Token.Literal }

func (b BlockExpression) String() string {
	var out strings.Builder
	out.WriteString("( ")
	for _, v := range b.Vars {
		out.WriteString("\t" + v.String() + "\n")
	}
	out.WriteString(b.Value.String() + " )")
	return out.String()
}

func (b BlockExpression) Pos() (start, end token.Pos) {
	_, end = b.Value.Pos()
	return b.StartPos, end
}

type ParenExpression struct {
	token.Token
	Inner  Expression
	RParen token.Token
}

func (p ParenExpression) ExpressionNode()             {}
func (p ParenExpression) TokenLiteral() string        { return p.Token.Literal }
func (p ParenExpression) String() string              { return "(" + p.Inner.String() + ")" }
func (p ParenExpression) Pos() (start, end token.Pos) { return p.Token.StartPos, p.RParen.EndPos }

// BlankExpression is an empty expression. Token is usually the preceding token.
type BlankExpression struct {
	token.Token
}

func (b BlankExpression) ExpressionNode()             {}
func (b BlankExpression) TokenLiteral() string        { return b.Token.Literal }
func (b BlankExpression) String() string              { return "" }
func (b BlankExpression) Pos() (start, end token.Pos) { return b.Token.StartPos, b.Token.EndPos }
