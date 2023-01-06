package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/communication"
)

type CommunicationDescriptionEntryCall struct {
	call.Call
	communicationDescriptionEntry communication.CommunicationDescriptionEntry
	callType                      call.CallType
}

func NewCommunicationDescriptionEntryCall(
	ctx antlr.ParserRuleContext,
	name string,
	communicationDescriptionEntry communication.CommunicationDescriptionEntry,
	programUnit model.ProgramUnit,
) call.CommunicationDescriptionEntryCall {
	return &CommunicationDescriptionEntryCall{
		Call:                          NewCall(ctx, name, programUnit),
		communicationDescriptionEntry: communicationDescriptionEntry,
		callType:                      call.COMMUNICATION_DESCRIPTION_ENTRY_CALL,
	}
}

func (e *CommunicationDescriptionEntryCall) Type() call.CallType {
	return e.callType
}

func (e *CommunicationDescriptionEntryCall) CommunicationDescriptionEntry() communication.CommunicationDescriptionEntry {
	return e.communicationDescriptionEntry
}
