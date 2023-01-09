package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/report"
	"github.com/kaisawind/cobol/asg/model/element"
)

type ReportDescriptionCall struct {
	call.Call

	reportDescription report.ReportDescription
	callType          call.CallType
}

func NewReportCall(
	ctx antlr.ParserRuleContext,
	name string,
	reportDescription report.ReportDescription,
	programUnit element.ProgramUnit,
) call.ReportDescriptionCall {
	return &ReportDescriptionCall{
		Call:              NewCall(ctx, name, programUnit),
		reportDescription: reportDescription,
		callType:          call.REPORT_DESCRIPTION_CALL,
	}
}

func (e *ReportDescriptionCall) ReportDescription() report.ReportDescription {
	return e.reportDescription
}

func (e *ReportDescriptionCall) Type() call.CallType {
	return e.callType
}
