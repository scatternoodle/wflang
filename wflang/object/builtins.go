package object

import "github.com/scatternoodle/wflang/wflang/types"

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
