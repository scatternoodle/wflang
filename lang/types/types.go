package types

type Type string

const (
	T_ANY       Type = "any"
	T_NUMBER    Type = "number"
	T_STRING    Type = "string"
	T_IDENT     Type = "ident"
	T_TIME      Type = "time"
	T_DTTM      Type = "dateTime"
	T_DTTMRNG   Type = "dateTimeRange"
	T_DATE      Type = "date"
	T_DATERNG   Type = "dateRange"
	T_BOOL      Type = "boolean"
	T_SCHEDREC  Type = "scheduleRecord"
	T_TIMEREC   Type = "timeRecord"
	T_EMPATTR   Type = "employeeAttribute"
	T_LDREC     Type = "LDRecord"
	T_TORDTL    Type = "TORDetailRecord"
	T_RESULTSET Type = "resultSet"
	T_TRGROUP   Type = "timeRecordGroup"
	T_EXCEPTION Type = "exception"
	T_DAY       Type = "day"
	T_WEEK      Type = "week"
	T_PERIOD    Type = "period"
	T_NULL      Type = "null"
)

// IsNullable returns true if the given WFLang base type can be null (and therefore
// requires nullchecking in formulae).
func IsNullable(t Type) bool {
	switch t {
	case T_STRING, T_SCHEDREC, T_TIMEREC, T_EMPATTR, T_LDREC, T_TORDTL, T_DATE, T_TIME, T_DTTM,
		T_DATERNG, T_DTTMRNG, T_TRGROUP, T_EXCEPTION:

		return true
	}
	return false
}
