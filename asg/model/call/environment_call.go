package call

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type EnvironmentCall struct {
	ctx  antlr.ParserRuleContext
	name string
}

func NewEnvironmentCall(ctx antlr.ParserRuleContext, name string) *EnvironmentCall {
	return &EnvironmentCall{
		ctx:  ctx,
		name: name,
	}
}

func (e *EnvironmentCall) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *EnvironmentCall) Name() string {
	return e.name
}

func (e *EnvironmentCall) Type() CallType {
	return ENVIRONMENT_CALL
}
