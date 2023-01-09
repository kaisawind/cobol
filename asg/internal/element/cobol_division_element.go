package element

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
)

type CobolDivisionElement struct {
	element.ProgramUnitElement
}

func init() {
	instances.RegisterNewCobolDivisionElementFunc(NewCobolDivisionElement)
}

func NewCobolDivisionElement(ctx antlr.ParserRuleContext, programUnit element.ProgramUnit) element.CobolDivisionElement {
	return &CobolDivisionElement{
		ProgramUnitElement: NewProgramUnitElement(ctx, programUnit),
	}
}
