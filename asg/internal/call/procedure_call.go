package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/procedure"
)

type ProcedureCall struct {
	call.Call

	paragraph procedure.Paragraph
	callType  call.CallType
}

func NewProcedureCall(
	ctx antlr.ParserRuleContext,
	name string,
	paragraph procedure.Paragraph,
	programUnit element.ProgramUnit,
) call.ProcedureCall {
	return &ProcedureCall{
		Call: NewCall(ctx, name, programUnit),
	}
}

func (e *ProcedureCall) Paragraph() procedure.Paragraph {
	return e.paragraph
}

func (e *ProcedureCall) Type() call.CallType {
	return e.callType
}
