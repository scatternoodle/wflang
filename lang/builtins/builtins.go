package builtins

type Builtin struct {
	Name string
}

func Builtins() map[string]Builtin {
	return map[string]Builtin{
		// control
		"if": {"if"},

		// maths
		"min": {"min"},
	}
}
