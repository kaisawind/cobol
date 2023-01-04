package model

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type DivisionElement struct {
	*ProgramUnitElement
}

func NewDivisionElement(ctx antlr.ParserRuleContext, programUnit *ProgramUnit) *DivisionElement {
	return &DivisionElement{
		ProgramUnitElement: NewProgramUnitElement(ctx, programUnit),
	}
}
