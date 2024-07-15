package analysis

import (
	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/builtins"
	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/types"
)

type object struct {
	baseType    types.BaseType
	environment map[string]object
}

func Evaluate(ast.Node) {
}

// isReserved returns true if string is a reserved language keyword.
func IsReserved(s string) bool {
	_, isKeyword := lexer.Keyword(s)
	_, isBuiltin := builtins.Builtins()[s]
	return isKeyword || isBuiltin
}
