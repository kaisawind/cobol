package datadescription

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type Index struct {
	ctx   cobol85.IIndexNameContext
	name  string
	calls []any // avoid cycle import use any
}

func NewIndex(ctx cobol85.IIndexNameContext, name string) *Index {
	return &Index{
		ctx:   ctx,
		name:  name,
		calls: []any{},
	}
}

func (e *Index) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *Index) Name() string {
	return e.name
}

func (e *Index) AddIndexCall(call any) {
	e.calls = append(e.calls, call)
}

func (e *Index) GetIndexCalls() []any {
	return e.calls
}
