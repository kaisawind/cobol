package element

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
)

type CobolDivision struct {
	element.ProgramUnitElement
}

func init() {
	instances.RegisterNewCobolDivisionFunc(NewCobolDivision)
}

func NewCobolDivision(ctx antlr.ParserRuleContext, programUnit element.ProgramUnit) element.CobolDivision {
	return &CobolDivision{
		ProgramUnitElement: NewProgramUnitElement(ctx, programUnit),
	}
}
