package communication

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CommunicationDescriptionEntryOutput struct {
	ctx cobol85.ICommunicationDescriptionEntryFormat3Context
}

func NewCommunicationDescriptionEntryOutput(ctx cobol85.ICommunicationDescriptionEntryFormat3Context) *CommunicationDescriptionEntryOutput {
	return &CommunicationDescriptionEntryOutput{
		ctx: ctx,
	}
}

func (e *CommunicationDescriptionEntryOutput) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *CommunicationDescriptionEntryOutput) Type() CDEType {
	return OUTPUT
}
