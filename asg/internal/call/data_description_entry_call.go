package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/asg/model/element"
)

type DataDescriptionEntryCall struct {
	call.Call

	dataDescriptionEntry datadescription.DataDescriptionEntry
	callType             call.CallType
}

func init() {
	instances.RegisterNewDataDescriptionEntryCallFunc(NewDataDescriptionEntryCall)
}

func NewDataDescriptionEntryCall(
	ctx antlr.ParserRuleContext,
	name string,
	dataDescriptionEntry datadescription.DataDescriptionEntry,
	programUnit element.ProgramUnit,
) call.DataDescriptionEntryCall {
	return &DataDescriptionEntryCall{
		Call:                 NewCall(ctx, name, programUnit),
		dataDescriptionEntry: dataDescriptionEntry,
		callType:             call.DATA_DESCRIPTION_ENTRY_CALL,
	}
}

func (e *DataDescriptionEntryCall) DataDescriptionEntry() datadescription.DataDescriptionEntry {
	return e.dataDescriptionEntry
}

func (e *DataDescriptionEntryCall) Type() call.CallType {
	return e.callType
}
