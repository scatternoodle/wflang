package server

import (
	"strings"

	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/lang/object"
	"github.com/scatternoodle/wflang/lang/token"
	"github.com/scatternoodle/wflang/server/docstring"
)

func (srv *Server) hover(pos lsp.Position) lsp.Hover {
	_, tok, ok := srv.getTokenAtPos(pos)
	if !ok {
		return lsp.Hover{}
	}
	if tok.Type == token.T_BUILTIN {
		return lsp.Hover{MarkupContent: *docMarkdown(strings.ToLower(tok.Literal))}
	}
	return lsp.Hover{}
}

func docMarkdown(name string) *lsp.MarkupContent {
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
	case object.Min:
		return docstring.Min, true
	case object.Max:
		return docstring.Max, true
	case object.Contains:
		return docstring.Contains, true
	case object.Sum:
		return docstring.Sum, true
	case object.Count:
		return docstring.Count, true
	case object.If:
		return docstring.If, true
	case object.SumTime:
		return docstring.SumTime, true
	case object.CountTime:
		return docstring.CountTime, true
	case object.FindFirstTime:
		return docstring.FindFirstTime, true
	case object.SumSchedule:
		return docstring.SumSchedule, true
	case object.CountSchedule:
		return docstring.CountSchedule, true
	case object.FindFirstSchedule:
		return docstring.FindFirstSchedule, true
	case object.CountException:
		return docstring.CountException, true
	case object.FindFirstTorDetail:
		return docstring.FindFirstTorDetail, true
	case object.FindFirstDayForward:
		return docstring.FindFirstDayForward, true
	case object.FindFirstDayBackward:
		return docstring.FindFirstDayBackward, true
	case object.FindFirstDeletedTime:
		return docstring.FindFirstDeletedTime, true
	case object.LongestConsecutiveRange:
		return docstring.LongestConsecutiveRange, true
	case object.FirstConsecutiveDay:
		return docstring.FirstConsecutiveDay, true
	case object.LastConsecutiveDay:
		return docstring.LastConsecutiveDay, true
	case object.FindNthTime:
		return docstring.FindNthTime, true
	case object.Accrued:
		return docstring.Accrued, true
	case object.BalanceAccruedBefore:
		return docstring.BalanceAccruedBefore, true
	case object.ConvertDttmByTimezone:
		return docstring.ConvertDttmByTimezone, true
	case object.CountHolidays:
		return docstring.CountHolidays, true
	case object.GetHoliday:
		return docstring.GetHoliday, true
	case object.EmployeeAttributeExists:
		return docstring.Employee_attribute_exists, true
	case object.EmployeeAttribute:
		return docstring.Employee_attribute, true
	case object.GetAttributeCalcDate:
		return docstring.GetAttributeCalculationDate, true
	case object.GetBooleanFieldFromTor:
		return docstring.GetBooleanFieldFromTor, true
	case object.GetDateFieldFromTor:
		return docstring.GetDateFieldFromTor, true
	case object.GetSelectFieldValueFromTor:
		return docstring.GetSelectFieldValueFromTor, true
	case object.GetNumberFieldFromTor:
		return docstring.GetNumberFieldFromTor, true
	case object.GetStringFieldFromTor:
		return docstring.GetStringValueFromTor, true
	case object.GetSysDateByTimezone:
		return docstring.GetSysDateByTimezone, true
	case object.LdLookup:
		return docstring.LDLookup, true
	case object.LdValidate:
		return docstring.LDValidate, true
	case object.IndexOf:
		return docstring.IndexOf, true
	case object.GetPayCurrencyCode:
		return docstring.GetPayCurrencyCode, true
	case object.CallSql:
		return docstring.CallSQL, true
	case object.LengthOfService:
		return docstring.LengthOfService, true
	case object.MakeDate:
		return docstring.MakeDate, true
	case object.MakeDateTime:
		return docstring.MakeDateTime, true
	case object.MakeDateTimeRange:
		return docstring.MakeDateTimeRange, true
	case object.PayCodeInScheduleMap:
		return docstring.PayCodeInScheduleMap, true
	case object.PayCodeInTimeSheetMap:
		return docstring.PayCodeInTimeSheetMap, true
	case object.Round:
		return docstring.Round, true
	case object.RoundDown:
		return docstring.RoundDown, true
	case object.RoundUp:
		return docstring.RoundUp, true
	case object.RoundToInt:
		return docstring.RoundToInt, true
	case object.SemiMonthlyPeriod:
		return docstring.SemiMonthlyPeriod, true
	case object.Substr:
		return docstring.Substr, true
	case object.ToLowerCase:
		return docstring.ToLowerCase, true
	case object.ToUpperCase:
		return docstring.ToUpperCase, true
	case object.MinSchedule:
		return docstring.MinSchedule, true
	case object.MaxSchedule:
		return docstring.MaxSchedule, true
	case object.AvgSchedule:
		return docstring.AvgSchedule, true
	case object.MinTime:
		return docstring.MinTime, true
	case object.MaxTime:
		return docstring.MaxTime, true
	case object.AvgTime:
		return docstring.AvgTime, true
	case object.SumException:
		return docstring.SumException, true
	case object.MinException:
		return docstring.MinException, true
	case object.MaxException:
		return docstring.MaxException, true
	case object.AvgException:
		return docstring.AvgException, true
	}

	return "", false
}
