package analysis

import (
	"github.com/scatternoodle/wflang/lang/ast"
	"github.com/scatternoodle/wflang/lang/builtins"
	"github.com/scatternoodle/wflang/lang/lexer"
)

func Evaluate(program *ast.AST) {}

// isReserved returns true if string is a reserved language keyword.
func IsReserved(s string) bool {
	_, isKeyword := lexer.Keyword(s)
	_, isBuiltin := builtins.Builtins()[s]
	return isKeyword || isBuiltin
}
