package element

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/element"
)

type CompilationUnitElement struct {
	element.Element
	compilationUnit element.CompilationUnit
}

func NewCompilationUnitElement(ctx antlr.ParserRuleContext, compilationUnit element.CompilationUnit) element.CompilationUnitElement {
	return &CompilationUnitElement{
		Element:         NewElement(ctx, compilationUnit.Program()),
		compilationUnit: compilationUnit,
	}
}

func (e *CompilationUnitElement) CompilationUnit() element.CompilationUnit {
	return e.compilationUnit
}
