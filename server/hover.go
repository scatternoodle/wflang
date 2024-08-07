package server

import (
	"strings"

	"github.com/scatternoodle/wflang/lang/object"
	"github.com/scatternoodle/wflang/lang/token"
	"github.com/scatternoodle/wflang/lsp"
	"github.com/scatternoodle/wflang/server/hovertext"
)

func (srv *Server) hover(pos lsp.Position) lsp.Hover {
	hov := lsp.Hover{
		MarkupContent: lsp.MarkupContent{
			Kind:  lsp.MarkupKindMarkdown,
			Value: "",
		},
	}

	tok, ok := srv.getTokenAtPos(pos)
	if !ok {
		return hov
	}

	if tok.Type == token.T_BUILTIN {
		hov.Value, _ = builtinHoverText(strings.ToLower(tok.Literal))
	}
	return hov
}

func builtinHoverText(name string) (text string, ok bool) {
	switch name {
	case object.Min:
		return hovertext.Min, true
	case object.Max:
		return hovertext.Max, true
	case object.Contains:
		return hovertext.Contains, true
	case object.Sum:
		return hovertext.Sum, true
	case object.Count:
		return hovertext.Count, true
	case object.If:
		return hovertext.If, true
	case object.SumTime:
		return hovertext.SumTime, true
	case object.CountTime:
		return hovertext.CountTime, true
	case object.FindFirstTime:
		return hovertext.FindFirstTime, true
	case object.SumSchedule:
		return hovertext.SumSchedule, true
	case object.CountSchedule:
		return hovertext.CountSchedule, true
	case object.FindFirstSchedule:
		return hovertext.FindFirstSchedule, true
	case object.CountException:
		return hovertext.CountException, true
	case object.FindFirstTorDetail:
		return hovertext.FindFirstTorDetail, true
	case object.FindFirstDayForward:
		return hovertext.FindFirstDayForward, true
	case object.FindFirstDayBackward:
		return hovertext.FindFirstDayBackward, true
	case object.FindFirstDeletedTime:
		return hovertext.FindFirstDeletedTime, true
	case object.LongestConsecutiveRange:
		return hovertext.LongestConsecutiveRange, true
	case object.FirstConsecutiveDay:
		return hovertext.FirstConsecutiveDay, true
	case object.LastConsecutiveDay:
		return hovertext.LastConsecutiveDay, true
	case object.FindNthTime:
		return hovertext.FindNthTime, true
	case object.Accrued:
		return hovertext.Accrued, true
	case object.BalanceAccruedBefore:
		return hovertext.BalanceAccruedBefore, true
	case object.ConvertDttmByTimezone:
		return hovertext.ConvertDttmByTimezone, true
	case object.CountHolidays:
		return hovertext.CountHolidays, true
	case object.GetHoliday:
		return hovertext.GetHoliday, true
	case object.EmployeeAttributeExists:
		return hovertext.Employee_attribute_exists, true
	case object.EmployeeAttribute:
		return hovertext.Employee_attribute, true
	case object.GetAttributeCalcDate:
		return hovertext.GetAttributeCalculationDate, true
	case object.GetBooleanFieldFromTor:
		return hovertext.GetBooleanFieldFromTor, true
	case object.GetDateFieldFromTor:
		return hovertext.GetDateFieldFromTor, true
	case object.GetSelectFieldValueFromTor:
		return hovertext.GetSelectFieldValueFromTor, true
	case object.GetNumberFieldFromTor:
		return hovertext.GetNumberFieldFromTor, true
	case object.GetStringFieldFromTor:
		return hovertext.GetStringValueFromTor, true
	case object.GetSysDateByTimezone:
		return hovertext.GetSysDateByTimezone, true
	case object.LdLookup:
		return hovertext.LDLookup, true
	case object.LdValidate:
		return hovertext.LDValidate, true
	case object.IndexOf:
		return hovertext.IndexOf, true
	case object.GetPayCurrencyCode:
		return hovertext.GetPayCurrencyCode, true
	case object.CallSql:
		return hovertext.CallSQL, true
	case object.LengthOfService:
		return hovertext.LengthOfService, true
	case object.MakeDate:
		return hovertext.MakeDate, true
	case object.MakeDateTime:
		return hovertext.MakeDateTime, true
	case object.MakeDateTimeRange:
		return hovertext.MakeDateTimeRange, true
	case object.PayCodeInScheduleMap:
		return hovertext.PayCodeInScheduleMap, true
	case object.PayCodeInTimeSheetMap:
		return hovertext.PayCodeInTimeSheetMap, true
	case object.Round:
		return hovertext.Round, true
	case object.RoundDown:
		return hovertext.RoundDown, true
	case object.RoundUp:
		return hovertext.RoundUp, true
	case object.RoundToInt:
		return hovertext.RoundToInt, true
	case object.SemiMonthlyPeriod:
		return hovertext.SemiMonthlyPeriod, true
	case object.Substr:
		return hovertext.Substr, true
	case object.ToLowerCase:
		return hovertext.ToLowerCase, true
	case object.ToUpperCase:
		return hovertext.ToUpperCase, true
	case object.MinSchedule:
		return hovertext.MinSchedule, true
	case object.MaxSchedule:
		return hovertext.MaxSchedule, true
	case object.AvgSchedule:
		return hovertext.AvgSchedule, true
	case object.MinTime:
		return hovertext.MinTime, true
	case object.MaxTime:
		return hovertext.MaxTime, true
	case object.AvgTime:
		return hovertext.AvgTime, true
	case object.SumException:
		return hovertext.SumException, true
	case object.MinException:
		return hovertext.MinException, true
	case object.MaxException:
		return hovertext.MaxException, true
	case object.AvgException:
		return hovertext.AvgException, true
	}

	return "", false
}
