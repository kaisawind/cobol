package element

import (
	"strings"

	"github.com/kaisawind/cobol/asg/model/element"
)

type Program struct {
	registry         element.Registry
	compilationUnits []element.CompilationUnit
	names            []string
}

func NewProgram() element.Program {
	return &Program{
		registry: NewRegistry(),
	}
}

func (e *Program) GetRegistry() element.Registry {
	return e.registry
}

func (e *Program) GetFirstCompilationUnit() element.CompilationUnit {
	if len(e.compilationUnits) == 0 {
		return nil
	}
	return e.compilationUnits[0]
}

func (e *Program) GetCompilationUnit(name string) element.CompilationUnit {
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

func (e *Program) GetCompilationUnits() []element.CompilationUnit {
	return e.compilationUnits
}

func (e *Program) RegisterCompilationUnit(element element.CompilationUnit) {
	name := strings.ToLower(element.Name())
	e.compilationUnits = append(e.compilationUnits, element)
	e.names = append(e.names, name)
}
