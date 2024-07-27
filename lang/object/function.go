package object

type Function struct {
	Name       string
	ReturnType Type
	Params     []Param
}

type Param struct {
	Name     string
	Types    pT // permitted types, can be many for some params
	Optional bool
	List     bool // function call can have N number of this param
	PairA    bool // is 1st in pair of params
	PairB    bool // is 2nd in pair of params
}

type pT []Type
