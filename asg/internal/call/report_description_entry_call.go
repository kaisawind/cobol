package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/report"
	"github.com/kaisawind/cobol/asg/model/element"
)

type ReportDescriptionEntryCall struct {
	call.Call

	reportDescriptionEntry report.ReportDescriptionEntry
	callType               call.CallType
}

func NewReportDescriptionEntryCall(
	ctx antlr.ParserRuleContext,
	name string,
	reportDescriptionEntry report.ReportDescriptionEntry,
	programUnit element.ProgramUnit,
) call.ReportDescriptionEntryCall {
	return &ReportDescriptionEntryCall{
		Call:                   NewCall(ctx, name, programUnit),
		reportDescriptionEntry: reportDescriptionEntry,
		callType:               call.REPORT_DESCRIPTION_ENTRY_CALL,
	}
}

func (e *ReportDescriptionEntryCall) ReportDescriptionEntry() report.ReportDescriptionEntry {
	return e.reportDescriptionEntry
}

func (e *ReportDescriptionEntryCall) Type() call.CallType {
	return e.callType
}
