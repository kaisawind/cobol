package file

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type FileSection struct {
	datadescription.BaseDataDescriptionEntry
	ctx cobol85.IFileSectionContext
}

func NewFileSection(ctx cobol85.IFileSectionContext) *FileSection {
	return &FileSection{
		BaseDataDescriptionEntry: *datadescription.NewBaseDataDescriptionEntry(),
		ctx:                      ctx,
	}
}

func (e *FileSection) Context() antlr.ParserRuleContext {
	return e.ctx
}
