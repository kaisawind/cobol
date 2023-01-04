package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
)

type Visitor struct {
	program *model.Program
}

func NewVisitor(program *model.Program) *Visitor {
	return &Visitor{
		program: program,
	}
}

func (v *Visitor) GetCompilationUnit(ctx antlr.Tree) (parent *model.CompilationUnit) {
	registry := v.program.GetRegistry()
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
	registry := v.program.GetRegistry()
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

func (v *Visitor) GetScope(ctx antlr.Tree) (parent *model.Scope) {
	registry := v.program.GetRegistry()
	current := ctx
	for {
		if parent != nil || current == nil {
			break
		}
		current = current.GetParent()
		element := registry.GetElement(current)
		if element != nil {
			if parent, ok := element.(*model.Scope); ok {
				return parent
			}
		}
	}
	return
}

func (v *Visitor) GetElement(ctx antlr.Tree) model.Element {
	return v.program.GetRegistry().GetElement(ctx)
}
