package call

import (
	"github.com/kaisawind/cobol/asg/internal/element"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type TableCall struct {
	call.DataDescriptionEntryCall

	ctx      *cobol85.TableCallContext
	callType call.CallType
}

func NewTableCall(
	ctx *cobol85.TableCallContext,
	name string,
	dataDescriptionEntry datadescription.DataDescriptionEntry,
	programUnit element.ProgramUnit,
) call.TableCall {
	return &TableCall{
		DataDescriptionEntryCall: NewDataDescriptionEntryCall(ctx, name, dataDescriptionEntry, programUnit),
		ctx:                      ctx,
		callType:                 call.TABLE_CALL,
	}
}

func (e *TableCall) TableCallContext() *cobol85.TableCallContext {
	return e.ctx
}

func (e *TableCall) Type() call.CallType {
	return e.callType
}
