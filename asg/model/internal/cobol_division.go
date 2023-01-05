package impl

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
)

type CobolDivision struct {
	model.ProgramUnitElement
}

func NewCobolDivision(ctx antlr.ParserRuleContext, programUnit model.ProgramUnit) model.CobolDivision {
	return &CobolDivision{
		ProgramUnitElement: NewProgramUnitElement(ctx, programUnit),
	}
}
