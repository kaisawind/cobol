package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/internal"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
)

type Call struct {
	model.CobolDivisionElement
	name string
}

func NewCall(ctx antlr.ParserRuleContext, name string, programUnit model.ProgramUnit) call.Call {
	return &Call{
		CobolDivisionElement: internal.NewCobolDivision(ctx, programUnit),
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
