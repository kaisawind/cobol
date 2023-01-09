package instances

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/element"
)

type NewCallDelegateFunc func(ctx antlr.ParserRuleContext, delegate call.Call, programUnit element.ProgramUnit) call.Call

var (
	NewCallDelegate NewCallDelegateFunc
)

func RegisterNewCallDelegateFunc(f NewCallDelegateFunc) {
	NewCallDelegate = f
}
