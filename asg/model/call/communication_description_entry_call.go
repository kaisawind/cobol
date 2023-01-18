package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/communication"
)

type CommunicationDescriptionEntryCall struct {
	ctx                           antlr.ParserRuleContext
	name                          string
	communicationDescriptionEntry communication.CommunicationDescriptionEntry
}

func NewCommunicationDescriptionEntryCall(
	ctx antlr.ParserRuleContext,
	name string,
	communicationDescriptionEntry communication.CommunicationDescriptionEntry,
) *CommunicationDescriptionEntryCall {
	return &CommunicationDescriptionEntryCall{
		ctx:                           ctx,
		name:                          name,
		communicationDescriptionEntry: communicationDescriptionEntry,
	}
}

func (e *CommunicationDescriptionEntryCall) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *CommunicationDescriptionEntryCall) Name() string {
	return e.name
}

func (e *CommunicationDescriptionEntryCall) Type() CallType {
	return COMMUNICATION_DESCRIPTION_ENTRY_CALL
}
