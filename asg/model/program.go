package model

import (
	"strings"
)

type Program struct {
	*BaseElement
	registry         *Registry
	compilationUnits []*CompilationUnit
	names            []string
}

func NewProgram() *Program {
	return &Program{
		BaseElement: NewBaseElement(nil, nil),
		registry:    NewRegistry(),
	}
}

func (e *Program) GetRegistry() *Registry {
	return e.registry
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

func (e *Program) RegistryCompilationUnit(element *CompilationUnit) {
	name := strings.ToLower(element.Name())
	e.compilationUnits = append(e.compilationUnits, element)
	e.names = append(e.names, name)
}
