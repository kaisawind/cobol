package instances

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type NewDataDivisionFunc func(ctx *cobol85.DataDivisionContext, programUnit element.ProgramUnit) data.DataDivision
type NewDataDescriptionEntryCallFunc func(
	ctx antlr.ParserRuleContext,
	name string,
	dataDescriptionEntry datadescription.DataDescriptionEntry,
	programUnit element.ProgramUnit,
) call.DataDescriptionEntryCall

var (
	NewDataDivision             NewDataDivisionFunc
	NewDataDescriptionEntryCall NewDataDescriptionEntryCallFunc
)

func RegisterNewDataDivisionFunc(f NewDataDivisionFunc) {
	NewDataDivision = f
}

func RegisterNewDataDescriptionEntryCallFunc(f NewDataDescriptionEntryCallFunc) {
	NewDataDescriptionEntryCall = f
}
