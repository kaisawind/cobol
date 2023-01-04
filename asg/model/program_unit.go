package model

import "github.com/kaisawind/cobol/gen/cobol85"

type ProgramUnit struct {
	*CompilationUnitElement
	ctx *cobol85.ProgramUnitContext
}

func NewProgramUnit(ctx *cobol85.ProgramUnitContext, compilationUnit *CompilationUnit) *ProgramUnit {
	return &ProgramUnit{
		CompilationUnitElement: NewCompilationUnitElement(ctx, compilationUnit),
		ctx:                    ctx,
	}
}
