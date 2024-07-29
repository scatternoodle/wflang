package object

import (
	"time"

	"github.com/scatternoodle/wflang/lang/token"
	"github.com/scatternoodle/wflang/lang/types"
)

// Object is the base interface for any formula object in WFLang, and is the ultimate
// expression of a formula. The main requirement of an object is to have a type.
type Object interface {
	Type() types.Type
	Methods() []Function
	OK() bool
}

type Number struct {
	Static bool
	Val    float64
}

func (n Number) Type() types.Type    { return types.T_NUMBER }
func (n Number) Methods() []Function { return nil }
func (n Number) OK() bool            { return false }

type Variable struct {
	Name string
	Val  Object
	Pos  token.Pos
}

func (v Variable) Methods() []Function { return nil }
func (v Variable) Type() types.Type    { return v.Val.Type() }
func (v Variable) OK() bool            { return false }

func (v Variable) String() string {
	return "var " + v.Name + ": " + string(v.Type())
}

type String struct {
	Static bool
	Val    string
}

func (s String) Type() types.Type    { return types.T_STRING }
func (s String) Methods() []Function { return nil }
func (s String) OK() bool            { return s.Static }

type Ident struct {
	Static bool
	Val    string
}

func (i Ident) Type() types.Type    { return types.T_IDENT }
func (i Ident) Methods() []Function { return nil }
func (i Ident) OK() bool            { return i.Static }

type Time struct {
	Static bool
	Val    time.Time
}

func (t Time) Type() types.Type    { return types.T_TIME }
func (t Time) Methods() []Function { return nil }
func (t Time) OK() bool            { return t.Static }

type DateTime struct {
	Static bool
	Val    time.Time
}

func (d DateTime) Type() types.Type    { return types.T_DTTM }
func (d DateTime) Methods() []Function { return nil }
func (d DateTime) OK() bool            { return d.Static && !d.Val.IsZero() }

type DateTimeRange struct {
	Static bool
	Val    struct {
		Start time.Time
		End   time.Time
	}
}

func (r DateTimeRange) Type() types.Type    { return types.T_DTTMRNG }
func (r DateTimeRange) Methods() []Function { return nil }
func (r DateTimeRange) OK() bool            { return r.Static && !r.Val.Start.IsZero() && !r.Val.End.IsZero() }

type Date struct {
	Static bool
	Val    time.Time
}

func (d Date) Type() types.Type    { return types.T_DATE }
func (d Date) Methods() []Function { return nil }
func (d Date) OK() bool            { return d.Static && !d.Val.IsZero() }

type DateRange struct {
	Static bool
	Val    struct {
		Start time.Time
		End   time.Time
	}
}

func (d DateRange) Type() types.Type    { return types.T_DATERNG }
func (d DateRange) Methods() []Function { return nil }
func (d DateRange) OK() bool            { return d.Static && !d.Val.Start.IsZero() && !d.Val.End.IsZero() }

type Boolean struct {
	Static bool
	Val    bool
}

func (b Boolean) Type() types.Type    { return types.T_BOOL }
func (b Boolean) Methods() []Function { return nil }
func (b Boolean) OK() bool            { return b.Static }

type ScheduleRecord struct{}

func (s ScheduleRecord) Type() types.Type    { return types.T_SCHEDREC }
func (s ScheduleRecord) Methods() []Function { return nil }
func (s ScheduleRecord) OK() bool            { return false }

type TimeRecord struct{}

func (t TimeRecord) Type() types.Type    { return types.T_TIMEREC }
func (t TimeRecord) Methods() []Function { return nil }
func (t TimeRecord) OK() bool            { return false }

type Attribute struct {
	Name    string
	AttType types.Type
}

func (a Attribute) Type() types.Type    { return types.T_EMPATTR }
func (a Attribute) Methods() []Function { return nil }
func (a Attribute) OK() bool            { return false }

type LDRecord struct{}

func (l LDRecord) Type() types.Type    { return types.T_LDREC }
func (l LDRecord) Methods() []Function { return nil }
func (l LDRecord) OK() bool            { return false }

type TORDetailRecord struct{}

func (t TORDetailRecord) Type() types.Type    { return types.T_TORDTL }
func (t TORDetailRecord) Methods() []Function { return nil }
func (t TORDetailRecord) OK() bool            { return false }

type ResultSet struct {
	Columns []struct {
		Name     string
		DataType types.Type
	}
}

func (r ResultSet) Type() types.Type    { return types.T_RESULTSET }
func (r ResultSet) Methods() []Function { return nil }
func (r ResultSet) OK() bool            { return false }

type TimeRecordGroup struct{}

func (t TimeRecordGroup) Type() types.Type    { return types.T_TRGROUP }
func (t TimeRecordGroup) Methods() []Function { return nil }
func (t TimeRecordGroup) OK() bool            { return false }

type Exception struct{}

func (s Exception) Type() types.Type    { return types.T_EXCEPTION }
func (s Exception) Methods() []Function { return nil }
func (s Exception) OK() bool            { return false }

type Day struct{}

func (d Day) Type() types.Type    { return types.T_DAY }
func (d Day) Methods() []Function { return nil }
func (d Day) OK() bool            { return false }

type Week struct{}

func (w Week) Type() types.Type    { return types.T_WEEK }
func (w Week) Methods() []Function { return nil }
func (w Week) OK() bool            { return false }

type Period struct{}

func (p Period) Type() types.Type    { return types.T_PERIOD }
func (p Period) Methods() []Function { return nil }
func (p Period) OK() bool            { return false }

type Null struct{}

func (n Null) Type() types.Type    { return types.T_NULL }
func (n Null) Methods() []Function { return nil }
func (n Null) OK() bool            { return false }

type Any struct{}

func (a Any) Type() types.Type    { return types.T_ANY }
func (a Any) Methods() []Function { return nil }
func (a Any) OK() bool            { return false }
