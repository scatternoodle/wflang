package builtins

import "github.com/scatternoodle/wflang/lang/types"

type Builtin struct {
	Name string
	Type types.BaseType
}

func Builtins() map[string]Builtin {
	return map[string]Builtin{
		"if":                           {"if", types.BOOL},
		"min":                          {"min", types.NUMBER},
		"max":                          {"max", types.NUMBER},
		"contains":                     {"contains", types.BOOL},
		"sum":                          {"sum", types.NUMBER},
		"count":                        {"count", types.NUMBER},
		"sumtime":                      {"sumTime", types.NUMBER},
		"counttime":                    {"countTime", types.NUMBER},
		"findfirsttime":                {"findFirstTime", types.NUMBER},
		"sumschedule":                  {"sumSchedule", types.NUMBER},
		"countschedule":                {"countSchedule", types.NUMBER},
		"findfirstschedule":            {"findFirstSchedule", types.SCHEDREC},
		"countexception":               {"countException", types.NUMBER},
		"findfirsttordetail":           {"findFirstTorDetail", types.TORDTL},
		"findfirstdayforward":          {"findFirstDayForward", types.DATE},
		"findfirstdaybackward":         {"findFirstDayBackward", types.DATE},
		"findfirstdeletedtime":         {"findFirstDeletedTime", types.DATE},
		"longestconsecutiverange":      {"longestConsecutiveRange", types.DATERNG},
		"firstconsecutiveday":          {"firstConsecutiveDay", types.DATE},
		"lastconsecutiveday":           {"lastConsecutiveDay", types.DATE},
		"findnthtime":                  {"findNthTime", types.TIME},
		"accrued":                      {"accrued", types.NUMBER},
		"accruedbefore":                {"accruedBefore", types.NUMBER},
		"balance":                      {"balance", types.NUMBER},
		"callsql":                      {"callSql", types.OBJECT},
		"convertdttmbytimezone":        {"convertDttmByTimezone", types.DTTM},
		"countgroupcalc":               {"countGroupCalc", types.NUMBER},
		"countholidays":                {"countHolidays", types.NUMBER},
		"getholiday":                   {"getHoliday", types.STRING},
		"counthomecrewmembers":         {"countHomeCrewMembers", types.NUMBER},
		"countshiftchanges":            {"countShiftChanges", types.NUMBER},
		"employee_attribute_exists":    {"employee_attribute_exists", types.BOOL},
		"employee_attribute":           {"employee_attribute", types.EMPATTR},
		"getattributecalcdate":         {"getAttributeCalcDate", types.DATE},
		"getbooleanfieldfromtor":       {"getBooleanFieldFromTor", types.BOOL},
		"getdatefieldfromtor":          {"getDateFieldFromTor", types.DATE},
		"getnumberfieldfromtor":        {"getNumberFieldFromTor", types.NUMBER},
		"getpaycurrencycode":           {"getPayCurrencyCode", types.STRING},
		"getselectfieldvaluefromtor":   {"getSelectFieldValueFromTor", types.STRING},
		"getstringfieldfromtor":        {"getStringFieldFromTor", types.STRING},
		"getsysdatebytimezone":         {"getSysDateByTimezone", types.POLICY_ID},
		"ldlookup":                     {"ldLookup", types.LDREC},
		"ldvalidate":                   {"ldValidate", types.BOOL},
		"indexof":                      {"indexof", types.NUMBER},
		"lengthofservice":              {"lengthOfService", types.NUMBER},
		"makedate":                     {"makeDate", types.DATE},
		"makedatetime":                 {"makeDateTime", types.DTTM},
		"makedatetimerange":            {"makeDateTimeRange", types.DTTMRNG},
		"paycodeinschedulemap":         {"payCodeInScheduleMap", types.BOOL},
		"paycodeintimesheetmap":        {"payCodeInTimeSheetMap", types.BOOL},
		"rangelookup":                  {"rangeLookup", types.NUMBER},
		"round":                        {"round", types.NUMBER},
		"roundup":                      {"roundUp", types.NUMBER},
		"rounddown":                    {"roundDown", types.NUMBER},
		"roundtoint":                   {"roundToInt", types.NUMBER},
		"semimonthlyperiod":            {"semiMonthlyPeriod", types.PERIOD},
		"substr":                       {"substr", types.STRING},
		"tolowercase":                  {"tolowercase", types.STRING},
		"touppercase":                  {"touppercase", types.STRING},
		"minschedule":                  {"minSchedule", types.NUMBER},
		"maxschedule":                  {"maxSchedule", types.NUMBER},
		"avgschedule":                  {"avgSchedule", types.NUMBER},
		"mintime":                      {"minTime", types.NUMBER},
		"maxtime":                      {"maxTime", types.NUMBER},
		"avgtime":                      {"avgTime", types.NUMBER},
		"swipe_in_latitude_in_range":   {"Swipe_in_latitude_in_range", types.BOOL},
		"swipe_in_longitude_in_range":  {"Swipe_in_longitude_in_range", types.BOOL},
		"swipe_out_latitude_in_range":  {"Swipe_out_latitude_in_range", types.BOOL},
		"swipe_out_longitude_in_range": {"Swipe_out_longitude_in_range", types.BOOL},
		"sumexception":                 {"sumException", types.NUMBER},
		"minexception":                 {"minException", types.NUMBER},
		"maxexception":                 {"maxException", types.NUMBER},
		"averageexception":             {"averageException", types.NUMBER},
	}
}
