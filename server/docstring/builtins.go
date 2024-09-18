package docstring

func BuiltinDoc(name string) (*FunctionDoc, bool) {
	f, ok := builtinDocs[name]
	return &f, ok
}

var builtinDocs = map[string]FunctionDoc{
	"min": {
		Name:      "Min",
		Signature: "min(args: number...)",
		Returns:   "number",
		Desc: "Returns the smallest of its arguments. It takes a list of arguments as long " +
			"as you like, which can be any expression that evaluates to a number.",
		Params: []*ParamDoc{
			{
				Name:  "args",
				Label: [2]int{0, 14},
				Type:  "number",
				Desc:  "list of numbers to compare",
			},
		},
	},

	"contains": {
		Name:      "Contains",
		Signature: "contains(x: string, y: string)",
		Returns:   "boolean",
		Desc:      "Returns true if `y` is a substring of `x`.",
		Params: []*ParamDoc{
			{Name: "x", Label: [2]int{9, 9}, Type: "string", Desc: "the string to search in"},
			{Name: "y", Label: [2]int{20, 20}, Type: "string", Desc: "the string to search in"},
		},
	},
}
