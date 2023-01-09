package instances

import (
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/environment"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type NewEnvironmentDivisionFunc func(ctx *cobol85.EnvironmentDivisionContext, programUnit element.ProgramUnit) environment.EnvironmentDivision

var (
	NewEnvironmentDivision NewEnvironmentDivisionFunc
)

func RegisterNewEnvironmentDivisionFunc(f NewEnvironmentDivisionFunc) {
	NewEnvironmentDivision = f
}
