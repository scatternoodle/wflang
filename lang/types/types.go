package types

type BaseType string

const (
	NUMBER    BaseType = "number"
	STRING    BaseType = "string"
	IDENT     BaseType = "ident"
	TIME      BaseType = "time"
	DTTM      BaseType = "datetime"
	DTTMRNG   BaseType = "datetime range"
	DATE      BaseType = "date"
	DATERNG   BaseType = "date range"
	BOOL      BaseType = "boolean"
	SCHEDREC  BaseType = "schedule record"
	TIMEREC   BaseType = "time record"
	EMPATTR   BaseType = "employee attribute"
	LDREC     BaseType = "LD record"
	TORDTL    BaseType = "TOR detail record"
	OBJECT    BaseType = "object" // misc type for things like SQL results
	POLICY_ID BaseType = "policy ID"
	PERIOD    BaseType = "period"
	TRGROUP   BaseType = "time record group"
	EXCEPTION BaseType = "exception"
)

// IsNullable returns true if the given WFLang base type can be null (and therefore
// requires nullchecking in formulae).
func IsNullable(t BaseType) bool {
	switch t {
	case STRING, SCHEDREC, TIMEREC, EMPATTR, LDREC, TORDTL, DATE, TIME, DTTM,
		DATERNG, DTTMRNG, TRGROUP, EXCEPTION:

		return true
	}
	return false
}
