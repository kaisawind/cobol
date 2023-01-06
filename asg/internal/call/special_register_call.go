package call

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type SpecialRegisterCall struct {
	call.Call

	ctx      *cobol85.SpecialRegisterContext
	callType call.CallType
}

func NewSpecialRegisterCall(
	ctx *cobol85.SpecialRegisterContext,
	name string,
	programUnit model.ProgramUnit,
) call.SpecialRegisterCall {
	return &SpecialRegisterCall{
		Call:     NewCall(ctx, "", programUnit),
		ctx:      ctx,
		callType: call.SPECIAL_REGISTER_CALL,
	}
}

func (e *SpecialRegisterCall) SpecialRegisterContext() *cobol85.SpecialRegisterContext {
	return e.ctx
}

func (e *SpecialRegisterCall) Type() call.CallType {
	return e.callType
}
