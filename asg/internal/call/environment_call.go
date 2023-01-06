package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
)

type EnvironmentCall struct {
	call.Call

	callType call.CallType
}

func NewEnvironmentCall(
	ctx antlr.ParserRuleContext,
	name string,
	programUnit model.ProgramUnit,
) call.EnvironmentCall {
	return &EnvironmentCall{
		Call:     NewCall(ctx, name, programUnit),
		callType: call.ENVIRONMENT_CALL,
	}
}

func (e *EnvironmentCall) Type() call.CallType {
	return e.callType
}
