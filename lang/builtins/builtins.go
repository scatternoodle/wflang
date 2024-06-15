package builtins

type Builtin struct {
	Name string
}

func Builtins() map[string]Builtin {
	return map[string]Builtin{
		"min": {"min"},
	}
}
