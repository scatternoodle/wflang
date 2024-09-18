package object

import (
	"log/slog"
	"strings"

	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/server/docstring"
	"github.com/scatternoodle/wflang/wflang/types"
)

// Builtin wraps a simple map lookup against Builtins(), making the check
// case-insensitive.
func Builtin(name string) (b Function, ok bool) {
	slog.Debug("called with name=" + name)
	b, ok = Builtins()[strings.ToLower(name)]
	slog.Debug("got", "b", b, "ok", ok)
	return
}

func Builtins() map[string]Function {
	return map[string]Function{
		If: {
			Name:       If,
			ReturnType: types.T_BOOL,
			Params: []Param{
				{Name: "condition", Types: pTypes{types.T_BOOL}},
				{Name: "then", Types: pTypes{types.T_ANY}},
				{Name: "else", Types: pTypes{types.T_ANY}},
			},
		},

		Min: {
			Name:       Min,
			ReturnType: types.T_NUMBER,
			Params: []Param{
				{Name: "args", Types: pTypes{types.T_NUMBER}, List: true},
			},
		},

		Max: {
			Name:       Max,
			ReturnType: types.T_NUMBER,
			Params: []Param{
				{Name: "args", Types: pTypes{types.T_NUMBER}, List: true},
			},
		},

		Contains: {
			Name:       Contains,
			ReturnType: types.T_BOOL,
			Params: []Param{
				{Name: "x", Types: pTypes{types.T_STRING}},
				{Name: "y", Types: pTypes{types.T_STRING}},
			},
		},

		Sum: {
			Name:       Sum,
			ReturnType: types.T_NUMBER,
			Params: []Param{
				{Name: "interval", Types: pTypes{types.T_DAY, types.T_WEEK, types.T_PERIOD}},
				{Name: "range", Types: pTypes{types.T_DAY, types.T_WEEK, types.T_PERIOD, types.T_DATERNG}},
				{Name: "aliasName", Types: pTypes{types.T_STRING}, Optional: true},
				{Name: "expression", Types: pTypes{types.T_NUMBER}},
				{Name: "condition", Types: pTypes{types.T_BOOL}, Optional: true},
			},
		},

		Count: {
			Name:       Count,
			ReturnType: types.T_NUMBER,
			Params: []Param{
				{Name: "interval", Types: pTypes{types.T_DAY, types.T_WEEK, types.T_PERIOD}},
				{Name: "range", Types: pTypes{types.T_DAY, types.T_WEEK, types.T_PERIOD, types.T_DATERNG}},
				{Name: "aliasName", Types: pTypes{types.T_STRING}, Optional: true},
				{Name: "condition", Types: pTypes{types.T_BOOL}},
			},
		},

		SumTime:                    {Name: SumTime, ReturnType: types.T_NUMBER},
		CountTime:                  {Name: CountTime, ReturnType: types.T_NUMBER},
		FindFirstTime:              {Name: FindFirstTime, ReturnType: types.T_NUMBER},
		SumSchedule:                {Name: SumSchedule, ReturnType: types.T_NUMBER},
		CountSchedule:              {Name: CountSchedule, ReturnType: types.T_NUMBER},
		FindFirstSchedule:          {Name: FindFirstSchedule, ReturnType: types.T_SCHEDREC},
		CountException:             {Name: CountException, ReturnType: types.T_NUMBER},
		FindFirstTorDetail:         {Name: FindFirstTorDetail, ReturnType: types.T_TORDTL},
		FindFirstDayForward:        {Name: FindFirstDayForward, ReturnType: types.T_DATE},
		FindFirstDayBackward:       {Name: FindFirstDayBackward, ReturnType: types.T_DATE},
		FindFirstDeletedTime:       {Name: FindFirstDeletedTime, ReturnType: types.T_DATE},
		LongestConsecutiveRange:    {Name: LongestConsecutiveRange, ReturnType: types.T_DATERNG},
		FirstConsecutiveDay:        {Name: FirstConsecutiveDay, ReturnType: types.T_DATE},
		LastConsecutiveDay:         {Name: LastConsecutiveDay, ReturnType: types.T_DATE},
		FindNthTime:                {Name: FindNthTime, ReturnType: types.T_TIME},
		Accrued:                    {Name: Accrued, ReturnType: types.T_NUMBER},
		BalanceAccruedBefore:       {Name: BalanceAccruedBefore, ReturnType: types.T_NUMBER},
		Balance:                    {Name: Balance, ReturnType: types.T_NUMBER},
		CallSql:                    {Name: CallSql, ReturnType: types.T_RESULTSET},
		ConvertDttmByTimezone:      {Name: ConvertDttmByTimezone, ReturnType: types.T_DTTM},
		CountGroupCalc:             {Name: CountGroupCalc, ReturnType: types.T_NUMBER},
		CountHolidays:              {Name: CountHolidays, ReturnType: types.T_NUMBER},
		GetHoliday:                 {Name: GetHoliday, ReturnType: types.T_STRING},
		CountHomeCrewMembers:       {Name: CountHomeCrewMembers, ReturnType: types.T_NUMBER},
		CountShiftChanges:          {Name: CountShiftChanges, ReturnType: types.T_NUMBER},
		EmployeeAttributeExists:    {Name: EmployeeAttributeExists, ReturnType: types.T_BOOL},
		EmployeeAttribute:          {Name: EmployeeAttribute, ReturnType: types.T_EMPATTR},
		GetAttributeCalcDate:       {Name: GetAttributeCalcDate, ReturnType: types.T_DATE},
		GetBooleanFieldFromTor:     {Name: GetBooleanFieldFromTor, ReturnType: types.T_BOOL},
		GetDateFieldFromTor:        {Name: GetDateFieldFromTor, ReturnType: types.T_DATE},
		GetNumberFieldFromTor:      {Name: GetNumberFieldFromTor, ReturnType: types.T_NUMBER},
		GetPayCurrencyCode:         {Name: GetPayCurrencyCode, ReturnType: types.T_STRING},
		GetSelectFieldValueFromTor: {Name: GetSelectFieldValueFromTor, ReturnType: types.T_STRING},
		GetStringFieldFromTor:      {Name: GetStringFieldFromTor, ReturnType: types.T_STRING},
		GetSysDateByTimezone:       {Name: GetSysDateByTimezone, ReturnType: types.T_IDENT},
		LdLookup:                   {Name: LdLookup, ReturnType: types.T_LDREC},
		LdValidate:                 {Name: LdValidate, ReturnType: types.T_BOOL},
		IndexOf:                    {Name: IndexOf, ReturnType: types.T_NUMBER},
		LengthOfService:            {Name: LengthOfService, ReturnType: types.T_NUMBER},
		MakeDate:                   {Name: MakeDate, ReturnType: types.T_DATE},
		MakeDateTime:               {Name: MakeDateTime, ReturnType: types.T_DTTM},
		MakeDateTimeRange:          {Name: MakeDateTimeRange, ReturnType: types.T_DTTMRNG},
		PayCodeInScheduleMap:       {Name: PayCodeInScheduleMap, ReturnType: types.T_BOOL},
		PayCodeInTimeSheetMap:      {Name: PayCodeInTimeSheetMap, ReturnType: types.T_BOOL},
		RangeLookup:                {Name: RangeLookup, ReturnType: types.T_NUMBER},
		Round:                      {Name: Round, ReturnType: types.T_NUMBER},
		RoundUp:                    {Name: RoundUp, ReturnType: types.T_NUMBER},
		RoundDown:                  {Name: RoundDown, ReturnType: types.T_NUMBER},
		RoundToInt:                 {Name: RoundToInt, ReturnType: types.T_NUMBER},
		SemiMonthlyPeriod:          {Name: SemiMonthlyPeriod, ReturnType: types.T_PERIOD},
		Substr:                     {Name: Substr, ReturnType: types.T_STRING},
		ToLowerCase:                {Name: ToLowerCase, ReturnType: types.T_STRING},
		ToUpperCase:                {Name: ToUpperCase, ReturnType: types.T_STRING},
		MinSchedule:                {Name: MinSchedule, ReturnType: types.T_NUMBER},
		MaxSchedule:                {Name: MaxSchedule, ReturnType: types.T_NUMBER},
		AvgSchedule:                {Name: AvgSchedule, ReturnType: types.T_NUMBER},
		MinTime:                    {Name: MinTime, ReturnType: types.T_NUMBER},
		MaxTime:                    {Name: MaxTime, ReturnType: types.T_NUMBER},
		AvgTime:                    {Name: AvgTime, ReturnType: types.T_NUMBER},
		SumException:               {Name: SumException, ReturnType: types.T_NUMBER},
		MinException:               {Name: MinException, ReturnType: types.T_NUMBER},
		MaxException:               {Name: MaxException, ReturnType: types.T_NUMBER},
		AvgException:               {Name: AvgException, ReturnType: types.T_NUMBER},
	}
}

