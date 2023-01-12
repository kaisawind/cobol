package model

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ElementType interface {
	*DataDivision | *ProgramUnit | *CompilationUnit
}

type Program struct {
	registry         *Registry
	compilationUnits []*CompilationUnit
	names            []string
}

func GetElement[T ElementType](ctx antlr.Tree, program *Program) (parent T) {
	registry := program.Registry()
	current := ctx
	for {
		if parent != nil || current == nil {
			break
		}
		current = current.GetParent()
		element := registry.GetElement(current)
		if element != nil {
			if parent, ok := element.(T); ok {
				return parent
			}
		}
	}
	return
}

func NewProgram() *Program {
	return &Program{
		registry: NewRegistry(),
	}
}

func (e *Program) Registry() *Registry {
	return e.registry
}

func (e *Program) AddElement(element Element) {
	e.Registry().AddElement(element)
}

func (e *Program) GetElement(ctx antlr.ParserRuleContext) Element {
	return e.Registry().GetElement(ctx)
}

func (e *Program) GetFirstCompilationUnit() *CompilationUnit {
	if len(e.compilationUnits) == 0 {
		return nil
	}
	return e.compilationUnits[0]
}

func (e *Program) GetCompilationUnit(name string) *CompilationUnit {
	name = strings.ToLower(name)
	for k, v := range e.names {
		if v == name {
			if len(e.compilationUnits) > k {
				return e.compilationUnits[k]
			} else {
				return nil
			}
		}
	}
	return nil
}

func (e *Program) GetCompilationUnits() []*CompilationUnit {
	return e.compilationUnits
}

func (e *Program) AddCompilationUnit(element *CompilationUnit) {
	name := strings.ToLower(element.Name())
	e.compilationUnits = append(e.compilationUnits, element)
	e.names = append(e.names, name)
}
