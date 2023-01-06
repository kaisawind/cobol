package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/internal"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
)

type CallDelegate struct {
	model.CobolDivisionElement
	delegate call.Call
}

func init() {
	internal.RegisterNewCallDelegateFunc(NewCallDelegate)
}

func NewCallDelegate(ctx antlr.ParserRuleContext, delegate call.Call, programUnit model.ProgramUnit) call.Call {
	return &CallDelegate{
		CobolDivisionElement: internal.NewCobolDivisionElement(ctx, programUnit),
		delegate:             delegate,
	}
}

func (e *CallDelegate) Name() string {
	return e.delegate.Name()
}

func (e *CallDelegate) Type() call.CallType {
	return e.delegate.Type()
}

func (e *CallDelegate) Unwrap() call.Call {
	return e.delegate.Unwrap()
}
