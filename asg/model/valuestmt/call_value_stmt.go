package valuestmt

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type CallValueStmt struct {
	*BaseValueStmt
	call any // call.Call
}

// NewCallValueStmt new call value stmt
//
//	prevent cycle import, any is call.Call
func NewCallValueStmt(ctx antlr.ParserRuleContext, call any) *CallValueStmt {
	return &CallValueStmt{
		BaseValueStmt: NewBaseValueStmt(ctx),
		call:          call,
	}
}

func (e *CallValueStmt) Call() any {
	return e.call
}
