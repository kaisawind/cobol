package communication

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CommunicationDescriptionEntryInputOutput struct {
	ctx  cobol85.ICommunicationDescriptionEntryFormat2Context
	name string
}

func NewCommunicationDescriptionEntryInputOutput(ctx cobol85.ICommunicationDescriptionEntryFormat2Context, name string) *CommunicationDescriptionEntryInputOutput {
	return &CommunicationDescriptionEntryInputOutput{
		ctx:  ctx,
		name: name,
	}
}

func (e *CommunicationDescriptionEntryInputOutput) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *CommunicationDescriptionEntryInputOutput) Name() string {
	return e.name
}

func (e *CommunicationDescriptionEntryInputOutput) Type() CDEType {
	return INPUT_OUTPUT
}
