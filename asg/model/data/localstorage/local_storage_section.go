package localstorage

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type LocalStorageSection struct {
	ctx cobol85.ILocalStorageSectionContext
}

func NewLocalStorageSection(ctx cobol85.ILocalStorageSectionContext) *LocalStorageSection {
	return &LocalStorageSection{
		ctx: ctx,
	}
}

func (e *LocalStorageSection) Context() antlr.ParserRuleContext {
	return e.ctx
}
