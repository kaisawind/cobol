package linkage

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type LinkageSection struct {
	datadescription.BaseDataDescriptionEntry
	ctx cobol85.ILinkageSectionContext
}

func NewLinkageSection(ctx cobol85.ILinkageSectionContext) *LinkageSection {
	return &LinkageSection{
		BaseDataDescriptionEntry: *datadescription.NewBaseDataDescriptionEntry(),
		ctx:                      ctx,
	}
}

func (e *LinkageSection) Context() antlr.ParserRuleContext {
	return e.ctx
}
