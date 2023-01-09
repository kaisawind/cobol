package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/environment/inputoutput/filecontrol"
)

type FileControlEntryCall struct {
	call.Call
	fileControlEntry filecontrol.FileControlEntry
	callType         call.CallType
}

func NewFileControlEntryCall(
	ctx antlr.ParserRuleContext,
	name string,
	fileControlEntry filecontrol.FileControlEntry,
	programUnit element.ProgramUnit,
) call.FileControlEntryCall {
	return &FileControlEntryCall{
		Call:             NewCall(ctx, name, programUnit),
		fileControlEntry: fileControlEntry,
		callType:         call.FILE_CONTROL_ENTRY_CALL,
	}
}

func (e *FileControlEntryCall) FileControlEntry() filecontrol.FileControlEntry {
	return e.fileControlEntry
}

func (e *FileControlEntryCall) Type() call.CallType {
	return e.callType
}
