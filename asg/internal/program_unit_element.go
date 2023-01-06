package internal

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
)

type ProgramUnitElement struct {
	model.CompilationUnitElement
	programUnit model.ProgramUnit
}

func NewProgramUnitElement(ctx antlr.ParserRuleContext, programUnit model.ProgramUnit) model.ProgramUnitElement {
	return &ProgramUnitElement{
		CompilationUnitElement: NewCompilationUnitElement(ctx, programUnit.GetCompilationUnit()),
		programUnit:            programUnit,
	}
}

func (e *ProgramUnitElement) GetProgramUnit() model.ProgramUnit {
	return e.programUnit
}

func (e *ProgramUnitElement) GetElement(ctx antlr.Tree) model.Element {
	return e.Program().GetRegistry().GetElement(ctx)
}

func (e *ProgramUnitElement) AddElement(element model.Element) {
	e.Program().GetRegistry().AddElement(element)
}
