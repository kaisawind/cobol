package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/element"
)

type Call struct {
	element.CobolDivisionElement
	name string
}

func NewCall(ctx antlr.ParserRuleContext, name string, programUnit element.ProgramUnit) call.Call {
	return &Call{
		CobolDivisionElement: instances.NewCobolDivision(ctx, programUnit),
		name:                 name,
	}
}

func (e *Call) Name() string {
	return e.name
}

func (e *Call) Type() call.CallType {
	return call.DEFAULT_CALL
}

func (e *Call) Unwrap() call.Call {
	return e
}
