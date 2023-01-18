package valuestmt

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type Subscript struct {
	*BaseValueStmt
	subscript ValueStmt
}

func NewSubscript(ctx antlr.ParserRuleContext) *Subscript {
	return &Subscript{
		BaseValueStmt: NewBaseValueStmt(ctx),
	}
}

func (e *Subscript) SetSubscriptValueStmt(vs ValueStmt) {
	e.subscript = vs
}

func (e *Subscript) GetSubscriptValueStmt() ValueStmt {
	return e.subscript
}
