package communication

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CommunicationSection struct {
	datadescription.BaseDataDescriptionEntry
	ctx                                  cobol85.ICommunicationSectionContext
	communicationDescriptionEntries      []CommunicationDescriptionEntry
	communicationDescriptionEntriesTable map[string]CommunicationDescriptionEntry
}

func NewCommunicationSection(ctx cobol85.ICommunicationSectionContext) *CommunicationSection {
	return &CommunicationSection{
		BaseDataDescriptionEntry:             *datadescription.NewBaseDataDescriptionEntry(),
		ctx:                                  ctx,
		communicationDescriptionEntries:      []CommunicationDescriptionEntry{},
		communicationDescriptionEntriesTable: map[string]CommunicationDescriptionEntry{},
	}
}

func (e *CommunicationSection) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *CommunicationSection) AddCommunicationDescriptionEntry(entry CommunicationDescriptionEntry) {
	e.communicationDescriptionEntries = append(e.communicationDescriptionEntries, entry)
	e.communicationDescriptionEntriesTable[entry.Name()] = entry
}

func (e *CommunicationSection) GetCommunicationDescriptionEntries() []CommunicationDescriptionEntry {
	return e.communicationDescriptionEntries
}

func (e *CommunicationSection) GetCommunicationDescriptionEntryByName(name string) CommunicationDescriptionEntry {
	for k, v := range e.communicationDescriptionEntriesTable {
		if k == name {
			return v
		}
	}
	return nil
}
