package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/screen"
	"github.com/kaisawind/cobol/asg/model/element"
)

type ScreenDescriptionEntryCall struct {
	call.Call

	screenDescriptionEntry screen.ScreenDescriptionEntry
	callType               call.CallType
}

func NewScreenDescriptionEntryCall(
	ctx antlr.ParserRuleContext,
	name string,
	screenDescriptionEntry screen.ScreenDescriptionEntry,
	programUnit element.ProgramUnit,
) call.ScreenDescriptionEntryCall {
	return &ScreenDescriptionEntryCall{
		Call:                   NewCall(ctx, name, programUnit),
		screenDescriptionEntry: screenDescriptionEntry,
		callType:               call.SCREEN_DESCRIPTION_ENTRY_CALL,
	}
}

func (e *ScreenDescriptionEntryCall) ScreenDescriptionEntry() screen.ScreenDescriptionEntry {
	return e.screenDescriptionEntry
}

func (e *ScreenDescriptionEntryCall) Type() call.CallType {
	return e.callType
}
