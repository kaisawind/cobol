package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type EnvironmentDivision struct {
	ctx cobol85.IEnvironmentDivisionContext
}

func NewEnvironmentDivision(ctx cobol85.IEnvironmentDivisionContext) *EnvironmentDivision {
	return &EnvironmentDivision{
		ctx: ctx,
	}
}

func (e *EnvironmentDivision) Context() antlr.ParserRuleContext {
	return e.ctx
}
