package model

type CompilationUnitElement interface {
	Element
	CompilationUnit() CompilationUnit
}
