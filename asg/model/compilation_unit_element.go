package model

type CompilationUnitElement interface {
	Element
	GetCompilationUnit() CompilationUnit
}
