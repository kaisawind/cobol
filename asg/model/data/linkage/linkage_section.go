package linkage

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type LinkageSection struct {
	ctx cobol85.ILinkageSectionContext
}

func NewLinkageSection(ctx cobol85.ILinkageSectionContext) *LinkageSection {
	return &LinkageSection{
		ctx: ctx,
	}
}

func (e *LinkageSection) Context() antlr.ParserRuleContext {
	return e.ctx
}
