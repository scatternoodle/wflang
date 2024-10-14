package object

import "github.com/scatternoodle/wflang/wflang/types"

type Function struct {
	Name       string
	ReturnType types.Type
	Params     []Param
}

type pTypes []types.Type

type Param struct {
	Name     string
	Types    []types.Type // permitted types, can be many for some params
	Optional bool
	List     bool // function call can have N number of this param
	PairA    bool // is 1st in pair of params
	PairB    bool // is 2nd in pair of params
}

// Variadic returns true if the Function is variadic; in that one or more parameters
// can take unlimited arguments - e.g. min(args...)
func (f Function) Variadic() bool {
	for _, param := range f.Params {
		if param.List {
			return true
		}
	}
	return false
}
