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
		"Returns the smallest of its arguments. It takes a list of arguments as long as you like, which can be any" +
		" expression that evaluates to a number.\n\n" +
		"@param `args: number...` - list of numbers to compare\n\n"

	Max string = CodeBlockStart +
		"max(args: number...)\n" +
		CodeBlockEnd +
		"### Max\n" +
		"Returns the largest of its arguments. It takes a list of arguments as long as you like, which can be any" +
		" expression that evaluates to a number.\n\n" +
		"@param `args: number...` - list of numbers to compare\n\n"

	Contains string = CodeBlockStart +
		"contains(string1: string, string2: string)\n" +
		CodeBlockEnd +
		"### Contains\n" +
		"Returns true if `string2` is a substring of `string1`.\n\n" +
		"@param `string1: string` - the string to search in\n\n" +
		"@param `string2: string` - the string to search for\n\n"

	Sum string = CodeBlockStart +
		"sum( by interval: day|week|period| over range: day|week|period|dateRange alias aliasName: string\n" +
		"   , expression: number \n" +
		"  [, where condition: boolean] )\n" +
		CodeBlockEnd +
		"### Sum\n" +
		"Calculates the sum of a numeric expression repeatedly over `range`.\n\n" +
		"@param `interval: day|week|period` - the time period to group by\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to sum\n\n" +
		"@param `aliasName: string` - alias for the interval currently being evaluated\n\n" +
		"@param `expression: number` - this is what will be summed for each qualifying interval\n\n" +
		"@param `condition?: boolean` - only intervals where this expression returns true will be evaluated. If omitted, all intervals will be summed\n\n"

	Count string = CodeBlockStart +
		"count( by interval: day|week|period over range: day|week|period|dateRange alias aliasName: string\n" +
		"     , where condition: boolean )\n" +
		CodeBlockEnd +
		"### Count\n" +
		"Calculates a total count of qualifying intervals across `range`.\n\n" +
		"@param `interval: day|week|period` - the time period to group by\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName: string` - alias for the interval currently being evaluated\n\n" +
		"@param `condition: boolean` - only intervals where this expression returns true will be counted.\n\n"

	If string = CodeBlockStart +
		"if( condition: boolean\n" +
		"  , then: any\n" +
		"  , else: any )\n" +
		CodeBlockEnd +
		"### If\n" +
		"If expressions in WFLang comprise a check condition, a consequence and an alternative." +
		" Unlike most languages, in WFLang the `then` and `else` expressions are mandatory, so all If expressions" +
		" are in fact If-Else expressions.\n\n" +
		"@param condition: boolean - the condition to evaluate\n\n" +
		"@param `then`: any - the expression to evaluate if the condition is true\n\n" +
		"@param `else`: any - the expression to evaluate if the condition is false\n\n"

	SumTime string = CodeBlockStart +
		"sumTime( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"       , expression: number\n" +
		"      [, where expression?: bool] )\n" +
		CodeBlockEnd +
		"### SumTime\n" +
		"Calculates the sum of a numeric expression over a range of time records.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to sum\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `expression: number` - numeric expression - this is what will be summed for each qualifying slice\n\n" +
		"@param `condition?: boolean` - if used, only slices where this expression returns true will be evaluated\n\n"

	CountTime string = CodeBlockStart +
		"countTime( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"        [, where expression?: bool] )\n" +
		CodeBlockEnd +
		"### CountTime\n" +
		"Calculates the count of time records in a range.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `expression?: boolean` - if used, only slices where this expression returns true will be counted\n\n"

	FindFirstTime string = CodeBlockStart +
		"findFirstTime( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"             , where condition: boolean\n" +
		"             , order by ordering: string|number|date|dateTime )\n" +
		CodeBlockEnd +
		"### FindFirstTime\n" +
		"Returns the first time record that meets `condition`, ordered by `ordering`.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n" +
		"@param `ordering: string|number|date|dateTime` - the value to order by\n\n"

	SumSchedule string = CodeBlockStart +
		"sumSchedule( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"           , expression: number\n" +
		"          [, where expression?: bool] )\n" +
		CodeBlockEnd +
		"### SumSchedule\n" +
		"Calculates the sum of a numeric expression over a range of schedule records.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to sum\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `expression: number` - numeric expression - this is what will be summed for each qualifying slice\n\n" +
		"@param `condition?: boolean` - if used, only slices where this expression returns true will be evaluated\n\n"

	CountSchedule string = CodeBlockStart +
		"countSchedule( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"            [, where expression?: bool] )\n" +
		CodeBlockEnd +
		"### CountSchedule\n" +
		"Calculates the count of schedule records in a range.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `expression?: boolean` - if used, only slices where this expression returns true will be counted\n\n"

	FindFirstSchedule string = CodeBlockStart +
		"findFirstSchedule( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                 , where condition: boolean\n" +
		"                 , order by ordering: string|number|date|dateTime )\n" +
		CodeBlockEnd +
		"### FindFirstSchedule\n" +
		"Returns the first schedule record that meets `condition`, ordered by `ordering`.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n" +
		"@param `ordering: string|number|date|dateTime` - the value to order by\n\n"

	CountException string = CodeBlockStart +
		"countException( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"             [, where expression?: bool] )\n" +
		CodeBlockEnd +
		"### CountException\n" +
		"Calculates the count of exception records in a range.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the exception currently being evaluated\n\n" +
		"@param `expression?: boolean` - if used, only exceptions where this expression returns true will be counted\n\n"

	FindFirstTorDetail string = CodeBlockStart +
		"findFirstTorDetail( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                  , where condition: boolean\n" +
		"                  , order by ordering: string|number|date|dateTime )\n" +
		CodeBlockEnd +
		"### FindFirstTorDetail\n" +
		"Returns the first TOR detail record that meets `condition`, ordered by `ordering`.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the record currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n" +
		"@param `ordering: string|number|date|dateTime` - the value to order by\n\n"

	FindFirstDayForward string = CodeBlockStart +
		"findFirstDayForward( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                   , where condition: boolean )\n" +
		CodeBlockEnd +
		"### FindFirstDayForward\n" +
		"Returns the first date that meets `condition`, going from the start to the end of `range`, excluding the start date.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the day currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n"
)