package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/element"
)

type CallDelegate struct {
	element.CobolDivisionElement
	delegate call.Call
}

func init() {
	instances.RegisterNewCallDelegateFunc(NewCallDelegate)
}

func NewCallDelegate(ctx antlr.ParserRuleContext, delegate call.Call, programUnit element.ProgramUnit) call.Call {
	return &CallDelegate{
		CobolDivisionElement: instances.NewCobolDivisionElement(ctx, programUnit),
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
