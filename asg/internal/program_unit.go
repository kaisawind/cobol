package internal

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	internalData "github.com/kaisawind/cobol/asg/internal/data"
	internalEnv "github.com/kaisawind/cobol/asg/internal/environment"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/data"
	"github.com/kaisawind/cobol/asg/model/environment"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ProgramUnit struct {
	model.CompilationUnitElement
	ctx                 *cobol85.ProgramUnitContext
	dataDivision        data.DataDivision
	environmentDivision environment.EnvironmentDivision
}

func NewProgramUnit(ctx *cobol85.ProgramUnitContext, compilationUnit model.CompilationUnit) model.ProgramUnit {
	return &ProgramUnit{
		CompilationUnitElement: NewCompilationUnitElement(ctx, compilationUnit),
		ctx:                    ctx,
	}
}

func (e *ProgramUnit) GetElement(ctx antlr.Tree) model.Element {
	return e.Program().GetRegistry().GetElement(ctx)
}

func (e *ProgramUnit) AddElement(element model.Element) {
	e.Program().GetRegistry().AddElement(element)
}

func (e *ProgramUnit) SetDataDivision(ctx *cobol85.DataDivisionContext) (ret data.DataDivision) {
	element := e.GetElement(ctx)
	if element != nil {
		ret = element.(data.DataDivision)
	} else {
		ret = internalData.NewDataDivision(ctx, NewCobolDivision(ctx, e))
		e.AddElement(ret)
	}
	return
}

func (e *ProgramUnit) GetDataDivision() data.DataDivision {
	return e.dataDivision
}

func (e *ProgramUnit) SetEnvironmentDivision(ctx *cobol85.EnvironmentDivisionContext) (ret environment.EnvironmentDivision) {
	element := e.GetElement(ctx)
	if element != nil {
		ret = element.(data.DataDivision)
	} else {
		ret = internalEnv.NewEnvironmentDivision(ctx, NewCobolDivision(ctx, e))
		e.AddElement(ret)
	}
	return
}

func (e *ProgramUnit) GetEnvironmentDivision() environment.EnvironmentDivision {
	return e.environmentDivision
}
