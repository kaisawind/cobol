package element

type CompilationUnitElement interface {
	Element
	CompilationUnit() CompilationUnit
}
