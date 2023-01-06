package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
)

type UndefinedCall struct {
	call.Call

	callType call.CallType
}

func NewUndefinedCall(ctx antlr.ParserRuleContext, name string, programUnit model.ProgramUnit) call.Call {
	return &UndefinedCall{
		Call:     NewCall(ctx, name, programUnit),
		callType: call.UNDEFINED_CALL,
	}
}

func (e *UndefinedCall) Type() call.CallType {
	return e.callType
}
