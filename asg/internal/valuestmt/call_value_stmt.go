package valuestmt

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
)

type CallValueStmt struct {
	valuestmt.ValueStmt
	call call.Call
}

func NewCallValueStmt(ctx antlr.ParserRuleContext, call call.Call, programUnit element.ProgramUnit) valuestmt.CallValueStmt {
	return &CallValueStmt{
		ValueStmt: NewValueStmt(ctx, programUnit),
		call:      call,
	}
}

func (e *CallValueStmt) Call() call.Call {
	return e.call
}
