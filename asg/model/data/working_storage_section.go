package data

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type WorkingStorageSection struct {
	datadescription.BaseDataDescriptionEntry
	ctx cobol85.IWorkingStorageSectionContext
}

func NewWorkingStorageSection(ctx cobol85.IWorkingStorageSectionContext) *WorkingStorageSection {
	return &WorkingStorageSection{
		BaseDataDescriptionEntry: *datadescription.NewBaseDataDescriptionEntry(),
		ctx:                      ctx,
	}
}

func (e *WorkingStorageSection) Context() antlr.ParserRuleContext {
	return e.ctx
}
