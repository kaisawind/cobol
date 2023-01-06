package model

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type ProgramUnitElement interface {
	CompilationUnitElement
	GetProgramUnit() ProgramUnit
	GetElement(ctx antlr.Tree) Element
	AddElement(element Element)
}
