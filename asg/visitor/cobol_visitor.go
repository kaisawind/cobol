package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CobolVisitor struct {
	cobol85.BaseCobol85Visitor
	program model.Program
}

func NewCobolVisitor(program model.Program) *CobolVisitor {
	return &CobolVisitor{
		program: program,
	}
}

func (v *CobolVisitor) GetCompilationUnit(ctx antlr.Tree) (parent model.CompilationUnit) {
	registry := v.program.GetRegistry()
	current := ctx
	for {
		if parent != nil || current == nil {
			break
		}
		current = current.GetParent()
		element := registry.GetElement(current)
		if element != nil {
			if parent, ok := element.(model.CompilationUnit); ok {
				return parent
			}
		}
	}
	return
}

func (v *CobolVisitor) GetProgramUnit(ctx antlr.Tree) (parent model.ProgramUnit) {
	registry := v.program.GetRegistry()
	current := ctx
	for {
		if parent != nil || current == nil {
			break
		}
		current = current.GetParent()
		element := registry.GetElement(current)
		if element != nil {
			if parent, ok := element.(model.ProgramUnit); ok {
				return parent
			}
		}
	}
	return
}

func (v *CobolVisitor) GetScope(ctx antlr.Tree) (parent model.Scope) {
	registry := v.program.GetRegistry()
	current := ctx
	for {
		if parent != nil || current == nil {
			break
		}
		current = current.GetParent()
		element := registry.GetElement(current)
		if element != nil {
			if parent, ok := element.(model.Scope); ok {
				return parent
			}
		}
	}
	return
}

func (v *CobolVisitor) GetElement(ctx antlr.Tree) model.Element {
	return v.program.GetRegistry().GetElement(ctx)
}
