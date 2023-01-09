package element

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/data"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/environment"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ProgramUnit struct {
	element.CompilationUnitElement
	ctx                 *cobol85.ProgramUnitContext
	dataDivision        data.DataDivision
	environmentDivision environment.EnvironmentDivision
}

func NewProgramUnit(ctx *cobol85.ProgramUnitContext, compilationUnit element.CompilationUnit) element.ProgramUnit {
	return &ProgramUnit{
		CompilationUnitElement: NewCompilationUnitElement(ctx, compilationUnit),
		ctx:                    ctx,
	}
}

func (e *ProgramUnit) GetElement(ctx antlr.Tree) element.Element {
	return e.Program().GetRegistry().GetElement(ctx)
}

func (e *ProgramUnit) AddElement(element element.Element) {
	e.Program().GetRegistry().AddElement(element)
}

func (e *ProgramUnit) SetDataDivision(ctx *cobol85.DataDivisionContext) (ret data.DataDivision) {
	element := e.GetElement(ctx)
	if element != nil {
		ret = element.(data.DataDivision)
	} else {
		if instances.NewDataDivision != nil {
			ret = instances.NewDataDivision(ctx, e)
			e.AddElement(ret)
		}
	}
	return
}

func (e *ProgramUnit) GetDataDivision() data.DataDivision {
	return e.dataDivision
}

func (e *ProgramUnit) SetEnvironmentDivision(ctx *cobol85.EnvironmentDivisionContext) (ret environment.EnvironmentDivision) {
	element := e.GetElement(ctx)
	if element != nil {
		ret = element.(environment.EnvironmentDivision)
	} else {
		if instances.NewEnvironmentDivision != nil {
			ret = instances.NewEnvironmentDivision(ctx, e)
			e.AddElement(ret)
		}
	}
	return
}

func (e *ProgramUnit) GetEnvironmentDivision() environment.EnvironmentDivision {
	return e.environmentDivision
}
