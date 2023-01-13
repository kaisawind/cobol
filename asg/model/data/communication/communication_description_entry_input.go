package communication

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CommunicationDescriptionEntryInput struct {
	ctx cobol85.ICommunicationDescriptionEntryFormat1Context
}

func NewCommunicationDescriptionEntryInput(ctx cobol85.ICommunicationDescriptionEntryFormat1Context) *CommunicationDescriptionEntryInput {
	return &CommunicationDescriptionEntryInput{
		ctx: ctx,
	}
}

func (e *CommunicationDescriptionEntryInput) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *CommunicationDescriptionEntryInput) Type() CDEType {
	return INPUT
}
