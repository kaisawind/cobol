package environment

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/environment"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type EnvironmentDivision struct {
	model.CobolDivision
	ctx *cobol85.EnvironmentDivisionContext
}

func NewEnvironmentDivision(ctx *cobol85.EnvironmentDivisionContext, cobolDivision model.CobolDivision) environment.EnvironmentDivision {
	return &EnvironmentDivision{
		CobolDivision: cobolDivision,
		ctx:           ctx,
	}
}
