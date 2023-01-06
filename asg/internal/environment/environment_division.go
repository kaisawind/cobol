package environment

import (
	"github.com/kaisawind/cobol/asg/internal"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/environment"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type EnvironmentDivision struct {
	model.CobolDivision
	ctx *cobol85.EnvironmentDivisionContext
}

func init() {
	internal.RegisterNewEnvironmentDivisionFunc(NewEnvironmentDivision)
}

func NewEnvironmentDivision(ctx *cobol85.EnvironmentDivisionContext, programUnit model.ProgramUnit) environment.EnvironmentDivision {
	return &EnvironmentDivision{
		CobolDivision: internal.NewCobolDivisionElement(ctx, programUnit),
		ctx:           ctx,
	}
}
