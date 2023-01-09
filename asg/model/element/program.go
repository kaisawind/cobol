package element

type Program interface {
	GetRegistry() Registry
	GetFirstCompilationUnit() CompilationUnit
	GetCompilationUnit(name string) CompilationUnit
	GetCompilationUnits() []CompilationUnit
	RegisterCompilationUnit(element CompilationUnit)
}
