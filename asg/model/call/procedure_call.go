package call

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type ProcedureCall struct {
	ctx  antlr.ParserRuleContext
	name string
}

func NewProcedureCall(ctx antlr.ParserRuleContext, name string) *ProcedureCall {
	return &ProcedureCall{
		ctx:  ctx,
		name: name,
	}
}

func (e *ProcedureCall) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *ProcedureCall) Name() string {
	return e.name
}

func (e *ProcedureCall) Type() CallType {
	return PROCEDURE_CALL
}
