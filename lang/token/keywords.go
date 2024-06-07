package token

func keywords() map[string]Type {
	// We define keywords here, including builtin functions.
	//
	// Global Variables are not included - those are handled by the parser along
	// with user-defined variable idents.
	//
	// Access methods are also delegated to the parser, as they are not keywords
	// and require context to be resolved.
	return map[string]Type{
		"var": T_VAR,

		// word literals
		"null":  T_NULL,
		"true":  T_TRUE,
		"false": T_FALSE,

		// logical
		// WFLang allows the actual words "and" and "or" to be used as logical operators.
		// TODO - check if case sensitive.
		// TODO - shall we just refuse to recognize this? This is an opinionated
		// tool, after all.
		"and": T_AND,
		"or":  T_OR,
		"not": T_BANG,

		// semantic keywords
		"alias": T_ALIAS,
		"over":  T_OVER,
		"where": T_WHERE,
		"order": T_ORDER,
		"by":    T_BY,
		"in":    T_IN,
		"set":   T_SET,

		// --- builtins...
		"if": T_IF, // if is actually implemented as a global function, not a keyword

		// ...math
		"min":        T_MIN,
		"max":        T_MAX,
		"round":      T_ROUND,
		"roundup":    T_ROUNDUP,
		"rounddown":  T_ROUNDDOWN,
		"roundtoint": T_ROUNDTOINT,

		// ...time summary
		"sumtime":   T_SUMTIME,
		"counttime": T_CNTTIME,
		"mintime":   T_MINTIME,
		"maxtime":   T_MAXTIME,
		"avgtime":   T_AVGTIME,

		// ...schedule summary
		"sumschedule":     T_SUMSCHED,
		"minschedule":     T_MINSCHED,
		"maxschedule":     T_MAXSCHED,
		"averageschedule": T_AVGSCHED,
		"countschedule":   T_CNTSCHED,

		// ...swipe location functions
		"swipe_in_latitude_in_range":   T_SWIPEINLATINRANGE,
		"swipe_in_longitude_in_range":  T_SWIPEINLONGINRANGE,
		"swipe_out_latitude_in_range":  T_SWIPEOUTLATINRANGE,
		"swipe_out_longitude_in_range": T_SWIPEOUTLONGINRANGE,

		// ...general summary
		"count": T_COUNT,
		"sum":   T_SUM,

		// ...exception summary
		"countexception":   T_CNTEXCEPT,
		"sumexception":     T_SUMEXCEPT,
		"minexception":     T_MINEXCEPT,
		"maxexception":     T_MAXEXCEPT,
		"averageexception": T_AVGEXCEPT,

		// ...findfirst
		"findfirsttime":     T_FINDTIME,
		"findfirstschedule": T_FINDSCHED,
		// TODO - check if findfirstexception is a thing??? Not in docs but that doesn't mean it ain't there.
		"findfirstdayforward":  T_FINDDAYFWD,
		"findfirstdaybackward": T_FINDDAYBWD,
		"findfirstdeletedtime": T_FINDDELETED,
		"findnthtime":          T_FINDNTH,

		// ...TOR
		"findfirsttordetail":         T_FINDTOR,
		"getbooleanfieldfromtor":     T_GETTORBOOL,
		"getdatefieldfromtor":        T_GETTORDATE,
		"getnumberfieldfromtor":      T_GETTORNUM,
		"getselectfieldvaluefromtor": T_GETTORSELECT,
		"getstringfieldfromtor":      T_GETTORSTR,

		// ...banks
		"accrued":       T_ACCRUED,
		"accruedbefore": T_ACCRUEDBEFORE,
		"balance":       T_BALANCE,
		"callsql":       T_CALLSQL,

		// ...consecutive ranges
		"longestconsecutiverange": T_LCONSECUTIVERANGE,
		"firstconsecutiveday":     T_FCONSECUTIVEDAY,
		"lastconsecutiveday":      T_LCONSECUTIVEDAY,

		// ...strings
		"contains":    T_CONTAINS,
		"indexof":     T_INDEXOF,
		"substr":      T_SUBSTR,
		"tolowercase": T_TOLOWER,
		"touppercase": T_TOUPPER,

		// ...LD
		"ldlookup":    T_LDLOOKUP,
		"ldvalidate":  T_LDVALIDATE,
		"rangelookup": T_RANGELOOKUP,

		// ...datetime
		"convertdttmbytimezone": T_CONVERTDTTMTZ,
		"lengthofservice":       T_LENGTHOFSERVICE,
		"makedate":              T_MAKEDATE,
		"makedatetime":          T_MAKEDTTM,
		"makedatetimerange":     T_MAKEDTTMRANGE,
		"semimonthlyperiod":     T_SEMIMONTHLYPERIOD,

		// ...holidays
		"countholidays": T_COUNTHOLIDAYS,
		"getholiday":    T_GETHOLIDAY,

		// ...employee attribute
		"employee_attribute_exists": T_EMPATTREXISTS,
		"employee_attribute":        T_EMPATTR,
		"getattributecalcdate":      T_GETATTRCALCDATE,

		// ...config
		"paycodeinschedulemap":  T_PAYCODEINSCHEDMAP,
		"paycodeintimesheetmap": T_PAYCODEINTSTMAP,

		// ...misc - the stuff I don't know if we'll ever need
		// we're skipping any crew-related functions, as we'll realistically never implement Crew Management.
		"getpaycurrencycode":   T_GETPAYCURRCODE,
		"getsysdatebytimezone": T_GETSYSDATEBYTZ,
		"countshiftchanges":    T_CNTSHIFTCHANGES,
	}
}
