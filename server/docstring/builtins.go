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
				Label: [2]int{4, 19},
				Type:  "number",
				Desc:  "list of numbers to compare",
			},
		},
	},

	"max": {
		Name:      "Max",
		Signature: "max(args: number...)",
		Returns:   "number",
		Desc: "Returns the largest of its arguments. It takes a list of arguments as long " +
			"as you like, which can be any expression that evaluates to a number.",
		Params: []*ParamDoc{
			{
				Name:  "args",
				Label: [2]int{4, 19},
				Type:  "number",
				Desc:  "list of numbers to compare",
			},
		},
	},
}
