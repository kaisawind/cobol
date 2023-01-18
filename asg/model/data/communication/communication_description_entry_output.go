package communication

import (
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CommunicationDescriptionEntryOutput struct {
	BaseCommunicationDescriptionEntry
}

func NewCommunicationDescriptionEntryOutput(ctx cobol85.ICommunicationDescriptionEntryFormat3Context, name string) *CommunicationDescriptionEntryOutput {
	return &CommunicationDescriptionEntryOutput{
		BaseCommunicationDescriptionEntry: *NewBaseCommunicationDescriptionEntry(ctx, name),
	}
}

func (e *CommunicationDescriptionEntryOutput) Type() CDEType {
	return OUTPUT
}
