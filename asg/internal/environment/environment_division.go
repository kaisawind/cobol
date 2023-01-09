package environment

import (
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/environment"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type EnvironmentDivision struct {
	element.CobolDivision
	ctx *cobol85.EnvironmentDivisionContext
}

func init() {
	instances.RegisterNewEnvironmentDivisionFunc(NewEnvironmentDivision)
}

func NewEnvironmentDivision(ctx *cobol85.EnvironmentDivisionContext, programUnit element.ProgramUnit) environment.EnvironmentDivision {
	return &EnvironmentDivision{
		CobolDivision: instances.NewCobolDivisionElement(ctx, programUnit),
		ctx:           ctx,
	}
}
