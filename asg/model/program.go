package model

type Program interface {
	Element
	GetRegistry() Registry
	GetFirstCompilationUnit() CompilationUnit
	GetCompilationUnit(name string) CompilationUnit
	GetCompilationUnits() []CompilationUnit
	RegisterCompilationUnit(element CompilationUnit)
}
