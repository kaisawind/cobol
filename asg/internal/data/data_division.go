package data

import (
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/data"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type DataDivision struct {
	element.CobolDivision
	ctx *cobol85.DataDivisionContext
}

func init() {
	instances.RegisterNewDataDivisionFunc(NewDataDivision)
}

func NewDataDivision(ctx *cobol85.DataDivisionContext, programUnit element.ProgramUnit) data.DataDivision {
	return &DataDivision{
		CobolDivision: instances.NewCobolDivision(ctx, programUnit),
		ctx:           ctx,
	}
}
