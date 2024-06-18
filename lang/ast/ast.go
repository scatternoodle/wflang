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

// MacroExpression brings the scope of a Macro into a formula. Macros are the
// closest thing that WFLang has to user-defined functions.
type MacroExpression struct {
	token.Token
	Name    Ident
	Params  []Expression // TODO - check - how expressive are we allowed to be with Macro params?
	RDollar token.Token
}

func (m MacroExpression) ExpressionNode()      {}
func (m MacroExpression) TokenLiteral() string { return m.Token.Literal }

func (m MacroExpression) String() string {
	var out strings.Builder

	out.WriteString("$" + m.Name.String() + "(")
	for i, p := range m.Params {
		out.WriteString(p.String())
		if i < len(m.Params)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")$")
	return out.String()
}

func (m MacroExpression) Pos() (start, end token.Pos) {
	return m.Token.StartPos, m.RDollar.EndPos
}

// FunctionCall - to the Parser, a function call is simply an Ident and an group of
// arguments, all of which are BlockExpressions.
type FunctionCall struct {
	token.Token
	Name   string
	Args   []Expression
	RParen token.Token
}

func (f FunctionCall) ExpressionNode()      {}
func (f FunctionCall) TokenLiteral() string { return f.Token.Literal }

func (f FunctionCall) String() string {
	var out strings.Builder
	out.WriteString(f.Name + "(")
	for i, arg := range f.Args {
		out.WriteString(arg.String())
		if i == len(f.Args)-1 {
			out.WriteString(",")
		}
	}
	out.WriteString(")")
	return out.String()
}

func (f FunctionCall) Pos() (start, end token.Pos) {
	return f.Token.StartPos, f.RParen.EndPos
}

// OverExpression represents the "over" keyword, which occurs as the first subExpression
// of all summary functions. It prefaces the summary context of the function call
// (e.g. day, period.end)
type OverExpression struct {
	token.Token
	Context Expression
}

func (o OverExpression) ExpressionNode()      {}
func (o OverExpression) TokenLiteral() string { return o.Token.Literal }

func (o OverExpression) String() string { return "over " + o.Context.String() }

func (o OverExpression) Pos() (start, end token.Pos) {
	start = o.Token.StartPos
	_, end = o.Context.Pos()
	return start, end
}

// WhereExpression - pretty much what it sounds like, very like SQL WHERE and occurs
// mostly in summary functions.
type WhereExpression struct {
	token.Token
	Condition Expression
}

func (w WhereExpression) ExpressionNode()      {}
func (w WhereExpression) TokenLiteral() string { return w.Token.Literal }
func (w WhereExpression) String() string       { return "where " + w.Condition.String() }

func (w WhereExpression) Pos() (start, end token.Pos) {
	start = w.Token.StartPos
	_, end = w.Condition.Pos()
	return start, end
}

// OrderByExpression - used in some summary functions to order results.
type OrderByExpression struct {
	token.Token
	Expression
	Asc *token.Token // the keyword "asc" or "desc". This is optional, hence the pointer.
}

func (o OrderByExpression) ExpressionNode()      {}
func (o OrderByExpression) TokenLiteral() string { return o.Token.Literal }

func (o OrderByExpression) String() string {
	return "order by " + o.Expression.String() + o.Ascending()
}

func (o OrderByExpression) Pos() (start, end token.Pos) {
	start = o.Token.StartPos
	if o.Asc != nil {
		end = o.Asc.EndPos
	} else {
		_, end = o.Expression.Pos()
	}
	return start, end
}

// Ascending returns the literal of the OrderByExpression's "asc/desc" clause, or
// empty string if not present.
func (o OrderByExpression) Ascending() string {
	if o.Asc != nil {
		return o.Asc.Literal
	}
	return ""
}

// AliasExpression - used in summary functions to assign a name to the iteration
// context of the function.
type AliasExpression struct {
	token.Token
	Alias Ident
}

func (a AliasExpression) ExpressionNode()      {}
func (a AliasExpression) TokenLiteral() string { return a.Token.Literal }
func (a AliasExpression) String() string       { return "alias " + a.Alias.String() }

func (a AliasExpression) Pos() (start, end token.Pos) {
	start = a.Token.StartPos
	_, end = a.Alias.Pos()
	return start, end
}

// InExpression - a very limited in expression that can check if certain items are
// in a list represented by either an array of string literals (ListLiteral) or
// a SetExpression referencing a policy set Ident.
type InExpression struct {
	token.Token
	List Expression
}

func (i InExpression) ExpressionNode()      {}
func (i InExpression) TokenLiteral() string { return i.Token.Literal }
func (i InExpression) String() string       { return "in " + i.List.String() }

func (i InExpression) Pos() (start, end token.Pos) {
	start = i.Token.StartPos
	_, end = i.List.Pos()
	return start, end
}

// ListLiteral is a list of string literals, used exclusively for InExpressions.
type ListLiteral struct {
	token.Token
	Strings  []StringLiteral
	RBracket token.Token
}

func (l ListLiteral) ExpressionNode()             {}
func (l ListLiteral) TokenLiteral() string        { return l.Token.Literal }
func (l ListLiteral) Pos() (start, end token.Pos) { return l.Token.StartPos, l.RBracket.EndPos }

func (l ListLiteral) String() string {
	var out strings.Builder
	out.WriteByte('[')

	for i, str := range l.Strings {
		out.WriteString(str.String())
		if i >= len(l.Strings)-1 {
			out.WriteString(", ")
		}
	}

	out.WriteByte(']')
	return out.String()
}

// SetExpression is a reference to a policy set Ident, used exclusively for InExpressions.
type SetExpression struct {
	token.Token
	Name Ident
}

func (s SetExpression) ExpressionNode()      {}
func (s SetExpression) TokenLiteral() string { return s.Token.Literal }
func (s SetExpression) String() string       { return "set " + s.Name.String() }
func (s SetExpression) Pos() (start, end token.Pos) {
	start = s.Token.StartPos
	_, end = s.Name.Pos()
	return start, end
}
