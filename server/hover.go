package server

import (
	"log/slog"
	"strings"

	"github.com/scatternoodle/wflang/lang/builtins"
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
		slog.Debug("No token found")
		return hov
	}
	slog.Debug("token found", "token", tok)

	if tok.Type == token.T_BUILTIN {
		hov.Value, _ = builtinHoverText(strings.ToLower(tok.Literal))
	}

	return hov
}

func builtinHoverText(name string) (text string, ok bool) {
	switch name {
	case builtins.Min:
		return hovertext.Min, true
	case builtins.Max:
		return hovertext.Max, true
	case builtins.Contains:
		return hovertext.Contains, true
	case builtins.Sum:
		return hovertext.Sum, true
	case builtins.Count:
		return hovertext.Count, true
	case builtins.If:
		return hovertext.If, true
	case builtins.SumTime:
		return hovertext.SumTime, true
	case builtins.CountTime:
		return hovertext.CountTime, true
	case builtins.FindFirstTime:
		return hovertext.FindFirstTime, true
	case builtins.SumSchedule:
		return hovertext.SumSchedule, true
	case builtins.CountSchedule:
		return hovertext.CountSchedule, true
	case builtins.FindFirstSchedule:
		return hovertext.FindFirstSchedule, true
	case builtins.CountException:
		return hovertext.CountException, true
	case builtins.FindFirstTorDetail:
		return hovertext.FindFirstTorDetail, true
	case builtins.FindFirstDayForward:
		return hovertext.FindFirstDayForward, true
	case builtins.FindFirstDayBackward:
		return hovertext.FindFirstDayBackward, true
	case builtins.FindFirstDeletedTime:
		return hovertext.FindFirstDeletedTime, true
	case builtins.LongestConsecutiveRange:
		return hovertext.LongestConsecutiveRange, true
	case builtins.FirstConsecutiveDay:
		return hovertext.FirstConsecutiveDay, true
	case builtins.LastConsecutiveDay:
		return hovertext.LastConsecutiveDay, true
	case builtins.FindNthTime:
		return hovertext.FindNthTime, true
	case builtins.Accrued:
		return hovertext.Accrued, true
	case builtins.BalanceAccruedBefore:
		return hovertext.BalanceAccruedBefore, true
	case builtins.ConvertDttmByTimezone:
		return hovertext.ConvertDttmByTimezone, true
	case builtins.CountHolidays:
		return hovertext.CountHolidays, true
	case builtins.GetHoliday:
		return hovertext.GetHoliday, true
	case builtins.EmployeeAttributeExists:
		return hovertext.Employee_attribute_exists, true
	case builtins.EmployeeAttribute:
		return hovertext.Employee_attribute, true
	case builtins.GetAttributeCalcDate:
		return hovertext.GetAttributeCalculationDate, true
	case builtins.GetBooleanFieldFromTor:
		return hovertext.GetBooleanFieldFromTor, true
	case builtins.GetDateFieldFromTor:
		return hovertext.GetDateFieldFromTor, true
	case builtins.GetSelectFieldValueFromTor:
		return hovertext.GetSelectFieldValueFromTor, true
	case builtins.GetNumberFieldFromTor:
		return hovertext.GetNumberFieldFromTor, true
	case builtins.GetStringFieldFromTor:
		return hovertext.GetStringValueFromTor, true
	case builtins.GetSysDateByTimezone:
		return hovertext.GetSysDateByTimezone, true
	case builtins.LdLookup:
		return hovertext.LDLookup, true
	case builtins.LdValidate:
		return hovertext.LDValidate, true
	case builtins.IndexOf:
		return hovertext.IndexOf, true
	case builtins.GetPayCurrencyCode:
		return hovertext.GetPayCurrencyCode, true
	case builtins.CallSql:
		return hovertext.CallSQL, true
	case builtins.LengthOfService:
		return hovertext.LengthOfService, true
	case builtins.MakeDate:
		return hovertext.MakeDate, true
	case builtins.MakeDateTime:
		return hovertext.MakeDateTime, true
	case builtins.MakeDateTimeRange:
		return hovertext.MakeDateTimeRange, true
	}

	return "", false
}
