package datadescription

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type OccursClause struct {
	ctx           cobol85.IDataOccursClauseContext
	occursIndexed *OccursIndexed
}

func NewOccursClause(ctx cobol85.IDataOccursClauseContext) *OccursClause {
	return &OccursClause{
		ctx: ctx,
	}
}

func (e *OccursClause) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *OccursClause) SetOccursIndexed(occursIndexed *OccursIndexed) {
	e.occursIndexed = occursIndexed
}

func (e *OccursClause) GetOccursIndexed() *OccursIndexed {
	return e.occursIndexed
}
