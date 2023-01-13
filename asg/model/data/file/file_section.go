package file

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type FileSection struct {
	ctx cobol85.IFileSectionContext
}

func NewFileSection(ctx cobol85.IFileSectionContext) *FileSection {
	return &FileSection{
		ctx: ctx,
	}
}

func (e *FileSection) Context() antlr.ParserRuleContext {
	return e.ctx
}
