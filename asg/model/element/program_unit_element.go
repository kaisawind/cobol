package element

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type ProgramUnitElement interface {
	CompilationUnitElement
	ProgramUnit() ProgramUnit
	GetElement(ctx antlr.Tree) Element
	AddElement(element Element)
}
