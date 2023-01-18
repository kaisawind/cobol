package communication

import (
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CommunicationDescriptionEntryInputOutput struct {
	BaseCommunicationDescriptionEntry
}

func NewCommunicationDescriptionEntryInputOutput(ctx cobol85.ICommunicationDescriptionEntryFormat2Context, name string) *CommunicationDescriptionEntryInputOutput {
	return &CommunicationDescriptionEntryInputOutput{
		BaseCommunicationDescriptionEntry: *NewBaseCommunicationDescriptionEntry(ctx, name),
	}
}

func (e *CommunicationDescriptionEntryInputOutput) Type() CDEType {
	return INPUT_OUTPUT
}
