package impl

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
)

type CompilationUnitElement struct {
	model.Element
	compilationUnit model.CompilationUnit
}

func NewCompilationUnitElement(ctx antlr.ParserRuleContext, compilationUnit model.CompilationUnit) model.CompilationUnitElement {
	return &CompilationUnitElement{
		Element:         NewElement(ctx, compilationUnit.Program()),
		compilationUnit: compilationUnit,
	}
}

func (e *CompilationUnitElement) GetCompilationUnit() model.CompilationUnit {
	return e.compilationUnit
}
