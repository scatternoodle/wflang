package hovertext

const (
	codeBlockStart   string = "```wflang\n"
	codeBlockReturns string = "\nreturns: "
	codeBlockEnd     string = "```\n"
)

const (
	Min string = codeBlockStart +
		"min(args: number...)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### Min\n" +
		"Returns the smallest of its arguments. It takes a list of arguments as long as you like, which can be any" +
		" expression that evaluates to a number.\n\n" +
		"@param `args: number...` - list of numbers to compare\n\n"

	Max string = codeBlockStart +
		"max(args: number...)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### Max\n" +
		"Returns the largest of its arguments. It takes a list of arguments as long as you like, which can be any" +
		" expression that evaluates to a number.\n\n" +
		"@param `args: number...` - list of numbers to compare\n\n"

	Contains string = codeBlockStart +
		"contains(string1: string, string2: string)\n" +
		codeBlockReturns + "boolean\n" +
		codeBlockEnd +
		"### Contains\n" +
		"Returns true if `string2` is a substring of `string1`.\n\n" +
		"@param `string1: string` - the string to search in\n\n" +
		"@param `string2: string` - the string to search for\n\n"

	Sum string = codeBlockStart +
		"sum( by interval: day|week|period| over range: day|week|period|dateRange alias aliasName: string\n" +
		"   , expression: number \n" +
		"  [, where condition: boolean] )\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### Sum\n" +
		"Calculates the sum of a numeric expression repeatedly over `range`.\n\n" +
		"@param `interval: day|week|period` - the time period to group by\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to sum\n\n" +
		"@param `aliasName: string` - alias for the interval currently being evaluated\n\n" +
		"@param `expression: number` - this is what will be summed for each qualifying interval\n\n" +
		"@param `condition?: boolean` - only intervals where this expression returns true will be evaluated. If omitted, all intervals will be summed\n\n"

	Count string = codeBlockStart +
		"count( by interval: day|week|period over range: day|week|period|dateRange alias aliasName: string\n" +
		"     , where condition: boolean )\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### Count\n" +
		"Calculates a total count of qualifying intervals across `range`.\n\n" +
		"@param `interval: day|week|period` - the time period to group by\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName: string` - alias for the interval currently being evaluated\n\n" +
		"@param `condition: boolean` - only intervals where this expression returns true will be counted.\n\n"

	If string = codeBlockStart +
		"if( condition: boolean\n" +
		"  , then: any\n" +
		"  , else: any )\n" +
		codeBlockReturns + "any\n" +
		codeBlockEnd +
		"### If\n" +
		"If expressions in WFLang comprise a check condition, a consequence and an alternative." +
		" Unlike most languages, in WFLang the `then` and `else` expressions are mandatory, so all If expressions" +
		" are in fact If-Else expressions.\n\n" +
		"@param condition: boolean - the condition to evaluate\n\n" +
		"@param `then`: any - the expression to evaluate if the condition is true\n\n" +
		"@param `else`: any - the expression to evaluate if the condition is false\n\n"

	SumTime string = codeBlockStart +
		"sumTime( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"       , expression: number\n" +
		"      [, where expression?: bool] )\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### SumTime\n" +
		"Calculates the sum of a numeric expression over a range of time records.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to sum\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `expression: number` - numeric expression - this is what will be summed for each qualifying slice\n\n" +
		"@param `condition?: boolean` - if used, only slices where this expression returns true will be evaluated\n\n"

	CountTime string = codeBlockStart +
		"countTime( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"        [, where expression?: bool] )\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### CountTime\n" +
		"Calculates the count of time records in a range.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `expression?: boolean` - if used, only slices where this expression returns true will be counted\n\n"

	FindFirstTime string = codeBlockStart +
		"findFirstTime( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"             , where condition: boolean\n" +
		"             , order by ordering: string|number|date|dateTime )\n" +
		codeBlockReturns + "timeRecord\n" +
		codeBlockEnd +
		"### FindFirstTime\n" +
		"Returns the first time record that meets `condition`, ordered by `ordering`.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n" +
		"@param `ordering: string|number|date|dateTime` - the value to order by\n\n"

	SumSchedule string = codeBlockStart +
		"sumSchedule( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"           , expression: number\n" +
		"          [, where expression?: bool] )\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### SumSchedule\n" +
		"Calculates the sum of a numeric expression over a range of schedule records.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to sum\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `expression: number` - numeric expression - this is what will be summed for each qualifying slice\n\n" +
		"@param `condition?: boolean` - if used, only slices where this expression returns true will be evaluated\n\n"

	CountSchedule string = codeBlockStart +
		"countSchedule( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"            [, where expression?: bool] )\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### CountSchedule\n" +
		"Calculates the count of schedule records in a range.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `expression?: boolean` - if used, only slices where this expression returns true will be counted\n\n"

	FindFirstSchedule string = codeBlockStart +
		"findFirstSchedule( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                 , where condition: boolean\n" +
		"                 , order by ordering: string|number|date|dateTime )\n" +
		codeBlockReturns + "scheduleRecord\n" +
		codeBlockEnd +
		"### FindFirstSchedule\n" +
		"Returns the first schedule record that meets `condition`, ordered by `ordering`.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n" +
		"@param `ordering: string|number|date|dateTime` - the value to order by\n\n"

	CountException string = codeBlockStart +
		"countException( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"             [, where condition?: bool] )\n" +
		codeBlockReturns + "exception\n" +
		codeBlockEnd +
		"### CountException\n" +
		"Calculates the count of exception records in a range.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the exception currently being evaluated\n\n" +
		"@param `condition?: boolean` - if used, only exceptions where this expression returns true will be counted\n\n"

	FindFirstTorDetail string = codeBlockStart +
		"findFirstTorDetail( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                  , where condition: boolean\n" +
		"                  , order by ordering: string|number|date|dateTime )\n" +
		codeBlockReturns + "torDetailRecord\n" +
		codeBlockEnd +
		"### FindFirstTorDetail\n" +
		"Returns the first TOR detail record that meets `condition`, ordered by `ordering`.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the record currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n" +
		"@param `ordering: string|number|date|dateTime` - the value to order by\n\n"

	FindFirstDayForward string = codeBlockStart +
		"findFirstDayForward( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                   , where condition: boolean )\n" +
		codeBlockReturns + "date\n" +
		codeBlockEnd +
		"### FindFirstDayForward\n" +
		"Returns the first date that meets `condition`, going from the start to the end of `range`, excluding the start date.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the day currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n"

	FindFirstDayBackward string = codeBlockStart +
		"findFirstBackward( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                 , where condition: boolean )\n" +
		codeBlockReturns + "date\n" +
		codeBlockEnd +
		"### FindFirstDayBackward\n" +
		"Returns the first date that meets `condition`, going from the end to the start of `range`, excluding the end date.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the day currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n"

	FindFirstDeletedTime string = codeBlockStart +
		"findFirstDeletedTime( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                    , where condition: boolean\n" +
		"                    , order by ordering: string|number|date|dateTime )\n" +
		codeBlockReturns + "timeRecord\n" +
		codeBlockEnd +
		"### FindFirstDeletedTime\n" +
		"Returns the first deleted time record that meets `condition`, ordered by `ordering`. Will only return time records" +
		" that have been deleted within the time entry window in the front-end (as opposed to within a script, for instance).\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n" +
		"@param `ordering: string|number|date|dateTime` - the value to order by\n\n"

	LongestConsecutiveRange string = codeBlockStart +
		"longestConsecutiveRange( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                       , where condition: bool )\n" +
		codeBlockReturns + "dateRange\n" +
		codeBlockEnd +
		"### LongestConsecutiveRange\n" +
		"evaluates a `condition` over a date range and returns the longest consecutive date range where `condition` is true.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the date currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate in finding the return range\n\n"

	FirstConsecutiveDay string = codeBlockStart +
		"firstConsecutiveDay( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                   , where condition: bool )\n" +
		codeBlockReturns + "date\n" +
		codeBlockEnd +
		"### FirstConsecutiveDay\n" +
		"evaluates a `condition` over a date range and returns the first date of the longest consecutive range where `condition` is true." +
		" This is effectively the same as `longestConsecutiveRange().start`\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the date currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate in finding the return date\n\n"

	LastConsecutiveDay string = codeBlockStart +
		"lastConsecutiveDay( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                   , where condition: bool )\n" +
		codeBlockReturns + "date\n" +
		codeBlockEnd +
		"### LastConsecutiveDay\n" +
		"evaluates a `condition` over a date range and returns the last date of the longest consecutive range where `condition` is true." +
		" This is effectively the same as `longestConsecutiveRange().end`\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the date currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate in finding the return date\n\n"
)
