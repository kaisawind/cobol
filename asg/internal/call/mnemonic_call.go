package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
)

type MnemonicCall struct {
	call.Call

	callType call.CallType
}

func NewMnemonicCall(
	ctx antlr.ParserRuleContext,
	name string,
	programUnit model.ProgramUnit,
) call.MnemonicCall {
	return &MnemonicCall{
		Call:     NewCall(ctx, name, programUnit),
		callType: call.MNEMONIC_CALL,
	}
}

func (e *MnemonicCall) Type() call.CallType {
	return e.callType
}
