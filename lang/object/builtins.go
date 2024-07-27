package object

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
			ReturnType: T_BOOL,
			Params: []Param{
				{Name: "condition", Types: pT{T_BOOL}},
				{Name: "then", Types: pT{T_ANY}},
				{Name: "else", Types: pT{T_ANY}},
			},
		},

		Min: {
			Name:       Min,
			ReturnType: T_NUMBER,
			Params: []Param{
				{Name: "args", Types: pT{T_NUMBER}, List: true},
			},
		},

		Max: {
			Name:       Max,
			ReturnType: T_NUMBER,
			Params: []Param{
				{Name: "args", Types: pT{T_NUMBER}, List: true},
			},
		},

		Contains: {
			Name:       Contains,
			ReturnType: T_BOOL,
			Params: []Param{
				{Name: "x", Types: pT{T_STRING}},
				{Name: "y", Types: pT{T_STRING}},
			},
		},

		Sum: {
			Name:       Sum,
			ReturnType: T_NUMBER,
			Params: []Param{
				{Name: "interval", Types: pT{T_DAY, T_WEEK, T_PERIOD}},
				{Name: "range", Types: pT{T_DAY, T_WEEK, T_PERIOD, T_DATERNG}},
				{Name: "aliasName", Types: pT{T_STRING}, Optional: true},
				{Name: "expression", Types: pT{T_NUMBER}},
				{Name: "condition", Types: pT{T_BOOL}, Optional: true},
			},
		},

		Count: {
			Name:       Count,
			ReturnType: T_NUMBER,
			Params: []Param{
				{Name: "interval", Types: pT{T_DAY, T_WEEK, T_PERIOD}},
				{Name: "range", Types: pT{T_DAY, T_WEEK, T_PERIOD, T_DATERNG}},
				{Name: "aliasName", Types: pT{T_STRING}, Optional: true},
				{Name: "condition", Types: pT{T_BOOL}},
			},
		},

		SumTime:                    {Name: SumTime, ReturnType: T_NUMBER},
		CountTime:                  {Name: CountTime, ReturnType: T_NUMBER},
		FindFirstTime:              {Name: FindFirstTime, ReturnType: T_NUMBER},
		SumSchedule:                {Name: SumSchedule, ReturnType: T_NUMBER},
		CountSchedule:              {Name: CountSchedule, ReturnType: T_NUMBER},
		FindFirstSchedule:          {Name: FindFirstSchedule, ReturnType: T_SCHEDREC},
		CountException:             {Name: CountException, ReturnType: T_NUMBER},
		FindFirstTorDetail:         {Name: FindFirstTorDetail, ReturnType: T_TORDTL},
		FindFirstDayForward:        {Name: FindFirstDayForward, ReturnType: T_DATE},
		FindFirstDayBackward:       {Name: FindFirstDayBackward, ReturnType: T_DATE},
		FindFirstDeletedTime:       {Name: FindFirstDeletedTime, ReturnType: T_DATE},
		LongestConsecutiveRange:    {Name: LongestConsecutiveRange, ReturnType: T_DATERNG},
		FirstConsecutiveDay:        {Name: FirstConsecutiveDay, ReturnType: T_DATE},
		LastConsecutiveDay:         {Name: LastConsecutiveDay, ReturnType: T_DATE},
		FindNthTime:                {Name: FindNthTime, ReturnType: T_TIME},
		Accrued:                    {Name: Accrued, ReturnType: T_NUMBER},
		BalanceAccruedBefore:       {Name: BalanceAccruedBefore, ReturnType: T_NUMBER},
		Balance:                    {Name: Balance, ReturnType: T_NUMBER},
		CallSql:                    {Name: CallSql, ReturnType: T_RESULTSET},
		ConvertDttmByTimezone:      {Name: ConvertDttmByTimezone, ReturnType: T_DTTM},
		CountGroupCalc:             {Name: CountGroupCalc, ReturnType: T_NUMBER},
		CountHolidays:              {Name: CountHolidays, ReturnType: T_NUMBER},
		GetHoliday:                 {Name: GetHoliday, ReturnType: T_STRING},
		CountHomeCrewMembers:       {Name: CountHomeCrewMembers, ReturnType: T_NUMBER},
		CountShiftChanges:          {Name: CountShiftChanges, ReturnType: T_NUMBER},
		EmployeeAttributeExists:    {Name: EmployeeAttributeExists, ReturnType: T_BOOL},
		EmployeeAttribute:          {Name: EmployeeAttribute, ReturnType: T_EMPATTR},
		GetAttributeCalcDate:       {Name: GetAttributeCalcDate, ReturnType: T_DATE},
		GetBooleanFieldFromTor:     {Name: GetBooleanFieldFromTor, ReturnType: T_BOOL},
		GetDateFieldFromTor:        {Name: GetDateFieldFromTor, ReturnType: T_DATE},
		GetNumberFieldFromTor:      {Name: GetNumberFieldFromTor, ReturnType: T_NUMBER},
		GetPayCurrencyCode:         {Name: GetPayCurrencyCode, ReturnType: T_STRING},
		GetSelectFieldValueFromTor: {Name: GetSelectFieldValueFromTor, ReturnType: T_STRING},
		GetStringFieldFromTor:      {Name: GetStringFieldFromTor, ReturnType: T_STRING},
		GetSysDateByTimezone:       {Name: GetSysDateByTimezone, ReturnType: T_IDENT},
		LdLookup:                   {Name: LdLookup, ReturnType: T_LDREC},
		LdValidate:                 {Name: LdValidate, ReturnType: T_BOOL},
		IndexOf:                    {Name: IndexOf, ReturnType: T_NUMBER},
		LengthOfService:            {Name: LengthOfService, ReturnType: T_NUMBER},
		MakeDate:                   {Name: MakeDate, ReturnType: T_DATE},
		MakeDateTime:               {Name: MakeDateTime, ReturnType: T_DTTM},
		MakeDateTimeRange:          {Name: MakeDateTimeRange, ReturnType: T_DTTMRNG},
		PayCodeInScheduleMap:       {Name: PayCodeInScheduleMap, ReturnType: T_BOOL},
		PayCodeInTimeSheetMap:      {Name: PayCodeInTimeSheetMap, ReturnType: T_BOOL},
		RangeLookup:                {Name: RangeLookup, ReturnType: T_NUMBER},
		Round:                      {Name: Round, ReturnType: T_NUMBER},
		RoundUp:                    {Name: RoundUp, ReturnType: T_NUMBER},
		RoundDown:                  {Name: RoundDown, ReturnType: T_NUMBER},
		RoundToInt:                 {Name: RoundToInt, ReturnType: T_NUMBER},
		SemiMonthlyPeriod:          {Name: SemiMonthlyPeriod, ReturnType: T_PERIOD},
		Substr:                     {Name: Substr, ReturnType: T_STRING},
		ToLowerCase:                {Name: ToLowerCase, ReturnType: T_STRING},
		ToUpperCase:                {Name: ToUpperCase, ReturnType: T_STRING},
		MinSchedule:                {Name: MinSchedule, ReturnType: T_NUMBER},
		MaxSchedule:                {Name: MaxSchedule, ReturnType: T_NUMBER},
		AvgSchedule:                {Name: AvgSchedule, ReturnType: T_NUMBER},
		MinTime:                    {Name: MinTime, ReturnType: T_NUMBER},
		MaxTime:                    {Name: MaxTime, ReturnType: T_NUMBER},
		AvgTime:                    {Name: AvgTime, ReturnType: T_NUMBER},
		SumException:               {Name: SumException, ReturnType: T_NUMBER},
		MinException:               {Name: MinException, ReturnType: T_NUMBER},
		MaxException:               {Name: MaxException, ReturnType: T_NUMBER},
		AvgException:               {Name: AvgException, ReturnType: T_NUMBER},
	}
}
