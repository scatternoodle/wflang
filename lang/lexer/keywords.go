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

		// --- builtins... shouldn't be here, they're moving on out to builtins package.
		// "if": token.T_IF, // if is actually implemented as a global function, not a keyword

		// ...math
		// "min":        token.T_MIN,
		// "max":        token.T_MAX,
		"round":      token.T_ROUND,
		"roundup":    token.T_ROUNDUP,
		"rounddown":  token.T_ROUNDDOWN,
		"roundtoint": token.T_ROUNDTOINT,

		// ...time summary
		// "sumtime":   token.T_SUMTIME,
		// "counttime": token.T_CNTTIME,
		"mintime": token.T_MINTIME,
		"maxtime": token.T_MAXTIME,
		"avgtime": token.T_AVGTIME,

		// ...schedule summary
		// "sumschedule":     token.T_SUMSCHED,
		"minschedule":     token.T_MINSCHED,
		"maxschedule":     token.T_MAXSCHED,
		"averageschedule": token.T_AVGSCHED,
		// "countschedule":   token.T_CNTSCHED,

		// ...swipe location functions
		"swipe_in_latitude_in_range":   token.T_SWIPEINLATINRANGE,
		"swipe_in_longitude_in_range":  token.T_SWIPEINLONGINRANGE,
		"swipe_out_latitude_in_range":  token.T_SWIPEOUTLATINRANGE,
		"swipe_out_longitude_in_range": token.T_SWIPEOUTLONGINRANGE,

		// ...general summary
		// "count": token.T_COUNT,
		// "sum":   token.T_SUM,

		// ...exception summary
		// "countexception":   token.T_CNTEXCEPT,
		"sumexception":     token.T_SUMEXCEPT,
		"minexception":     token.T_MINEXCEPT,
		"maxexception":     token.T_MAXEXCEPT,
		"averageexception": token.T_AVGEXCEPT,

		// ...findfirst
		"findfirsttime":     token.T_FINDTIME,
		"findfirstschedule": token.T_FINDSCHED,
		// TODO - check if findfirstexception is a thing??? Not in docs but that doesn't mean it ain't there.
		"findfirstdayforward":  token.T_FINDDAYFWD,
		"findfirstdaybackward": token.T_FINDDAYBWD,
		"findfirstdeletedtime": token.T_FINDDELETED,
		"findnthtime":          token.T_FINDNTH,

		// ...TOR
		"findfirsttordetail":         token.T_FINDTOR,
		"getbooleanfieldfromtor":     token.T_GETTORBOOL,
		"getdatefieldfromtor":        token.T_GETTORDATE,
		"getnumberfieldfromtor":      token.T_GETTORNUM,
		"getselectfieldvaluefromtor": token.T_GETTORSELECT,
		"getstringfieldfromtor":      token.T_GETTORSTR,

		// ...banks
		"accrued":       token.T_ACCRUED,
		"accruedbefore": token.T_ACCRUEDBEFORE,
		"balance":       token.T_BALANCE,

		"callsql": token.T_CALLSQL,

		// ...consecutive ranges
		"longestconsecutiverange": token.T_LCONSECUTIVERANGE,
		"firstconsecutiveday":     token.T_FCONSECUTIVEDAY,
		"lastconsecutiveday":      token.T_LCONSECUTIVEDAY,

		// ...strings
		"contains":    token.T_CONTAINS,
		"indexof":     token.T_INDEXOF,
		"substr":      token.T_SUBSTR,
		"tolowercase": token.T_TOLOWER,
		"touppercase": token.T_TOUPPER,

		// ...LD
		"ldlookup":    token.T_LDLOOKUP,
		"ldvalidate":  token.T_LDVALIDATE,
		"rangelookup": token.T_RANGELOOKUP,

		// ...datetime
		"convertdttmbytimezone": token.T_CONVERTDTTMTZ,
		"lengthofservice":       token.T_LENGTHOFSERVICE,
		"makedate":              token.T_MAKEDATE,
		"makedatetime":          token.T_MAKEDTTM,
		"makedatetimerange":     token.T_MAKEDTTMRANGE,
		"semimonthlyperiod":     token.T_SEMIMONTHLYPERIOD,

		// ...holidays
		"countholidays": token.T_COUNTHOLIDAYS,
		"getholiday":    token.T_GETHOLIDAY,

		// ...employee attribute
		"employee_attribute_exists": token.T_EMPATTREXISTS,
		"employee_attribute":        token.T_EMPATTR,
		"getattributecalcdate":      token.T_GETATTRCALCDATE,

		// ...config
		"paycodeinschedulemap":  token.T_PAYCODEINSCHEDMAP,
		"paycodeintimesheetmap": token.T_PAYCODEINTSTMAP,

		// ...misc - the stuff I don't know if we'll ever need
		// we're skipping any crew-related functions, as we'll realistically never implement Crew Management.
		"getpaycurrencycode":   token.T_GETPAYCURRCODE,
		"getsysdatebytimezone": token.T_GETSYSDATEBYTZ,
		"countshiftchanges":    token.T_CNTSHIFTCHANGES,
	}
}
