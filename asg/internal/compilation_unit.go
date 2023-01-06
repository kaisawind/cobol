package internal

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CompilationUnit struct {
	model.Element
	ctx    *cobol85.CompilationUnitContext
	name   string
	tokens *antlr.CommonTokenStream
}

func NewCompilationUnit(name string, program *Program, tokens *antlr.CommonTokenStream, ctx *cobol85.CompilationUnitContext) model.CompilationUnit {
	compilationUnit := &CompilationUnit{
		Element: NewElement(ctx, program),
		ctx:     ctx,
		name:    name,
		tokens:  tokens,
	}
	return compilationUnit
}

func (e *CompilationUnit) RegisterElement(element model.Element) {
	if element == nil || element.Context() == nil {
		return
	}
	e.Program().GetRegistry().AddElement(element)
}

func (e *CompilationUnit) Name() string {
	return e.name
}

func (e *CompilationUnit) Context() antlr.ParserRuleContext {
	return e.ctx
}
