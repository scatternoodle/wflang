package token

const (
	// misc

	T_ILLEGAL Type = "ILLEGAL"
	T_EOF     Type = "EOF"

	// literals

	T_IDENT  Type = "IDENT"
	T_NUM    Type = "NUMBER" // A la Javascript, a number is atokn = newToken(l, token.T_SLASH, '/') number is a runtime error.
	T_STRING Type = "STRING"

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
	T_LBRACE      Type = "{" // TODO - check - rare - only use case I'm aware of is to express times.
	T_RBRACE      Type = "}"
	T_LBRACKET    Type = "["
	T_RBRACKET    Type = "]"  // For specific array-like use cases such as "in" expressions.
	T_PERIOD      Type = "."  // Period can denote a decimal point or member access.
	T_DOLLAR      Type = "$"  // Dollas signs wrap Macros in WFLang.
	T_DOUBLEQUOTE Type = "\"" // For string literals. Single quotes are not used.

	// keywords / builtins

	T_VAR                 Type = "var"
	T_IF                  Type = "if"
	T_OVER                Type = "over"
	T_WHERE               Type = "where"
	T_ORDER               Type = "order"
	T_BY                  Type = "by"
	T_ALIAS               Type = "alias"
	T_IN                  Type = "in"
	T_SET                 Type = "set"
	T_NULL                Type = "null"
	T_TRUE                Type = "true"
	T_FALSE               Type = "false"
	T_MIN                 Type = "min"
	T_MAX                 Type = "max"
	T_SUM                 Type = "sum"
	T_SUMTIME             Type = "sumTime"
	T_SUMSCHED            Type = "sumSchedule"
	T_COUNT               Type = "count"
	T_CNTTIME             Type = "countTime"
	T_CNTSCHED            Type = "countSchedule"
	T_CNTEXCEPT           Type = "countException"
	T_FINDTIME            Type = "findFirstTime"
	T_FINDSCHED           Type = "findFirstSched"
	T_FINDTOR             Type = "findFirstTorDetail"
	T_FINDDAYFWD          Type = "findFirstDayForward"
	T_FINDDAYBWD          Type = "findFirstDayBackward"
	T_FINDDELETED         Type = "findFirstDeletedTime"
	T_FINDNTH             Type = "findNthTime"
	T_ACCRUED             Type = "accrued"
	T_ACCRUEDBEFORE       Type = "accruedBefore"
	T_BALANCE             Type = "balance"
	T_CALLSQL             Type = "callSql"
	T_LCONSECUTIVERANGE   Type = "longestConsecutiveRange"
	T_FCONSECUTIVEDAY     Type = "firstConsecutiveDay"
	T_LCONSECUTIVEDAY     Type = "lastConsecutiveDay"
	T_CONTAINS            Type = "contains"
	T_CONVERTDTTMTZ       Type = "convertDttmByTimezone"
	T_COUNTGRPCALC        Type = "countGroupCalc"
	T_COUNTHOLIDAYS       Type = "countHolidays"
	T_GETHOLIDAY          Type = "getHoliday"
	T_CNTHOMECREWMEMS     Type = "countHomeCrewMembers"
	T_CNTSHIFTCHANGES     Type = "countShiftChanges"
	T_EMPATTREXISTS       Type = "employee_attribute_exists"
	T_EMPATTR             Type = "employee_attribute"
	T_GETATTRCALCDATE     Type = "getAttributeCalcDate"
	T_GETTORBOOL          Type = "getBooleanFieldFromTor"
	T_GETTORDATE          Type = "getDateFieldFromTor"
	T_GETTORNUM           Type = "getNumberFieldFromTor"
	T_GETPAYCURRCODE      Type = "getPayCurrencyCode"
	T_GETTORSELECT        Type = "getSelectFieldValueFromTor"
	T_GETTORSTR           Type = "getStringFieldFromTor"
	T_GETSYSDATEBYTZ      Type = "getSysDateByTimezone"
	T_INDEXOF             Type = "indexOf"
	T_LDLOOKUP            Type = "LDLookup"
	T_LDVALIDATE          Type = "LDValidate"
	T_LENGTHOFSERVICE     Type = "lengthOfService"
	T_MAKEDATE            Type = "makeDate"
	T_MAKEDTTM            Type = "makeDateTime"
	T_MAKEDTTMRANGE       Type = "makeDateTimeRange"
	T_PAYCODEINSCHEDMAP   Type = "payCodeInScheduleMap"
	T_PAYCODEINTSTMAP     Type = "payCodeInTimeSheetMap"
	T_RANGELOOKUP         Type = "rangeLookup"
	T_ROUND               Type = "round"
	T_ROUNDUP             Type = "roundUp"
	T_ROUNDDOWN           Type = "roundDown"
	T_ROUNDTOINT          Type = "roundToInt"
	T_SEMIMONTHLYPERIOD   Type = "semiMonthlyPeriod"
	T_SUBSTR              Type = "substr"
	T_TOLOWER             Type = "toLowerCase"
	T_TOUPPER             Type = "toUpperCase"
	T_MINSCHED            Type = "minSchedule"
	T_MAXSCHED            Type = "maxSchedule"
	T_AVGSCHED            Type = "averageSchedule"
	T_MINTIME             Type = "minTime"
	T_MAXTIME             Type = "maxTime"
	T_AVGTIME             Type = "averageTime"
	T_SWIPEINLATINRANGE   Type = "Swipe_in_latitude_in_range"
	T_SWIPEINLONGINRANGE  Type = "Swipe_in_longitude_in_range"
	T_SWIPEOUTLATINRANGE  Type = "Swipe_out_latitude_in_range"
	T_SWIPEOUTLONGINRANGE Type = "Swipe_out_longitude_in_range"
	T_SUMEXCEPT           Type = "sumException"
	T_MINEXCEPT           Type = "minException"
	T_MAXEXCEPT           Type = "maxException"
	T_AVGEXCEPT           Type = "averageException"
)
