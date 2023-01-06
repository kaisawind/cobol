package internal

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/data"
	"github.com/kaisawind/cobol/asg/model/environment"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type NewDataDivisionFunc func(ctx *cobol85.DataDivisionContext, programUnit model.ProgramUnit) data.DataDivision

type NewEnvironmentDivisionFunc func(ctx *cobol85.EnvironmentDivisionContext, programUnit model.ProgramUnit) environment.EnvironmentDivision

var (
	NewDataDivision        NewDataDivisionFunc
	NewEnvironmentDivision NewEnvironmentDivisionFunc
)

func RegisterNewDataDivisionFunc(f NewDataDivisionFunc) {
	NewDataDivision = f
}

func RegisterNewEnvironmentDivisionFunc(f NewEnvironmentDivisionFunc) {
	NewEnvironmentDivision = f
}
