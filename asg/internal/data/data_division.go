package data

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/data"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type DataDivision struct {
	model.CobolDivision
	ctx *cobol85.DataDivisionContext
}

func NewDataDivision(ctx *cobol85.DataDivisionContext, cobolDivision model.CobolDivision) data.DataDivision {
	return &DataDivision{
		CobolDivision: cobolDivision,
		ctx:           ctx,
	}
}
