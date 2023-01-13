package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type IdentificationDivision struct {
	ctx cobol85.IIdentificationDivisionContext
}

func NewIdentificationDivision(ctx cobol85.IIdentificationDivisionContext) *IdentificationDivision {
	return &IdentificationDivision{
		ctx: ctx,
	}
}

func (e *IdentificationDivision) Context() antlr.ParserRuleContext {
	return e.ctx
}
