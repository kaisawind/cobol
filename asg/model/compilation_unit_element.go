package model

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type CompilationUnitElement struct {
	*BaseElement
	compilationUnit *CompilationUnit
}

func NewCompilationUnitElement(ctx antlr.ParserRuleContext, compilationUnit *CompilationUnit) *CompilationUnitElement {
	return &CompilationUnitElement{
		BaseElement:     NewBaseElement(ctx, compilationUnit.Program()),
		compilationUnit: compilationUnit,
	}
}

func (e *CompilationUnitElement) GetCompilationUnit() *CompilationUnit {
	return e.compilationUnit
}