const (
	If                         string = "if"
	Min                        string = "min"
	Max                        string = "max"
	Contains                   string = "contains"
	Sum                        string = "sum"
	Count                      string = "count"
	SumTime                    string = "sumtime"
	CountTime                  string = "counttime"
	FindFirstTime              string = "findfirsttime"
	SumSchedule                string = "sumschedule"
	CountSchedule              string = "countschedule"
	FindFirstSchedule          string = "findfirstschedule"
	CountException             string = "countexception"
	FindFirstTorDetail         string = "findfirsttordetail"
	FindFirstDayForward        string = "findfirstdayforward"
	FindFirstDayBackward       string = "findfirstdaybackward"
	FindFirstDeletedTime       string = "findfirstdeletedtime"
	LongestConsecutiveRange    string = "longestconsecutiverange"
	FirstConsecutiveDay        string = "firstconsecutiveday"
	LastConsecutiveDay         string = "lastconsecutiveday"
	FindNthTime                string = "findnthtime"
	Accrued                    string = "accrued"
	BalanceAccruedBefore       string = "balanceaccruedbefore"
	Balance                    string = "balance"
	CallSql                    string = "callsql"
	ConvertDttmByTimezone      string = "convertdttmbytimezone"
	CountGroupCalc             string = "countgroupcalc"
	CountHolidays              string = "countholidays"
	GetHoliday                 string = "getholiday"
	CountHomeCrewMembers       string = "counthomecrewmembers"
	CountShiftChanges          string = "countshiftchanges"
	EmployeeAttributeExists    string = "employee_attribute_exists"
	EmployeeAttribute          string = "employee_attribute"
	GetAttributeCalcDate       string = "getattributecalculationdate"
	GetBooleanFieldFromTor     string = "getbooleanfieldfromtor"
	GetDateFieldFromTor        string = "getdatefieldfromtor"
	GetNumberFieldFromTor      string = "getnumberfieldfromtor"
	GetPayCurrencyCode         string = "getpaycurrencycode"
	GetSelectFieldValueFromTor string = "getselectfieldvaluefromtor"
	GetStringFieldFromTor      string = "getstringfieldfromtor"
	GetSysDateByTimezone       string = "getsysdatebytimezone"
	LdLookup                   string = "ldlookup"
	LdValidate                 string = "ldvalidate"
	IndexOf                    string = "indexof"
	LengthOfService            string = "lengthofservice"
	MakeDate                   string = "makedate"
	MakeDateTime               string = "makedatetime"
	MakeDateTimeRange          string = "makedatetimerange"
	PayCodeInScheduleMap       string = "paycodeinschedulemap"
	PayCodeInTimeSheetMap      string = "paycodeintimesheetmap"
	RangeLookup                string = "rangelookup"
	Round                      string = "round"
	RoundUp                    string = "roundup"
	RoundDown                  string = "rounddown"
	RoundToInt                 string = "roundtoint"
	SemiMonthlyPeriod          string = "semimonthlyperiod"
	Substr                     string = "substr"
	ToLowerCase                string = "tolowercase"
	ToUpperCase                string = "touppercase"
	MinSchedule                string = "minschedule"
	MaxSchedule                string = "maxschedule"
	AvgSchedule                string = "avgschedule"
	MinTime                    string = "mintime"
	MaxTime                    string = "maxtime"
	AvgTime                    string = "avgtime"
	SumException               string = "sumexception"
	MinException               string = "minexception"
	MaxException               string = "maxexception"
	AvgException               string = "averageexception"
)

