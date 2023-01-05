package internal

import (
	"strings"

	"github.com/kaisawind/cobol/asg/model"
)

type Program struct {
	model.Element
	registry         model.Registry
	compilationUnits []model.CompilationUnit
	names            []string
}

func NewProgram() model.Program {
	return &Program{
		Element:  NewElement(nil, nil),
		registry: NewRegistry(),
	}
}

func (e *Program) GetRegistry() model.Registry {
	return e.registry
}

func (e *Program) GetFirstCompilationUnit() model.CompilationUnit {
	if len(e.compilationUnits) == 0 {
		return nil
	}
	return e.compilationUnits[0]
}

func (e *Program) GetCompilationUnit(name string) model.CompilationUnit {
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

func (e *Program) GetCompilationUnits() []model.CompilationUnit {
	return e.compilationUnits
}

func (e *Program) RegistryCompilationUnit(element model.CompilationUnit) {
	name := strings.ToLower(element.Name())
	e.compilationUnits = append(e.compilationUnits, element)
	e.names = append(e.names, name)
}
