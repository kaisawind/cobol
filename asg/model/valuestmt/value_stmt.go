package valuestmt

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type ValueStmt interface {
	AddSubValueStmt(vs ValueStmt)
	GetSubValueStmts() []ValueStmt
}

type BaseValueStmt struct {
	ctx  antlr.ParserRuleContext
	subs []ValueStmt
}

func NewBaseValueStmt(ctx antlr.ParserRuleContext) *BaseValueStmt {
	return &BaseValueStmt{
		ctx: ctx,
	}
}

func (e *BaseValueStmt) AddSubValueStmt(vs ValueStmt) {
	e.subs = append(e.subs, vs)
}

func (e *BaseValueStmt) GetSubValueStmts() []ValueStmt {
	return e.subs
}

func (e *BaseValueStmt) Context() antlr.ParserRuleContext {
	return e.ctx
}
