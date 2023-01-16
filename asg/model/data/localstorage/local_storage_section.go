package localstorage

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type LocalStorageSection struct {
	datadescription.BaseDataDescriptionEntry
	ctx cobol85.ILocalStorageSectionContext
}

func NewLocalStorageSection(ctx cobol85.ILocalStorageSectionContext) *LocalStorageSection {
	return &LocalStorageSection{
		BaseDataDescriptionEntry: *datadescription.NewBaseDataDescriptionEntry(),
		ctx:                      ctx,
	}
}

func (e *LocalStorageSection) Context() antlr.ParserRuleContext {
	return e.ctx
}
