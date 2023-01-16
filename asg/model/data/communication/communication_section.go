package communication

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CommunicationSection struct {
	datadescription.BaseDataDescriptionEntry
	ctx                             cobol85.ICommunicationSectionContext
	communicationDescriptionEntries []CommunicationDescriptionEntry
}

func NewCommunicationSection(ctx cobol85.ICommunicationSectionContext) *CommunicationSection {
	return &CommunicationSection{
		BaseDataDescriptionEntry:        *datadescription.NewBaseDataDescriptionEntry(),
		ctx:                             ctx,
		communicationDescriptionEntries: []CommunicationDescriptionEntry{},
	}
}

func (e *CommunicationSection) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *CommunicationSection) AddCommunicationDescriptionEntry(entry CommunicationDescriptionEntry) {
	e.communicationDescriptionEntries = append(e.communicationDescriptionEntries, entry)
}

func (e *CommunicationSection) GetCommunicationDescriptionEntries() []CommunicationDescriptionEntry {
	return e.communicationDescriptionEntries
}
