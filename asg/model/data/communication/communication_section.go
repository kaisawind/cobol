package communication

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CommunicationSection struct {
	ctx     cobol85.ICommunicationSectionContext
	entries []CommunicationDescriptionEntry
}

func NewCommunicationSection(ctx cobol85.ICommunicationSectionContext) *CommunicationSection {
	return &CommunicationSection{
		ctx:     ctx,
		entries: []CommunicationDescriptionEntry{},
	}
}

func (e *CommunicationSection) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *CommunicationSection) AddCommunicationDescriptionEntry(entry CommunicationDescriptionEntry) {
	e.entries = append(e.entries, entry)
}
