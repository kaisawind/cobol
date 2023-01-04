package model

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type ProgramUnitElement struct {
	*CompilationUnitElement
	programUnit *ProgramUnit
}

func NewProgramUnitElement(ctx antlr.ParserRuleContext, programUnit *ProgramUnit) *ProgramUnitElement {
	return &ProgramUnitElement{
		CompilationUnitElement: NewCompilationUnitElement(ctx, programUnit.GetCompilationUnit()),
		programUnit:            programUnit,
	}
}

func (e *ProgramUnitElement) GetProgramUnit() *ProgramUnit {
	return e.programUnit
}
