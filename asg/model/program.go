package model

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/communication"
)

type ElementType interface {
	*DataDivision | *ProgramUnit | *CompilationUnit | *communication.CommunicationDescriptionEntryInput
}

type program struct {
	registry         *Registry
	compilationUnits []*CompilationUnit
	names            []string
}

func newProgram() *program {
	return &program{
		registry:         NewRegistry(),
		compilationUnits: []*CompilationUnit{},
		names:            []string{},
	}
}

type Program struct {
	*programCall
}

func NewProgram() *Program {
	return &Program{
		programCall: newProgramCall(),
	}
}

func (e *program) Registry() *Registry {
	return e.registry
}

func (e *program) AddElement(element Element) {
	e.Registry().AddElement(element)
}

func (e *program) GetElement(ctx antlr.ParserRuleContext) Element {
	return e.Registry().GetElement(ctx)
}

func (e *program) GetFirstCompilationUnit() *CompilationUnit {
	if len(e.compilationUnits) == 0 {
		return nil
	}
	return e.compilationUnits[0]
}

func (e *program) GetCompilationUnit(name string) *CompilationUnit {
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

func (e *program) GetCompilationUnits() []*CompilationUnit {
	return e.compilationUnits
}

func (e *program) AddCompilationUnit(element *CompilationUnit) {
	name := strings.ToLower(element.Name())
	e.compilationUnits = append(e.compilationUnits, element)
	e.names = append(e.names, name)
}

func (e *program) GetProgramUnit(ctx antlr.Tree) *ProgramUnit {
	return getParent[*ProgramUnit](e, ctx)
}

func getParent[T ElementType](program *program, ctx antlr.Tree) (parent T) {
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
