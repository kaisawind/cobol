package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
)

type FunctionCall struct {
	call.Call

	callType call.CallType
}

func NewFunctionCall(
	ctx antlr.ParserRuleContext,
	name string,
	programUnit model.ProgramUnit,
) call.FunctionCall {
	return &FunctionCall{
		Call:     NewCall(ctx, name, programUnit),
		callType: call.FUNCTION_CALL,
	}
}

func (e *FunctionCall) Type() call.CallType {
	return e.callType
}
