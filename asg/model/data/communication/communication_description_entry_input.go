package communication

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CommunicationDescriptionEntryInput struct {
	ctx  cobol85.ICommunicationDescriptionEntryFormat1Context
	name string

	symbolicQueueClause *SymbolicQueueClause
}

func NewCommunicationDescriptionEntryInput(ctx cobol85.ICommunicationDescriptionEntryFormat1Context, name string) *CommunicationDescriptionEntryInput {
	return &CommunicationDescriptionEntryInput{
		ctx:  ctx,
		name: name,
	}
}

func (e *CommunicationDescriptionEntryInput) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *CommunicationDescriptionEntryInput) Name() string {
	return e.name
}

func (e *CommunicationDescriptionEntryInput) Type() CDEType {
	return INPUT
}

func (e *CommunicationDescriptionEntryInput) SetSymbolicQueueClause(symbolicQueueClause *SymbolicQueueClause) {
	e.symbolicQueueClause = symbolicQueueClause
}
