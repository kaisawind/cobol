package call

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type FunctionCall struct {
	ctx  antlr.ParserRuleContext
	name string
}

func NewFunctionCall(ctx antlr.ParserRuleContext, name string) *FunctionCall {
	return &FunctionCall{
		ctx:  ctx,
		name: name,
	}
}

func (e *FunctionCall) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *FunctionCall) Name() string {
	return e.name
}

func (e *FunctionCall) Type() CallType {
	return FUNCTION_CALL
}
