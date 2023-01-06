package valuestmt

import (
	"github.com/kaisawind/cobol/asg/internal"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type Subscript struct {
	model.CobolDivisionElement
	ctx *cobol85.SubscriptContext
}

func NewSubscript(ctx *cobol85.SubscriptContext, programUnit model.ProgramUnit) valuestmt.Subscript {
	return &Subscript{
		CobolDivisionElement: internal.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
	}
}
