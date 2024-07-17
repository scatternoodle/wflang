package hovertext

const (
	SumTime string = "```wflang\n" +
		"sumTime( over range: timePeriod [alias aliasName?: string]\n" +
		"       , sumExpression: number\n" +
		"      [, where whereExpression: bool] )\n" +
		"```\n" +
		"### SumTime\n" +
		"SumTime calculates the sum of a numeric expression over a range of time records.\n\n" +
		"@param `range: timePeriod` - the time period over which to sum\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `sumExpression: number` - numeric expression - this is what will be summed for each qualifying slice\n\n" +
		"@param `whereExpression?: boolean` - if used, only slices where this expression returns true will be evaluated\n\n"
)
