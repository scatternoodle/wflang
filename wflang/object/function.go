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