func DocMarkdown(name string) *lsp.MarkupContent {
	content := lsp.MarkupContent{
		Kind:  lsp.MarkupKindMarkdown,
		Value: "",
	}
	if doc, ok := builtinDocString(name); ok {
		content.Value = doc
	}
	return &content
}

func builtinDocString(name string) (text string, ok bool) {
	switch name {
	case Min:
		return docstring.Min, true
	case Max:
		return docstring.Max, true
	case Contains:
		return docstring.Contains, true
	case Sum:
		return docstring.Sum, true
	case Count:
		return docstring.Count, true
	case If:
		return docstring.If, true
	case SumTime:
		return docstring.SumTime, true
	case CountTime:
		return docstring.CountTime, true
	case FindFirstTime:
		return docstring.FindFirstTime, true
	case SumSchedule:
		return docstring.SumSchedule, true
	case CountSchedule:
		return docstring.CountSchedule, true
	case FindFirstSchedule:
		return docstring.FindFirstSchedule, true
	case CountException:
		return docstring.CountException, true
	case FindFirstTorDetail:
		return docstring.FindFirstTorDetail, true
	case FindFirstDayForward:
		return docstring.FindFirstDayForward, true
	case FindFirstDayBackward:
		return docstring.FindFirstDayBackward, true
	case FindFirstDeletedTime:
		return docstring.FindFirstDeletedTime, true
	case LongestConsecutiveRange:
		return docstring.LongestConsecutiveRange, true
	case FirstConsecutiveDay:
		return docstring.FirstConsecutiveDay, true
	case LastConsecutiveDay:
		return docstring.LastConsecutiveDay, true
	case FindNthTime:
		return docstring.FindNthTime, true
	case Accrued:
		return docstring.Accrued, true
	case BalanceAccruedBefore:
		return docstring.BalanceAccruedBefore, true
	case ConvertDttmByTimezone:
		return docstring.ConvertDttmByTimezone, true
	case CountHolidays:
		return docstring.CountHolidays, true
	case GetHoliday:
		return docstring.GetHoliday, true
	case EmployeeAttributeExists:
		return docstring.Employee_attribute_exists, true
	case EmployeeAttribute:
		return docstring.Employee_attribute, true
	case GetAttributeCalcDate:
		return docstring.GetAttributeCalculationDate, true
	case GetBooleanFieldFromTor:
		return docstring.GetBooleanFieldFromTor, true
	case GetDateFieldFromTor:
		return docstring.GetDateFieldFromTor, true
	case GetSelectFieldValueFromTor:
		return docstring.GetSelectFieldValueFromTor, true
	case GetNumberFieldFromTor:
		return docstring.GetNumberFieldFromTor, true
	case GetStringFieldFromTor:
		return docstring.GetStringValueFromTor, true
	case GetSysDateByTimezone:
		return docstring.GetSysDateByTimezone, true
	case LdLookup:
		return docstring.LDLookup, true
	case LdValidate:
		return docstring.LDValidate, true
	case IndexOf:
		return docstring.IndexOf, true
	case GetPayCurrencyCode:
		return docstring.GetPayCurrencyCode, true
	case CallSql:
		return docstring.CallSQL, true
	case LengthOfService:
		return docstring.LengthOfService, true
	case MakeDate:
		return docstring.MakeDate, true
	case MakeDateTime:
		return docstring.MakeDateTime, true
	case MakeDateTimeRange:
		return docstring.MakeDateTimeRange, true
	case PayCodeInScheduleMap:
		return docstring.PayCodeInScheduleMap, true
	case PayCodeInTimeSheetMap:
		return docstring.PayCodeInTimeSheetMap, true
	case Round:
		return docstring.Round, true
	case RoundDown:
		return docstring.RoundDown, true
	case RoundUp:
		return docstring.RoundUp, true
	case RoundToInt:
		return docstring.RoundToInt, true
	case SemiMonthlyPeriod:
		return docstring.SemiMonthlyPeriod, true
	case Substr:
		return docstring.Substr, true
	case ToLowerCase:
		return docstring.ToLowerCase, true
	case ToUpperCase:
		return docstring.ToUpperCase, true
	case MinSchedule:
		return docstring.MinSchedule, true
	case MaxSchedule:
		return docstring.MaxSchedule, true
	case AvgSchedule:
		return docstring.AvgSchedule, true
	case MinTime:
		return docstring.MinTime, true
	case MaxTime:
		return docstring.MaxTime, true
	case AvgTime:
		return docstring.AvgTime, true
	case SumException:
		return docstring.SumException, true
	case MinException:
		return docstring.MinException, true
	case MaxException:
		return docstring.MaxException, true
	case AvgException:
		return docstring.AvgException, true
	}

	return "", false
}
