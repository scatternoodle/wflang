package hovertext

const (
	CodeBlockStart string = "```wflang\n"
	CodeBlockEnd   string = "```\n"
)

const (
	Min string = CodeBlockStart +
		"min(args: number...)\n" +
		CodeBlockEnd +
		"### Min\n" +
		"Min returns the smallest of its arguments. It takes a list of arguments as long as you like, which can be any" +
		" expression that evaluates to a number.\n\n" +
		"@param `args: number...` - list of numbers to compare\n\n"

	Max string = CodeBlockStart +
		"max(args: number...)\n" +
		CodeBlockEnd +
		"### Max\n" +
		"Max returns the largest of its arguments. It takes a list of arguments as long as you like, which can be any" +
		" expression that evaluates to a number.\n\n" +
		"@param `args: number...` - list of numbers to compare\n\n"

	If string = CodeBlockStart +
		"if( condition: boolean\n" +
		"  , then: expression\n" +
		"  , else: expression )\n" +
		CodeBlockEnd +
		"### If\n" +
		"If expressions in WFLang comprise a check condition, a consequence and an alternative." +
		" Unlike most languages, in WFLang the `then` and `else` expressions are mandatory, so all If expressions" +
		" are in fact If-Else expressions.\n\n" +
		"@param condition: boolean - the condition to evaluate\n\n" +
		"@param `then`: expression - the expression to evaluate if the condition is true\n\n" +
		"@param `else`: expression - the expression to evaluate if the condition is false\n\n"

	SumTime string = CodeBlockStart +
		"sumTime( over range: timePeriod [alias aliasName?: string]\n" +
		"       , sumExpression: number\n" +
		"      [, where whereExpression: bool] )\n" +
		CodeBlockEnd +
		"### SumTime\n" +
		"SumTime calculates the sum of a numeric expression over a range of time records.\n\n" +
		"@param `range: timePeriod` - the time period over which to sum\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `sumExpression: number` - numeric expression - this is what will be summed for each qualifying slice\n\n" +
		"@param `whereExpression?: boolean` - if used, only slices where this expression returns true will be evaluated\n\n"
)
