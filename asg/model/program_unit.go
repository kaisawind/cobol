package model

import "github.com/kaisawind/cobol/gen/cobol85"

type ProgramUnit struct {
	Element
	dataDivision           *DataDivision
	environmentDivision    *EnvironmentDivision
	identificationDivision *IdentificationDivision
	procedureDivision      *ProcedureDivision
}

func NewProgramUnit(ctx *cobol85.ProgramUnitContext) *ProgramUnit {
	return &ProgramUnit{
		Element: NewBaseElement(ctx),
	}
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
