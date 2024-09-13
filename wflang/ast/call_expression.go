package ast

import (
	"strings"

	"github.com/scatternoodle/wflang/wflang/token"
)

// CallExpression is implemented by any expression in WFlang that represents a
// call to a builtin function, type method, invocation or macro.
type CallExpression interface {
	LParen() token.Pos
	RParen() token.Pos
	Params() []Expression
}

// BuiltinCall represents a call to a builtin function in WFLang and implements
// the CallExpression interface.
type BuiltinCall struct {
	token.Token
	Name string
	Args []Expression
	LPar token.Token
	RPar token.Token
}

func (f BuiltinCall) ExpressionNode()      {}
func (f BuiltinCall) TokenLiteral() string { return f.Token.Literal }

func (f BuiltinCall) String() string {
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

func (f BuiltinCall) Pos() (start, end token.Pos) {
	return f.Token.StartPos, f.RPar.EndPos
}

func (b BuiltinCall) LParen() token.Pos { return b.LPar.StartPos }
func (b BuiltinCall) RParen() token.Pos { return b.RPar.StartPos }

// MacroExpression brings the scope of a Macro into a formula. Macros are the
// closest thing that WFLang has to user-defined functions. Implements the
// CallExpression interface.
type MacroExpression struct {
	token.Token
	Name    Ident
	Args    []Expression // TODO - check - how expressive are we allowed to be with Macro params?
	RDollar token.Token
	LPar    token.Token
	RPar    token.Token
}

func (m MacroExpression) ExpressionNode()      {}
func (m MacroExpression) TokenLiteral() string { return m.Token.Literal }

func (m MacroExpression) String() string {
	var out strings.Builder

	out.WriteString("$" + m.Name.String() + "(")
	for i, p := range m.Args {
		out.WriteString(p.String())
		if i < len(m.Args)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")$")
	return out.String()
}

func (m MacroExpression) Pos() (start, end token.Pos) {
	return m.Token.StartPos, m.RDollar.EndPos
}

func (m MacroExpression) LParen() token.Pos { return m.LPar.StartPos }
func (m MacroExpression) RParen() token.Pos { return m.RPar.StartPos }
