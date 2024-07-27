package builtins

import "github.com/scatternoodle/wflang/lang/types"

type Builtin struct {
	Name       string
	ReturnType types.BaseType
	Params     []types.BaseType
}

func Builtins() map[string]Builtin {
	return map[string]Builtin{
		If:                         {Name: If, ReturnType: types.BOOL},
		Min:                        {Name: Min, ReturnType: types.NUMBER},
		Max:                        {Name: Max, ReturnType: types.NUMBER},
		Contains:                   {Name: Contains, ReturnType: types.BOOL},
		Sum:                        {Name: Sum, ReturnType: types.NUMBER},
		Count:                      {Name: Count, ReturnType: types.NUMBER},
		SumTime:                    {Name: SumTime, ReturnType: types.NUMBER},
		CountTime:                  {Name: CountTime, ReturnType: types.NUMBER},
		FindFirstTime:              {Name: FindFirstTime, ReturnType: types.NUMBER},
		SumSchedule:                {Name: SumSchedule, ReturnType: types.NUMBER},
		CountSchedule:              {Name: CountSchedule, ReturnType: types.NUMBER},
		FindFirstSchedule:          {Name: FindFirstSchedule, ReturnType: types.SCHEDREC},
		CountException:             {Name: CountException, ReturnType: types.NUMBER},
		FindFirstTorDetail:         {Name: FindFirstTorDetail, ReturnType: types.TORDTL},
		FindFirstDayForward:        {Name: FindFirstDayForward, ReturnType: types.DATE},
		FindFirstDayBackward:       {Name: FindFirstDayBackward, ReturnType: types.DATE},
		FindFirstDeletedTime:       {Name: FindFirstDeletedTime, ReturnType: types.DATE},
		LongestConsecutiveRange:    {Name: LongestConsecutiveRange, ReturnType: types.DATERNG},
		FirstConsecutiveDay:        {Name: FirstConsecutiveDay, ReturnType: types.DATE},
		LastConsecutiveDay:         {Name: LastConsecutiveDay, ReturnType: types.DATE},
		FindNthTime:                {Name: FindNthTime, ReturnType: types.TIME},
		Accrued:                    {Name: Accrued, ReturnType: types.NUMBER},
		BalanceAccruedBefore:       {Name: BalanceAccruedBefore, ReturnType: types.NUMBER},
		Balance:                    {Name: Balance, ReturnType: types.NUMBER},
		CallSql:                    {Name: CallSql, ReturnType: types.OBJECT},
		ConvertDttmByTimezone:      {Name: ConvertDttmByTimezone, ReturnType: types.DTTM},
		CountGroupCalc:             {Name: CountGroupCalc, ReturnType: types.NUMBER},
		CountHolidays:              {Name: CountHolidays, ReturnType: types.NUMBER},
		GetHoliday:                 {Name: GetHoliday, ReturnType: types.STRING},
		CountHomeCrewMembers:       {Name: CountHomeCrewMembers, ReturnType: types.NUMBER},
		CountShiftChanges:          {Name: CountShiftChanges, ReturnType: types.NUMBER},
		EmployeeAttributeExists:    {Name: EmployeeAttributeExists, ReturnType: types.BOOL},
		EmployeeAttribute:          {Name: EmployeeAttribute, ReturnType: types.EMPATTR},
		GetAttributeCalcDate:       {Name: GetAttributeCalcDate, ReturnType: types.DATE},
		GetBooleanFieldFromTor:     {Name: GetBooleanFieldFromTor, ReturnType: types.BOOL},
		GetDateFieldFromTor:        {Name: GetDateFieldFromTor, ReturnType: types.DATE},
		GetNumberFieldFromTor:      {Name: GetNumberFieldFromTor, ReturnType: types.NUMBER},
		GetPayCurrencyCode:         {Name: GetPayCurrencyCode, ReturnType: types.STRING},
		GetSelectFieldValueFromTor: {Name: GetSelectFieldValueFromTor, ReturnType: types.STRING},
		GetStringFieldFromTor:      {Name: GetStringFieldFromTor, ReturnType: types.STRING},
		GetSysDateByTimezone:       {Name: GetSysDateByTimezone, ReturnType: types.POLICY_ID},
		LdLookup:                   {Name: LdLookup, ReturnType: types.LDREC},
		LdValidate:                 {Name: LdValidate, ReturnType: types.BOOL},
		IndexOf:                    {Name: IndexOf, ReturnType: types.NUMBER},
		LengthOfService:            {Name: LengthOfService, ReturnType: types.NUMBER},
		MakeDate:                   {Name: MakeDate, ReturnType: types.DATE},
		MakeDateTime:               {Name: MakeDateTime, ReturnType: types.DTTM},
		MakeDateTimeRange:          {Name: MakeDateTimeRange, ReturnType: types.DTTMRNG},
		PayCodeInScheduleMap:       {Name: PayCodeInScheduleMap, ReturnType: types.BOOL},
		PayCodeInTimeSheetMap:      {Name: PayCodeInTimeSheetMap, ReturnType: types.BOOL},
		RangeLookup:                {Name: RangeLookup, ReturnType: types.NUMBER},
		Round:                      {Name: Round, ReturnType: types.NUMBER},
		RoundUp:                    {Name: RoundUp, ReturnType: types.NUMBER},
		RoundDown:                  {Name: RoundDown, ReturnType: types.NUMBER},
		RoundToInt:                 {Name: RoundToInt, ReturnType: types.NUMBER},
		SemiMonthlyPeriod:          {Name: SemiMonthlyPeriod, ReturnType: types.PERIOD},
		Substr:                     {Name: Substr, ReturnType: types.STRING},
		ToLowerCase:                {Name: ToLowerCase, ReturnType: types.STRING},
		ToUpperCase:                {Name: ToUpperCase, ReturnType: types.STRING},
		MinSchedule:                {Name: MinSchedule, ReturnType: types.NUMBER},
		MaxSchedule:                {Name: MaxSchedule, ReturnType: types.NUMBER},
		AvgSchedule:                {Name: AvgSchedule, ReturnType: types.NUMBER},
		MinTime:                    {Name: MinTime, ReturnType: types.NUMBER},
		MaxTime:                    {Name: MaxTime, ReturnType: types.NUMBER},
		AvgTime:                    {Name: AvgTime, ReturnType: types.NUMBER},
		SumException:               {Name: SumException, ReturnType: types.NUMBER},
		MinException:               {Name: MinException, ReturnType: types.NUMBER},
		MaxException:               {Name: MaxException, ReturnType: types.NUMBER},
		AverageException:           {Name: AverageException, ReturnType: types.NUMBER},
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
	AverageException           string = "averageexception"
)
