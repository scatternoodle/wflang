package analysis

import "github.com/scatternoodle/wflang/lang/types"

type Object struct {
	Type    types.BaseType
	Context Environment
}

type Environment map[string]Object
