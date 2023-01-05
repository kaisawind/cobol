package impl

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ProgramUnit struct {
	model.CompilationUnitElement
	ctx *cobol85.ProgramUnitContext
}

func NewProgramUnit(ctx *cobol85.ProgramUnitContext, compilationUnit model.CompilationUnit) model.ProgramUnit {
	return &ProgramUnit{
		CompilationUnitElement: NewCompilationUnitElement(ctx, compilationUnit),
		ctx:                    ctx,
	}
}
