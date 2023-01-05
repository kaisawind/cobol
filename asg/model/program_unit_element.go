package model

type ProgramUnitElement interface {
	CompilationUnitElement
	GetProgramUnit() ProgramUnit
}
