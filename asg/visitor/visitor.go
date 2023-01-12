package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type Visitor struct {
	cobol85.BaseCobol85Visitor
	program *model.Program
}

func NewVisitor(program *model.Program) *Visitor {
	return &Visitor{
		program: program,
	}
}

func (v *Visitor) AddCompilationUnit(element *model.CompilationUnit) {
	v.program.AddCompilationUnit(element)
}

func (e *Visitor) AddElement(element model.Element) {
	e.program.Registry().AddElement(element)
}

func (v *Visitor) GetElement(ctx antlr.Tree) model.Element {
	return v.program.Registry().GetElement(ctx)
}
