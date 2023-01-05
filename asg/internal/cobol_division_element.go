package internal

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
)

type CobolDivisionElement struct {
	model.ProgramUnitElement
}

func NewCobolDivisionElement(ctx antlr.ParserRuleContext, programUnit model.ProgramUnit) model.CobolDivisionElement {
	return &CobolDivisionElement{
		ProgramUnitElement: NewProgramUnitElement(ctx, programUnit),
	}
}
