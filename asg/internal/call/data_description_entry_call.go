package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
)

type DataDescriptionEntryCall struct {
	call.Call

	dataDescriptionEntry datadescription.DataDescriptionEntry
	callType             call.CallType
}

func NewDataDescriptionEntryCall(
	ctx antlr.ParserRuleContext,
	name string,
	dataDescriptionEntry datadescription.DataDescriptionEntry,
	programUnit model.ProgramUnit,
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
