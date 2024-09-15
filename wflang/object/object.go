package object

import (
	"time"

	"github.com/scatternoodle/wflang/wflang/ast"
	"github.com/scatternoodle/wflang/wflang/types"
)

// Object is the base interface for any formula object in WFLang, and is the ultimate
// expression of a formula. The main requirement of an object is to have a type.
type Object interface {
	Type() types.Type
	Methods() []Function
	Value() (v any, ok bool)
}

// Undefined is used when an object type cannot be resolved, and likely indicates
// a syntax or semantic error in the WFLang code.
type Undefined struct {
	Val ast.Node
}

func (u Undefined) Type() types.Type        { return types.T_UNDEFINED }
func (u Undefined) Methods() []Function     { return nil }
func (u Undefined) Value() (v any, ok bool) { return u.Val, u.Val != nil }

type Number struct {
	Static bool
	Val    float64
}

func (n Number) Type() types.Type        { return types.T_NUMBER }
func (n Number) Methods() []Function     { return nil }
func (n Number) Value() (v any, ok bool) { return n.Val, n.Static }

type Variable struct {
	Name      string
	Statement *ast.VarStatement
	Val       Object
}

func (vr Variable) Methods() []Function { return nil }
func (vr Variable) Type() types.Type    { return vr.Val.Type() }

func (vr Variable) Value() (v any, ok bool) {
	if v, ok = vr.Val.Value(); !ok {
		return nil, false
	}
	return v, true
}

func (v Variable) String() string {
	return "var " + v.Name + ": " + string(v.Type())
}

type String struct {
	Static bool
	Val    string
}

func (s String) Type() types.Type        { return types.T_STRING }
func (s String) Methods() []Function     { return nil }
func (s String) Value() (v any, ok bool) { return s.Val, s.Static }

type Ident struct {
	Static bool
	Val    string
}

func (i Ident) Type() types.Type        { return types.T_IDENT }
func (i Ident) Methods() []Function     { return nil }
func (i Ident) Value() (v any, ok bool) { return i.Val, i.Static }

type Time struct {
	Static bool
	Val    time.Time
}

func (t Time) Type() types.Type        { return types.T_TIME }
func (t Time) Methods() []Function     { return nil }
func (t Time) Value() (v any, ok bool) { return t.Val, t.Static }

type DateTime struct {
	Static bool
	Val    time.Time
}

func (d DateTime) Type() types.Type        { return types.T_DTTM }
func (d DateTime) Methods() []Function     { return nil }
func (d DateTime) Value() (v any, ok bool) { return d.Val, d.Static && !d.Val.IsZero() }

type DateTimeRange struct {
	Static bool
	Val    struct {
		Start time.Time
		End   time.Time
	}
}

func (r DateTimeRange) Type() types.Type    { return types.T_DTTMRNG }
func (r DateTimeRange) Methods() []Function { return nil }
func (r DateTimeRange) Value() (v any, ok bool) {
	return r.Val, r.Static && !r.Val.Start.IsZero() && !r.Val.End.IsZero()
}

type Date struct {
	Static bool
	Val    time.Time
}

func (d Date) Type() types.Type        { return types.T_DATE }
func (d Date) Methods() []Function     { return nil }
func (d Date) Value() (v any, ok bool) { return d.Val, d.Static && !d.Val.IsZero() }

type DateRange struct {
	Static bool
	Val    struct {
		Start time.Time
		End   time.Time
	}
}

func (d DateRange) Type() types.Type    { return types.T_DATERNG }
func (d DateRange) Methods() []Function { return nil }
func (d DateRange) Value() (v any, ok bool) {
	return d.Val, d.Static && !d.Val.Start.IsZero() && !d.Val.End.IsZero()
}

type Boolean struct {
	Static bool
	Val    bool
}

func (b Boolean) Type() types.Type        { return types.T_BOOL }
func (b Boolean) Methods() []Function     { return nil }
func (b Boolean) Value() (v any, ok bool) { return b.Val, b.Static }

type ScheduleRecord struct{}

func (s ScheduleRecord) Type() types.Type        { return types.T_SCHEDREC }
func (s ScheduleRecord) Methods() []Function     { return nil }
func (s ScheduleRecord) Value() (v any, ok bool) { return nil, false }

type TimeRecord struct{}

func (t TimeRecord) Type() types.Type        { return types.T_TIMEREC }
func (t TimeRecord) Methods() []Function     { return nil }
func (t TimeRecord) Value() (v any, ok bool) { return nil, false }

type Attribute struct {
	Name    string
	AttType types.Type
}

func (a Attribute) Type() types.Type        { return types.T_EMPATTR }
func (a Attribute) Methods() []Function     { return nil }
func (a Attribute) Value() (v any, ok bool) { return nil, false }

type LDRecord struct{}

func (l LDRecord) Type() types.Type        { return types.T_LDREC }
func (l LDRecord) Methods() []Function     { return nil }
func (l LDRecord) Value() (v any, ok bool) { return nil, false }

type TORDetailRecord struct{}

func (t TORDetailRecord) Type() types.Type        { return types.T_TORDTL }
func (t TORDetailRecord) Methods() []Function     { return nil }
func (t TORDetailRecord) Value() (v any, ok bool) { return nil, false }

type ResultSet struct {
	Columns []struct {
		Name     string
		DataType types.Type
	}
}

func (r ResultSet) Type() types.Type        { return types.T_RESULTSET }
func (r ResultSet) Methods() []Function     { return nil }
func (r ResultSet) Value() (v any, ok bool) { return nil, false }

type TimeRecordGroup struct{}

func (t TimeRecordGroup) Type() types.Type        { return types.T_TRGROUP }
func (t TimeRecordGroup) Methods() []Function     { return nil }
func (t TimeRecordGroup) Value() (v any, ok bool) { return nil, false }

type Exception struct{}

func (s Exception) Type() types.Type        { return types.T_EXCEPTION }
func (s Exception) Methods() []Function     { return nil }
func (s Exception) Value() (v any, ok bool) { return nil, false }

type Day struct{}

func (d Day) Type() types.Type        { return types.T_DAY }
func (d Day) Methods() []Function     { return nil }
func (d Day) Value() (v any, ok bool) { return nil, false }

type Week struct{}

func (w Week) Type() types.Type        { return types.T_WEEK }
func (w Week) Methods() []Function     { return nil }
func (w Week) Value() (v any, ok bool) { return nil, false }

type Period struct{}

func (p Period) Type() types.Type        { return types.T_PERIOD }
func (p Period) Methods() []Function     { return nil }
func (p Period) Value() (v any, ok bool) { return nil, false }

type Null struct{}

func (n Null) Type() types.Type        { return types.T_NULL }
func (n Null) Methods() []Function     { return nil }
func (n Null) Value() (v any, ok bool) { return nil, false }

type Any struct{}

func (a Any) Type() types.Type        { return types.T_ANY }
func (a Any) Methods() []Function     { return nil }
func (a Any) Value() (v any, ok bool) { return nil, false }
