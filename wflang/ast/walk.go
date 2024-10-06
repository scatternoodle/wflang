package ast

import "fmt"

// Visitor is a function that is called recursively while walking the AST. Used
// to enclose logic for static analysis such as type filtering and building maps.
// To stop walking, Visit should return nil.
type Visitor interface {
	Visit(n Node) Visitor
}

type inspector func(Node) bool

func (f inspector) Visit(node Node) Visitor {
	if f(node) {
		return f
	}
	return nil
}

// Inspect provides a closure around Walk that allows you to walk the AST with
// an inspector function f, which is transformed into a Visitor and passed to the
// Walk function.
func Inspect(n Node, f func(Node) bool) {
	Walk(inspector(f), n)
}

// Walk traverses an AST in depth-first order: It starts by calling
// v.Visit(node); node must not be nil. If the visitor w returned by
// v.Visit(node) is not nil, Walk is invoked recursively with visitor
// w for each of the non-nil children of node, followed by a call of
// w.Visit(nil).
func Walk(v Visitor, node Node) {
	if v = v.Visit(node); v == nil {
		return
	}

	switch n := node.(type) {
	case *AST:
		if n.Statements != nil {
			walkList(v, n.Statements)
		}
	case ExpressionStatement:
		if n.Expression != nil {
			Walk(v, n.Expression)
		}
	case VarStatement:
		if n.Value != nil {
			Walk(v, n.Value)
		}
	case PrefixExpression:
		if n.Right != nil {
			Walk(v, n.Right)
		}
	case InfixExpression:
		if n.Left != nil {
			Walk(v, n.Left)
		}
		if n.Right != nil {
			Walk(v, n.Right)
		}
	case BlockExpression:
		if len(n.Vars) != 0 {
			walkList(v, n.Vars)
		}
		if n.Value != nil {
			Walk(v, n.Value)
		}
	case ParenExpression:
		if n.Inner != nil {
			Walk(v, n.Inner)
		}
	case MacroExpression:
		if len(n.Args) > 0 {
			walkList(v, n.Args)
		}
	case BuiltinCall:
		if len(n.Args) > 0 {
			walkList(v, n.Args)
		}
	case OverExpression:
		if n.Context != nil {
			Walk(v, n.Context)
		}
		if n.HasAlias {
			Walk(v, n.Alias)
		}
	case WhereExpression:
		if n.Condition != nil {
			Walk(v, n.Condition)
		}
	case OrderByExpression:
		if n.Expression != nil {
			Walk(v, n.Expression)
		}
	case InExpression:
		if n.Left != nil {
			Walk(v, n.Left)
		}
		if n.List != nil {
			Walk(v, n.List)
		}
	case Ident, LineCommentStatement, BlockCommentStatement, NumberLiteral,
		StringLiteral, BooleanLiteral, BlankExpression, ListLiteral,
		SetExpression, DateLiteral, TimeLiteral:
		// nothing to do as these do not have child nodes
	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
	}
}

func walkList[N Node](v Visitor, list []N) {
	for _, n := range list {
		Walk(v, n)
	}
}
