package lexer

import "github.com/scatternoodle/wflang/lang/token"

func keywords() map[string]token.Type {
	// We define keywords here, including builtin functions.
	//
	// Global Variables are not included - those are handled by the parser along
	// with user-defined variable idents.
	//
	// Access methods are also delegated to the parser, as they are not keywords
	// and require context to be resolved.
	return map[string]token.Type{
		"var": token.T_VAR,

		// word literals
		"null":  token.T_NULL,
		"true":  token.T_TRUE,
		"false": token.T_FALSE,

		// logical
		// WFLang allows the actual words "and" and "or" to be used as logical operators.
		// TODO - check if case sensitive.
		// TODO - shall we just refuse to recognize this? This is an opinionated
		// tool, after all.
		"and": token.T_AND,
		"or":  token.T_OR,
		"not": token.T_BANG,

		// semantic keywords
		"alias": token.T_ALIAS,
		"over":  token.T_OVER,
		"where": token.T_WHERE,
		"order": token.T_ORDER,
		"by":    token.T_BY,
		"in":    token.T_IN,
		"set":   token.T_SET,
	}
}
