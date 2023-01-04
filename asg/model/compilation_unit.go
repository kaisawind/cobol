package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CompilationUnit struct {
	*BaseElement
	ctx    *cobol85.CompilationUnitContext
	name   string
	tokens *antlr.CommonTokenStream
}

func NewCompilationUnit(name string, program *Program, tokens *antlr.CommonTokenStream, ctx *cobol85.CompilationUnitContext) *CompilationUnit {
	compilationUnit := &CompilationUnit{
		BaseElement: NewBaseElement(ctx, program),
		ctx:         ctx,
		name:        name,
		tokens:      tokens,
	}
	return compilationUnit
}

func (e *CompilationUnit) RegistryElement(element Element) {
	if element == nil || element.Context() == nil {
		return
	}
	e.program.GetRegistry().AddElement(element)
}

func (e *CompilationUnit) Name() string {
	return e.name
}

func (e *CompilationUnit) Context() antlr.ParserRuleContext {
	return e.ctx
}
