package token

const (
	// misc

	T_ILLEGAL Type = "ILLEGAL"
	T_EOF     Type = "EOF"

	// literals

	T_IDENT  Type = "IDENT"
	T_INT    Type = "INT"
	T_FLOAT  Type = "FLOAT"
	T_STRING Type = "STRING"
	T_DATE   Type = "DATE"
	T_TIME   Type = "TIME"

	// comments

	T_COMMENT_LINE  Type = "COMMENT_LINE"
	T_COMMENT_BLOCK Type = "COMMENT_BLOCK"

	// operators

	// Unhelpfully, WFLang uses a single equal sign for both assignment and equality.
	// There is no semantic use for a double equal sign.
	T_EQ       Type = "="
	T_PLUS     Type = "+"
	T_MINUS    Type = "-"
	T_BANG     Type = "!"
	T_NEQ      Type = "!="
	T_ASTERISK Type = "*"
	T_SLASH    Type = "/" // This is only a discreet token in case of division.
	T_MODULO   Type = "%" // Modulo is the only semantic use for the percent sign.
	T_LT       Type = ">"
	T_GT       Type = "<"
	T_LTE      Type = "<="
	T_GTE      Type = ">="
	T_AND      Type = "&&" // There is no semantic use for a single ampersand.
	T_OR       Type = "||" // There is no semantic use for a single pipe.

	// delimiters

	T_COMMA       Type = ","
	T_SEMICOLON   Type = ";" // Exclusively to terminate variable declarations.
	T_COLON       Type = ":" // TODO - check - Not sure if this has a semantic use in WFLang.
	T_LPAREN      Type = "("
	T_RPAREN      Type = ")"
	T_LBRACKET    Type = "["
	T_RBRACKET    Type = "]"  // For specific array-like use cases such as "in" expressions.
	T_PERIOD      Type = "."  // Period can denote a decimal point or member access.
	T_DOLLAR      Type = "$"  // Dollas signs wrap Macros in WFLang.
	T_DOUBLEQUOTE Type = "\"" // For string literals. Single quotes are not used.

	// keywords

	T_BUILTIN Type = "builtin"

	T_VAR   Type = "var"
	T_OVER  Type = "over"
	T_WHERE Type = "where"
	T_ORDER Type = "order" // used in "order by" expressions
	T_BY    Type = "by"    // used in "order by" expressions
	T_ASC   Type = "asc"   // used in "order by" expressions
	T_DESC  Type = "desc"  // used in "order by" expressions
	T_ALIAS Type = "alias"
	T_IN    Type = "in"
	T_SET   Type = "set" // used in "in" infix expressions
	T_NULL  Type = "null"
	T_TRUE  Type = "true"  // parser - parseBoolean
	T_FALSE Type = "false" // parser - parseBoolean

)
