package internal

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/data"
	modelInternal "github.com/kaisawind/cobol/asg/model/internal"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type DataDivision struct {
	model.CobolDivision
	ctx *cobol85.DataDivisionContext
}

func NewDataDivision(ctx *cobol85.DataDivisionContext, programUnit model.ProgramUnit) data.DataDivision {
	return &DataDivision{
		CobolDivision: modelInternal.NewCobolDivision(ctx, programUnit),
		ctx:           ctx,
	}
}
