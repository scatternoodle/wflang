package hovertext

const (
	codeBlockStart   string = "```wflang\nSYNTAX:\n"
	codeBlockReturns string = "\nRETURNS: "
	codeBlockEnd     string = "```\n\n---\n\n"
)

const (
	Min string = codeBlockStart +
		"min(args: number...)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### Min\n\n" +
		"Returns the smallest of its arguments. It takes a list of arguments as long as you like, which can be any" +
		" expression that evaluates to a number.\n\n" +
		"@param `args: number...` - list of numbers to compare\n\n"

	Max string = codeBlockStart +
		"max(args: number...)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### Max\n\n" +
		"Returns the largest of its arguments. It takes a list of arguments as long as you like, which can be any" +
		" expression that evaluates to a number.\n\n" +
		"@param `args: number...` - list of numbers to compare\n\n"

	Contains string = codeBlockStart +
		"contains(string1: string, string2: string)\n" +
		codeBlockReturns + "boolean\n" +
		codeBlockEnd +
		"### Contains\n\n" +
		"Returns true if `string2` is a substring of `string1`.\n\n" +
		"@param `string1: string` - the string to search in\n\n" +
		"@param `string2: string` - the string to search for\n\n"

	Sum string = codeBlockStart +
		"sum( by interval: day|week|period| over range: day|week|period|dateRange alias aliasName: string\n" +
		"   , expression: number \n" +
		"  [, where condition: boolean] )\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### Sum\n\n" +
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
		"### Count\n\n" +
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
		"### If\n\n" +
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
		"### SumTime\n\n" +
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
		"### CountTime\n\n" +
		"Calculates the count of time records in a range.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `expression?: boolean` - if used, only slices where this expression returns true will be counted\n\n"

	FindFirstTime string = codeBlockStart +
		"findFirstTime( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"             , where condition: boolean\n" +
		"             , order by ordering: string|number|date|dateTime )\n" +
		codeBlockReturns + "timeRecord|null\n" +
		codeBlockEnd +
		"### FindFirstTime\n\n" +
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
		"### SumSchedule\n\n" +
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
		"### CountSchedule\n\n" +
		"Calculates the count of schedule records in a range.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `expression?: boolean` - if used, only slices where this expression returns true will be counted\n\n"

	FindFirstSchedule string = codeBlockStart +
		"findFirstSchedule( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                 , where condition: boolean\n" +
		"                 , order by ordering: string|number|date|dateTime )\n" +
		codeBlockReturns + "scheduleRecord|null\n" +
		codeBlockEnd +
		"### FindFirstSchedule\n\n" +
		"Returns the first schedule record that meets `condition`, ordered by `ordering`.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n" +
		"@param `ordering: string|number|date|dateTime` - the value to order by\n\n"

	CountException string = codeBlockStart +
		"countException( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"             [, where condition?: bool] )\n" +
		codeBlockReturns + "exception|null\n" +
		codeBlockEnd +
		"### CountException\n\n" +
		"Calculates the count of exception records in a range.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the exception currently being evaluated\n\n" +
		"@param `condition?: boolean` - if used, only exceptions where this expression returns true will be counted\n\n"

	FindFirstTorDetail string = codeBlockStart +
		"findFirstTorDetail( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                  , where condition: boolean\n" +
		"                  , order by ordering: string|number|date|dateTime )\n" +
		codeBlockReturns + "torDetailRecord|null\n" +
		codeBlockEnd +
		"### FindFirstTorDetail\n\n" +
		"Returns the first TOR detail record that meets `condition`, ordered by `ordering`.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the record currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n" +
		"@param `ordering: string|number|date|dateTime` - the value to order by\n\n"

	FindFirstDayForward string = codeBlockStart +
		"findFirstDayForward( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                   , where condition: boolean )\n" +
		codeBlockReturns + "date|null\n" +
		codeBlockEnd +
		"### FindFirstDayForward\n\n" +
		"Returns the first date that meets `condition`, going from the start to the end of `range`, excluding the start date.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the day currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n"

	FindFirstDayBackward string = codeBlockStart +
		"findFirstBackward( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                 , where condition: boolean )\n" +
		codeBlockReturns + "date|null\n" +
		codeBlockEnd +
		"### FindFirstDayBackward\n\n" +
		"Returns the first date that meets `condition`, going from the end to the start of `range`, excluding the end date.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the day currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n"

	FindFirstDeletedTime string = codeBlockStart +
		"findFirstDeletedTime( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                    , where condition: boolean\n" +
		"                    , order by ordering: string|number|date|dateTime )\n" +
		codeBlockReturns + "timeRecord|null\n" +
		codeBlockEnd +
		"### FindFirstDeletedTime\n\n" +
		"Returns the first deleted time record that meets `condition`, ordered by `ordering`. Will only return time records" +
		" that have been deleted within the time entry window in the front-end (as opposed to within a script, for instance).\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n" +
		"@param `ordering: string|number|date|dateTime` - the value to order by\n\n"

	LongestConsecutiveRange string = codeBlockStart +
		"longestConsecutiveRange( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                       , where condition: bool )\n" +
		codeBlockReturns + "dateRange|null\n" +
		codeBlockEnd +
		"### LongestConsecutiveRange\n\n" +
		"evaluates a `condition` over a date range and returns the longest consecutive date range where `condition` is true.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the date currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate in finding the return range\n\n"

	FirstConsecutiveDay string = codeBlockStart +
		"firstConsecutiveDay( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                   , where condition: bool )\n" +
		codeBlockReturns + "date|null\n" +
		codeBlockEnd +
		"### FirstConsecutiveDay\n\n" +
		"evaluates a `condition` over a date range and returns the first date of the longest consecutive range where `condition` is true." +
		" This is effectively the same as `longestConsecutiveRange().start`\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the date currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate in finding the return date\n\n"

	LastConsecutiveDay string = codeBlockStart +
		"lastConsecutiveDay( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"                   , where condition: bool )\n" +
		codeBlockReturns + "date|null\n" +
		codeBlockEnd +
		"### LastConsecutiveDay\n\n" +
		"evaluates a `condition` over a date range and returns the last date of the longest consecutive range where `condition` is true." +
		" This is effectively the same as `longestConsecutiveRange().end`\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to count\n\n" +
		"@param `aliasName?: string` - alias for the date currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate in finding the return date\n\n"

	FindNthTime string = codeBlockStart +
		"findNthTime( over range: day|week|period|dateRange [alias aliasName?: string]\n" +
		"             , where condition: boolean\n" +
		"             , order by ordering: string|number|date|dateTime\n" +
		"             , n: number )\n" +
		codeBlockReturns + "timeRecord|null\n" +
		codeBlockEnd +
		"### FindNthTime\n\n" +
		"Returns the Nth time record that meets `condition`, ordered by `ordering`.\n\n" +
		"@param `range: day|week|period|dateRange` - the time period over which to search\n\n" +
		"@param `aliasName?: string` - alias for the slice currently being evaluated\n\n" +
		"@param `condition: boolean` - the condition to evaluate\n\n" +
		"@param `ordering: string|number|date|dateTime` - the value to order by\n\n" +
		"@param `n`: number - the 'N' in 'Nth', e.g. 2 gets you the 2nd timeRecord meeting `condition`\n\n"

	Accrued string = codeBlockStart +
		"accrued(bank: ident, start: day|date to end: day|date)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### Accrued\n\n" +
		"Returns the balance accrued to `bank` over date range of `start` to `end`. Only returns the product of **positive** transactions.\n\n" +
		"@param `bank: ident` - the identifier of the bank\n\n" +
		"@param `start: day|date` - start of the range\n\n" +
		"@param `end: day|date` - end of the range\n\n"

	BalanceAccruedBefore string = codeBlockStart +
		"balanceAccruedBefore(bank: string, asOfDate: day|date)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### BalanceAccruedBefore\n\n" +
		"Returns all balance accrued to `bank` prior to `asOfDate`. Only returns the product of **positive** transactions." +
		" Note that syntax is inconsistent to the `accrued` function which takes an `ident` for its `bank` parameter, whereas" +
		" for `balanceAccruedBefore`, `bank` must be a string literal.\n\n" +
		"@param `bank: string` - the policy ID of the bank\n\n" +
		"@param `asOfDate: day|date` - accruals are summed prior to this date\n\n"

	ConvertDttmByTimezone string = codeBlockStart +
		"convertDttmByTimezone( timezone: string, dttm: dateTime )\n" +
		codeBlockReturns + "dateTime|null\n" +
		codeBlockEnd +
		"### ConvertDttmByTimezone\n\n" +
		"Converts the given dateTime to the specified timezone.\n\n" +
		"@param `timezone: string` - the timezone to convert to, which must be the policy ID of a valid Timezones Policy\n\n" +
		"@param `dttm: dateTime` - the dateTime to be converted\n\n"

	CountHolidays string = codeBlockStart +
		"countHolidays(timePeriod: day|date|dateRange|week|period, holidaySet: ident)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### CountHolidays\n\n" +
		"Returns a count of all holidays in `holidaySet` over `timePeriod`. Note that `holidaySet` must be an explicit ident" +
		" for a valid Holiday Set; you cannot use a string literal or assignment.HOLIDAYS, for instance. This means that," +
		" unlike with `getHoliday()`, the Policy ID of the Holiday Set in question must be known at compile time.\n\n" +
		"@param `timePeriod: day|date|dateRange|week|period` - date range to count over\n\n" +
		"@param `holidaySet: ident` - the policy ID of the holiday set to check\n\n"

	GetHoliday string = codeBlockStart +
		"getHoliday(holidaySet: string, date: day|date)\n" +
		codeBlockReturns + "string\n" +
		codeBlockEnd +
		"### GetHoliday\n\n" +
		"Checks the given holiday set on `date`. If no holidays found or if `holidaySet` is invalid, returns an empty string.\n\n" +
		"@param `holidaySet: string` - the Policy ID of a Holiday Set\n\n" +
		"@param `date: day|date` - the date to check in the set\n\n"

	Employee_attribute_exists string = codeBlockStart +
		"employee_attribute_exists(id: ident, asOf?: day|date)\n" +
		codeBlockReturns + "boolean\n" +
		codeBlockEnd +
		"### employee_attribute_exists\n\n" +
		"Returns true if the employee attribute with `id` exists on the timesheet on `asOf`. Use to nullcheck Employee Attributes" +
		" before calling `employee_attribute()`, usually nested within an if statement like so:\n\n" +
		"```wflang\n" +
		"if( employee_attribute_exists(MY_NUMBER_ATTRIBUTE, day)\n" +
		"  , employee_attribute(MY_NUMBER_ATTRIBUTE, day)\n" +
		"  , 0 )\n" +
		"```\n\n" +
		"@param `id: ident` - the Policy ID of the employee attribute\n\n" +
		"@param `asOf?: day|date` - date on which to check. If omitted, {TBC} is used \n\n" // TODO - confirm behaviour when asOf is omitted.

	Employee_attribute string = codeBlockStart +
		"employee_attribute(id: ident, asOf?: day|date)\n" +
		codeBlockReturns + "date|dateTime|number|string|boolean|null\n" +
		codeBlockEnd +
		"### employee_attribute\n\n" +
		"Returns the value of the Employee Attribute with `id` on `asOf`.\n\nThe type of the return value depends on the type" +
		" of the Employee Attribute (parent node >> Main >> Attribute Type). Interestingly, the Attribute Type dropdown lists" +
		" quite a lot of options (e.g. exception severity, color), but only permits saving the policy with one of [string, date," +
		" dateTime, number, boolean].\n\n" +
		"@param `id: ident` - the Policy ID of the employee attribute\n\n" +
		"@param `asOf?: day|date` - date on which to check. If omitted, {TBC} is used \n\n" // TODO - confirm behaviour when asOf is omitted.

	GetAttributeCalculationDate string = codeBlockStart +
		"getAttributeCalculationDate(id: string, asOf: day|date)\n" +
		codeBlockReturns + "date|null\n" +
		codeBlockEnd +
		"### GetAttributeCalculationDate\n\n" +
		"Returns the date on which a given employee attribute value was calculated. Useful if you need to go and pull other data" +
		" from the timesheet or employee record on that date.\n\nNote that despite having the similar syntax to `employee_attribute()` and" +
		" `employee_attribute_exists()`, `getAttributeCalculationDate()` differs in that `id` is a string not an ident, and" +
		" `asOf` is *not* optional. Easy to trip up on this.\n\n" +
		"@param `id: string` - the Policy ID of the Employee Attribute\n\n" +
		"@param `asOf: day|date` - the date on which to check the value of the Employee Attribute\n\n"

	GetBooleanFieldFromTor string = codeBlockStart +
		"getBooleanFieldFromTor(torId: string, fieldId: string)\n" +
		codeBlockReturns + "boolean\n" +
		codeBlockEnd +
		"### GetBooleanFieldFromTor\n\n" +
		"Takes a Time Off Request ID, a policy ID for a boolean field, and returns the value of that boolean. It has some nuances" +
		" to be aware of:\n\n" +
		"- `torId` param is string type but **must** resolve to an integer otherwise it'll compile then error at runtime.\n\n" +
		"- If `fieldId` is not a valid Time Off Request Field Policy ID, formula compiles then errors at runtime.\n\n" +
		"- If `fieldId` exists but is of a different type, you guessed it... compiles but errors at runtime.\n\n" +
		"- If no TOR with `torId` is found at runtime, returns false.\n\n" +
		"@param `torId: string` - TOR ID - must resolve to an integer\n\n" +
		"@param `fieldId: string` - the Policy ID of the TOR field\n\n"

	GetDateFieldFromTor string = codeBlockStart +
		"getDateFieldFromTor(torId: string, fieldId: string)\n" +
		codeBlockReturns + "date|null\n" +
		codeBlockEnd +
		"### GetDateFieldFromTor\n\n" +
		"Takes a Time Off Request ID, a policy ID for a boolean field, and returns the value of that boolean. It has some nuances" +
		" to be aware of:\n\n" +
		"- `torId` param is string type but **must** resolve to an integer otherwise it'll compile then error at runtime.\n\n" +
		"- If `fieldId` is not a valid Time Off Request Field Policy ID, formula compiles then errors at runtime.\n\n" +
		"- If `fieldId` exists but is of a different type, you guessed it... compiles but errors at runtime.\n\n" +
		"- If no TOR with `torId` is found at runtime, returns null.\n\n" +
		"@param `torId: string` - TOR ID - must resolve to an integer\n\n" +
		"@param `fieldId: string` - the Policy ID of the TOR field\n\n"

	GetNumberFieldFromTor string = codeBlockStart +
		"getNumberFieldFromTor(torId: string, fieldId: string)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### GetNumberFieldFromTor\n\n" +
		"Takes a Time Off Request ID, a policy ID for a number field, and returns the value of that number. It has some nuances" +
		" to be aware of:\n\n" +
		"- `torId` param is string type but **must** resolve to an integer otherwise it'll compile then error at runtime.\n\n" +
		"- If `fieldId` is not a valid Time Off Request Field Policy ID, formula compiles then errors at runtime.\n\n" +
		"- If `fieldId` exists but is of a different type, you guessed it... compiles but errors at runtime.\n\n" +
		"- If no TOR with `torId` is found at runtime, returns zero.\n\n" +
		"@param `torId: string` - TOR ID - must resolve to an integer\n\n" +
		"@param `fieldId: string` - the Policy ID of the TOR field\n\n"

	GetSelectFieldValueFromTor string = codeBlockStart +
		"getSelectFieldValueFromTor(torId: string, fieldId: string)\n" +
		codeBlockReturns + "string|null\n" +
		codeBlockEnd +
		"### GetSelectFieldValueFromTor\n\n" +
		"Takes a Time Off Request ID, a policy ID for a select field, and returns the value of that field (selects are always strings). It has some nuances" +
		" to be aware of:\n\n" +
		"- `torId` param is string type but **must** resolve to an integer otherwise it'll compile then error at runtime.\n\n" +
		"- If `fieldId` is not a valid Time Off Request Field Policy ID, formula compiles then errors at runtime.\n\n" +
		"- If `fieldId` exists but is of a different type, you guessed it... compiles but errors at runtime.\n\n" +
		"- If no TOR with `torId` is found at runtime, returns null.\n\n" +
		"@param `torId: string` - TOR ID - must resolve to an integer\n\n" +
		"@param `fieldId: string` - the Policy ID of the TOR field\n\n"

	GetStringValueFromTor string = codeBlockStart +
		"getStringFieldFromTor(torId: string, fieldId: string)\n" +
		codeBlockReturns + "string|null\n" +
		codeBlockEnd +
		"### GetStringFieldFromTor\n\n" +
		"Takes a Time Off Request ID, a policy ID for a string field, and returns the value of that string. It has some nuances" +
		" to be aware of:\n\n" +
		"- `torId` param is string type but **must** resolve to an integer otherwise it'll compile then error at runtime.\n\n" +
		"- If `fieldId` is not a valid Time Off Request Field Policy ID, formula compiles then errors at runtime.\n\n" +
		"- If `fieldId` exists but is of a different type, you guessed it... compiles but errors at runtime.\n\n" +
		"- If no TOR with `torId` is found at runtime, returns null.\n\n" +
		"@param `torId: string` - TOR ID - must resolve to an integer\n\n" +
		"@param `fieldId: string` - the Policy ID of the TOR field\n\n"

	GetSysDateByTimezone string = codeBlockStart +
		"getSysDateByTimezone(timezone: IDENT)\n" +
		codeBlockReturns + "date|null\n" +
		codeBlockEnd +
		"### GetSysDateByTimezone\n\n" +
		"Returns the current sysdate based on the TIME_ZONE value of the assignment. Interestingly, this only compiles when" +
		" we pass `timezone` as `assignment(context).TIME_ZONE` and will not allow us to explicitly specify a Time Zones" +
		" Policy ident. If assignment(context).TIME_ZONE is blank, returns the current sysdate according to the default" +
		" server timezone.\n\n" +
		"@param `timezone: ident` - Time Zones Policy Ident - can only take `assignment(context).TIME_ZONE`!\n\n"

	LDLookup string = codeBlockStart +
		"ldLookup( policyName: ident\n" +
		"        , field1: ident => value1: string\n" +
		"       [, field2?: ident, value2?: string\n" +
		"        , ...] )" +
		codeBlockReturns + "ldRecord|null\n" +
		codeBlockEnd +
		"### LDLookup\n\n" +
		"Looks up and returns an ldRecord object from the given LD Table and field / value pairs.\n\n" +
		"@param `policyName: ident` - the Policy ID of the LD Field that points to the table's primary key.\n\n" +
		"@param `fieldX: ident` - the Policy ID of the LD field to look up - normally just the same LD Field policy as `policyName`" +
		" but can be different for more complex multifield lookups where an LD Table supplies multiple LD Field Policies.\n\n" +
		"@param `valueX: string` - the value to match on. Always a string value.\n\n"

	LDValidate string = codeBlockStart +
		"ldValidate(field: ident, slice: timeRecord, asOf?: date)\n" +
		codeBlockReturns + "boolean\n" +
		codeBlockEnd +
		"### LDValidate\n\n" +
		"Returns true if the value of `field` in time record `slice` is a valid value in the LD table as of `asOf`.\n\n" +
		"@param `field: ident` - LD Field Policy ID\n\n" +
		"@param `slice: timeRecord` - the time record on which to check the LD field value\n\n" +
		"@param `asOf?: date` - function checks only for records in the LD table that are active on the `asOf` date\n\n" // TODO confirm default behaviour if omitted

	IndexOf string = codeBlockStart +
		"indexOf(x: string, y: string)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### IndexOf\n\n" +
		"Returns the index of string `y` in string `x`, or -1 if `y` is not in `x`.\n\n" +
		"@param `x: string`\n\n" +
		"@param `y: string`\n\n"

	GetPayCurrencyCode string = codeBlockStart +
		"getPayCurrencyCode(over timePeriod: day|date|dateRange|period)\n" +
		codeBlockReturns + "string\n" +
		codeBlockEnd +
		"### GetPayCurrencyCode\n\n" +
		"Returns the ISO Currency Code for `timePeriod`, or the string literal `\"MULTI\"` if there are multiple over the range." +
		" Determined by `assignment.PAY_CURRENCY_CODE`.\n\n" +
		"@param `timePeriod: day|date|dateRange|period`\n\n"

	CallSQL string = codeBlockStart +
		"callSQL(policyId: ident, [param1?: ident, value1?: any, ..., paramN?: ident, valueN?: any])\n" +
		codeBlockReturns + "any|null\n" +
		codeBlockEnd +
		"### CallSQL\n\n" +
		"Returns the results of a SQL Invocation. The results can be a `resultSet` object containing named columns, or" +
		" one of many other types.\n\n" +
		"@param `policyId: ident` - the SQL Invocation Policy ID\n\n" +
		"@param `paramN?: ident` - must correspond with a parameter on the Parameters tab of the SQL Invocation\n\n" +
		"@param `valueN?: any`\n\n"

	LengthOfService string = codeBlockStart +
		"lengthOfService(startDate: date, endDate: date, units: ident{days|months|years}, adjustTo?: TBC)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### LengthOfService\n\n" +
		"Returns the length of time in `units` between `startDate` `endDate`. It is acceptable to use a `startDate` that is" +
		" after `endDate`, which will return a negative number.\n\n" +
		"@param `startDate: date`\n\n" +
		"@param `endDate: date`\n\n" +
		"@param `units: ident` - accepted values are days, months, or years\n\n" +
		"@param `adjustTo?: ???` - TBC" // TODO confirm behaviour of adjustTo param

	MakeDate string = codeBlockStart +
		"makeDate(year: int, month: int, day: int)\n" +
		codeBlockReturns + "date\n" +
		codeBlockEnd +
		"### MakeDate\n\n" +
		"Returns a date object with the given `year`, `month`, `day`. Will compile but error at runtime if `month` > 12" +
		" or `day` > 31. Caution: it will let you save, and there **will be no runtime error** if `day` is 31 for a month" +
		" that has fewer than 31 days, which could cause unexpected behavior / errors in formulas that consume the date.\n\n" +
		"There appear to be no restrictions on `year`, and using a negative integer interestingly results in the product" +
		" of `2000-year`. e.g. `year = -42` results in `1958` being used as the year for the returned date. Using a negative" +
		" number below 2000 is also fine! e.g. `-2001` results in `0002`. Needless to say, I recommend just using positive" +
		" integers for the `year` value.\n\n" +
		"@param `year: int`\n\n" +
		"@param `month: int` - must be > 0 \n\n" +
		"@param `day: int`\n\n"

	MakeDateTime string = codeBlockStart +
		"makeDateTime(date: date, time: time, useDSTFallback?: boolean)\n" +
		codeBlockReturns + "dateTime\n" +
		codeBlockEnd +
		"### MakeDateTime\n\n" +
		"Returns a dateTime object with the given `date` and `time`. If `useDSTFallback` is true, the DST Fallback Hour will" +
		" be used if applicable. The Fallback is the extra hour from the DST Autumn change.\n\n" +
		"For more information on DST, see https://en.wikipedia.org/wiki/Daylight_saving_time\n\n" +
		"@param `date: date`\n\n" +
		"@param `time: time`\n\n" +
		"@param `useDSTFallback?: boolean` - if true, the DST Fallback Hour will be used if applicable\n\n"

	MakeDateTimeRange string = codeBlockStart +
		"makeDateTimeRange(start: dateTime, end: dateTime)\n" +
		codeBlockReturns + "dateTimeRange\n" +
		codeBlockEnd +
		"### MakeDateTimeRange\n\n" +
		"Returns a dateTimeRange object for the range of `start` to `end`. `start` must be <= `end`, otherwise formula will" +
		" compile but will error at runtime.\n\n" +
		"@param `start: dateTime`\n\n" +
		"@param `end: dateTime`\n\n"

	PayCodeInScheduleMap string = codeBlockStart +
		"payCodeInScheduleMap(payCode: {timeRecord|scheduleRecord}.PAY_CODE, asOf?: date)\n" +
		codeBlockReturns + "boolean\n" +
		codeBlockEnd +
		"### PayCodeInScheduleMap\n\n" +
		"Returns true if the given `timeRecord.PAY_CODE` is in the schedule map on `asOf` date.\n\n" +
		"This function must be used either in a slice context formula, or within a slice-level scope, e.g.\n\n" +
		"```wflang\n" +
		"countSchedule( over day alias x\n" +
		"             , where payCodeInScheduleMap(x.PAY_CODE) )\n" +
		"```\n\n" +
		"You cannot directly use a pay code IDENT or string literal for `payCode` - it can only be the PAY_CODE field on" +
		" a timeRecord or scheduleRecord object.\n\n" +
		"Arguably, this function has limited value when we consider that all codes should be in all maps as a best practice" +
		" to avoid application errors when employees switch Policy Profile mid-period.\n\n" +
		"@param `payCode: {timeRecord|scheduleRecord}.PAY_CODE` - must be the actual timeRecord.PAY_CODE field access," +
		" either in a slice-level formula or within a slice-level scope such as in a summary function clause\n\n" +
		"@param `asOf?: date` - if omitted, period.end is used\n\n"

	PayCodeInTimeSheetMap string = codeBlockStart +
		"payCodeInTimeSheetMap(payCode: {timeRecord|scheduleRecord}.PAY_CODE, asOf?: date)\n" +
		codeBlockReturns + "boolean\n" +
		codeBlockEnd +
		"### PayCodeInScheduleMap\n\n" +
		"Returns true if the given `timeRecord.PAY_CODE` is in the timesheet paycodes map on `asOf` date.\n\n" +
		"This function must be used either in a slice context formula, or within a slice-level scope, e.g.\n\n" +
		"```wflang\n" +
		"countSchedule( over day alias x\n" +
		"             , where payCodeInTimeSheetMap(x.PAY_CODE) )\n" +
		"```\n\n" +
		"You cannot directly use a pay code IDENT or string literal for `payCode` - it can only be the PAY_CODE field on" +
		" a timeRecord or scheduleRecord object.\n\n" +
		"Arguably, this function has limited value when we consider that all codes should be in all maps as a best practice" +
		" to avoid application errors when employees switch Policy Profile mid-period.\n\n" +
		"@param `payCode: {timeRecord|scheduleRecord}.PAY_CODE` - must be the actual timeRecord.PAY_CODE field access," +
		" either in a slice-level formula or within a slice-level scope such as in a summary function clause\n\n" +
		"@param `asOf?: date` - if omitted, period.end is used\n\n"

	Round string = codeBlockStart +
		"round(x: number, precision?: number=0)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### Round\n\n" +
		"Returns the number `x` rounded to `precision` decimal places. If `precision` is omitted, the number is rounded" +
		" to the nearest integer.\n\n" +
		"@param `x: number`\n\n" +
		"@param `precision?: number=0` - the number of decimal places to round to\n\n"

	RoundUp string = codeBlockStart +
		"roundUp(x: number, precision?: number=0)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### RoundUp\n\n" +
		"Returns the number `x` rounded up to `precision` decimal places. If `precision` is omitted, the number is rounded" +
		" to the nearest integer.\n\n" +
		"@param `x: number`\n\n" +
		"@param `precision?: number=0` - the number of decimal places to round to\n\n"

	RoundDown string = codeBlockStart +
		"roundDown(x: number, precision?: number=0)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### RoundDown\n\n" +
		"Returns the number `x` rounded down to `precision` decimal places. If `precision` is omitted, the number is rounded" +
		" to the nearest integer.\n\n" +
		"@param `x: number`\n\n" +
		"@param `precision?: number=0` - the number of decimal places to round to\n\n"

	RoundToInt string = codeBlockStart +
		"roundToInt(x: number)\n" +
		codeBlockReturns + "number\n" +
		codeBlockEnd +
		"### RoundToInt\n\n" +
		"Shorthand for `round(x, 0)` - returns the number `x` rounded to the nearest integer. If you need to force direction," +
		" use `roundUp()` or `roundDown()` instead.\n\n" +
		"@param `x: number`\n\n"

	SemiMonthlyPeriod string = codeBlockStart +
		"semiMonthlyPeriod(date: date)\n" +
		codeBlockReturns + "period\n" +
		codeBlockEnd +
		"### SemiMonthlyPeriod\n\n" +
		"Returns the period in which the given `date` falls, based on the semi-monthly pay period. The returned period is" +
		" based on calendar months, and is not connected with actual Policy Profile config in any way; effectively this is" +
		" just simple date-maths.\n\n" +
		"@param `date: date`\n\n"

	Substr string = codeBlockStart +
		"substr(x: string, start: number, length: number)\n" +
		codeBlockReturns + "string\n" +
		codeBlockEnd +
		"### Substr\n\n" +
		"Returns a substring of `x` starting at `start` with length `length`. Some usage considerations:\n\n" +
		"- `start` is 1-indexed (i.e. the first character is at index 1, not 0)\n\n" +
		"- If `start` is < 1, it is treated as 1\n\n" +
		"- If `length` is greater than the remaining length of the string, the substring will be truncated to the end of the string\n\n" +
		"- ***If `length` is negative, formula compiles but errors at runtime***\n\n" +
		"- If `length` is 0, an empty string is returned\n\n" +
		"@param `x: string`\n\n" +
		"@param `start: number`\n\n" +
		"@param `length: number`\n\n"
)
