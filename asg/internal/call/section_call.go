package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/procedure"
)

type SectionCall struct {
	call.Call

	section  procedure.Section
	callType call.CallType
}

func NewSectionCall(
	ctx antlr.ParserRuleContext,
	name string,
	section procedure.Section,
	programUnit model.ProgramUnit,
) call.SectionCall {
	return &SectionCall{
		Call:     NewCall(ctx, name, programUnit),
		section:  section,
		callType: call.SECTION_CALL,
	}
}

func (e *SectionCall) Section() procedure.Section {
	return e.section
}

func (e *SectionCall) Type() call.CallType {
	return e.callType
}
