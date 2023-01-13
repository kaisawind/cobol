package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
)

type IndexCall struct {
	ctx   antlr.ParserRuleContext
	name  string
	index *datadescription.Index
}

func NewIndexCall(ctx antlr.ParserRuleContext, name string, index *datadescription.Index) *IndexCall {
	return &IndexCall{
		ctx:   ctx,
		name:  name,
		index: index,
	}
}

func (e *IndexCall) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *IndexCall) Name() string {
	return e.name
}

func (e *IndexCall) Index() *datadescription.Index {
	return e.index
}

func (e *IndexCall) Type() CallType {
	return INDEX_CALL
}
