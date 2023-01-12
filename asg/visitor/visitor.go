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

func (v *Visitor) GetCompilationUnit(ctx antlr.Tree) (parent *model.CompilationUnit) {
	registry := v.program.Registry()
	current := ctx
	for {
		if parent != nil || current == nil {
			break
		}
		current = current.GetParent()
		element := registry.GetElement(current)
		if element != nil {
			if parent, ok := element.(*model.CompilationUnit); ok {
				return parent
			}
		}
	}
	return
}

func (v *Visitor) GetProgramUnit(ctx antlr.Tree) (parent *model.ProgramUnit) {
	registry := v.program.Registry()
	current := ctx
	for {
		if parent != nil || current == nil {
			break
		}
		current = current.GetParent()
		element := registry.GetElement(current)
		if element != nil {
			if parent, ok := element.(*model.ProgramUnit); ok {
				return parent
			}
		}
	}
	return
}

func (e *Visitor) AddElement(element model.Element) {
	e.program.Registry().AddElement(element)
}

func (v *Visitor) GetElement(ctx antlr.Tree) model.Element {
	return v.program.Registry().GetElement(ctx)
}
