package types

type BaseType string

const (
	NUMBER    BaseType = "number"
	STRING    BaseType = "string"
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
)
