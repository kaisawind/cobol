package valuestmt

import (
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type Argument struct {
	element.CobolDivisionElement
	ctx *cobol85.ArgumentContext
}

func NewArgument(ctx *cobol85.ArgumentContext, programUnit element.ProgramUnit) valuestmt.Argument {
	return &Argument{
		CobolDivisionElement: instances.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
	}
}
