package data

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type WorkingStorageSection struct {
	ctx     cobol85.IWorkingStorageSectionContext
	entries []datadescription.DataDescriptionEntry
}

func NewWorkingStorageSection(ctx cobol85.IWorkingStorageSectionContext) *WorkingStorageSection {
	return &WorkingStorageSection{
		ctx: ctx,
	}
}

func (e *WorkingStorageSection) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *WorkingStorageSection) AddCommunicationDescriptionEntry(entry datadescription.DataDescriptionEntry) {
	e.entries = append(e.entries, entry)
}

func (e *WorkingStorageSection) GetCommunicationDescriptionEntries() []datadescription.DataDescriptionEntry {
	return e.entries
}
