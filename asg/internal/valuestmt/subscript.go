package valuestmt

import (
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type Subscript struct {
	element.CobolDivisionElement
	ctx *cobol85.SubscriptContext
}

func NewSubscript(ctx *cobol85.SubscriptContext, programUnit element.ProgramUnit) valuestmt.Subscript {
	return &Subscript{
		CobolDivisionElement: instances.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
	}
}
