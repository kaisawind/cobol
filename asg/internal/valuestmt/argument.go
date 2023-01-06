package valuestmt

import (
	"github.com/kaisawind/cobol/asg/internal"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type Argument struct {
	model.CobolDivisionElement
	ctx *cobol85.ArgumentContext
}

func NewArgument(ctx *cobol85.ArgumentContext, programUnit model.ProgramUnit) valuestmt.Argument {
	return &Argument{
		CobolDivisionElement: internal.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
	}
}
