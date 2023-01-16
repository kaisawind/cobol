package call

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type UndefinedCall struct {
	ctx  antlr.ParserRuleContext
	name string
}

func NewUndefinedCall(ctx antlr.ParserRuleContext, name string) *UndefinedCall {
	return &UndefinedCall{
		ctx:  ctx,
		name: name,
	}
}

func (e *UndefinedCall) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *UndefinedCall) Name() string {
	return e.name
}

func (e *UndefinedCall) Type() CallType {
	return UNDEFINED_CALL
}
