package instances

import (
	"github.com/kaisawind/cobol/asg/model/data"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type NewDataDivisionFunc func(ctx *cobol85.DataDivisionContext, programUnit element.ProgramUnit) data.DataDivision

var (
	NewDataDivision NewDataDivisionFunc
)

func RegisterNewDataDivisionFunc(f NewDataDivisionFunc) {
	NewDataDivision = f
}
