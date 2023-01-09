package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/element"
)

type MnemonicCall struct {
	call.Call

	callType call.CallType
}

func NewMnemonicCall(
	ctx antlr.ParserRuleContext,
	name string,
	programUnit element.ProgramUnit,
) call.MnemonicCall {
	return &MnemonicCall{
		Call:     NewCall(ctx, name, programUnit),
		callType: call.MNEMONIC_CALL,
	}
}

func (e *MnemonicCall) Type() call.CallType {
	return e.callType
}
