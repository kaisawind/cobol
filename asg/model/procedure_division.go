package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ProcedureDivision struct {
	ctx cobol85.IProcedureDivisionContext
}

func NewProcedureDivision(ctx cobol85.IProcedureDivisionContext) *ProcedureDivision {
	return &ProcedureDivision{
		ctx: ctx,
	}
}

func (e *ProcedureDivision) Context() antlr.ParserRuleContext {
	return e.ctx
}
