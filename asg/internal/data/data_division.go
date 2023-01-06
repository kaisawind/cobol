package data

import (
	"github.com/kaisawind/cobol/asg/internal"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/data"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type DataDivision struct {
	model.CobolDivision
	ctx *cobol85.DataDivisionContext
}

func init() {
	internal.RegisterNewDataDivisionFunc(NewDataDivision)
}

func NewDataDivision(ctx *cobol85.DataDivisionContext, programUnit model.ProgramUnit) data.DataDivision {
	return &DataDivision{
		CobolDivision: internal.NewCobolDivisionElement(ctx, programUnit),
		ctx:           ctx,
	}
}
