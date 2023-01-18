package communication

import (
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CommunicationDescriptionEntryInput struct {
	BaseCommunicationDescriptionEntry

	symbolicQueueClause *SymbolicQueueClause
}

func NewCommunicationDescriptionEntryInput(ctx cobol85.ICommunicationDescriptionEntryFormat1Context, name string) *CommunicationDescriptionEntryInput {
	return &CommunicationDescriptionEntryInput{
		BaseCommunicationDescriptionEntry: *NewBaseCommunicationDescriptionEntry(ctx, name),
	}
}

func (e *CommunicationDescriptionEntryInput) Type() CDEType {
	return INPUT
}

func (e *CommunicationDescriptionEntryInput) SetSymbolicQueueClause(symbolicQueueClause *SymbolicQueueClause) {
	e.symbolicQueueClause = symbolicQueueClause
}
