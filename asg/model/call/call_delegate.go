package call

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type CallDelegate struct {
	ctx      antlr.ParserRuleContext
	delegate Call
}

func NewCallDelegate(ctx antlr.ParserRuleContext, delegate Call) *CallDelegate {
	return &CallDelegate{
		ctx:      ctx,
		delegate: delegate,
	}
}

func (e *CallDelegate) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *CallDelegate) Name() string {
	return e.delegate.Name()
}

func (e *CallDelegate) Type() CallType {
	return e.delegate.Type()
}
