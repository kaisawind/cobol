package communication

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type SymbolicQueueClause struct {
	ctx          cobol85.ISymbolicQueueClauseContext
	dataDescCall call.Call
}

func NewSymbolicQueueClause(ctx cobol85.ISymbolicQueueClauseContext) *SymbolicQueueClause {
	return &SymbolicQueueClause{
		ctx: ctx,
	}
}

func (e *SymbolicQueueClause) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *SymbolicQueueClause) SetDataDescCall(dataDescCall call.Call) {
	e.dataDescCall = dataDescCall
}
