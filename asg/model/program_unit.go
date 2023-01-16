package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ProgramUnit struct {
	ctx                    cobol85.IProgramUnitContext
	identificationDivision *IdentificationDivision
	dataDivision           *DataDivision
	environmentDivision    *EnvironmentDivision
	procedureDivision      *ProcedureDivision
}

func NewProgramUnit(ctx cobol85.IProgramUnitContext) *ProgramUnit {
	return &ProgramUnit{
		ctx: ctx,
	}
}

func (e *ProgramUnit) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *ProgramUnit) SetDataDivision(dataDivision *DataDivision) {
	e.dataDivision = dataDivision
}

func (e *ProgramUnit) GetDataDivision() *DataDivision {
	return e.dataDivision
}

func (e *ProgramUnit) SetEnvironmentDivision(environmentDivision *EnvironmentDivision) {
	e.environmentDivision = environmentDivision
}

func (e *ProgramUnit) GetEnvironmentDivision() *EnvironmentDivision {
	return e.environmentDivision
}

func (e *ProgramUnit) SetIdentificationDivision(identificationDivision *IdentificationDivision) {
	e.identificationDivision = identificationDivision
}

func (e *ProgramUnit) GetIdentificationDivision() *IdentificationDivision {
	return e.identificationDivision
}

func (e *ProgramUnit) SetProcedureDivision(procedureDivision *ProcedureDivision) {
	e.procedureDivision = procedureDivision
}

func (e *ProgramUnit) GetProcedureDivision() *ProcedureDivision {
	return e.procedureDivision
}
