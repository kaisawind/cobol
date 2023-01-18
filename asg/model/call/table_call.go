package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
)

type TableCall struct {
	ctx                  antlr.ParserRuleContext
	name                 string
	dataDescriptionEntry datadescription.DataDescriptionEntry
	subscripts           []*valuestmt.Subscript
}

func NewTableCall(ctx antlr.ParserRuleContext, name string, dataDescriptionEntry datadescription.DataDescriptionEntry) *TableCall {
	return &TableCall{
		ctx:                  ctx,
		name:                 name,
		dataDescriptionEntry: dataDescriptionEntry,
	}
}

func (e *TableCall) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *TableCall) Name() string {
	return e.name
}

func (e *TableCall) Type() CallType {
	return TABLE_CALL
}

func (e *TableCall) AddSubscript(subscript *valuestmt.Subscript) {
	e.subscripts = append(e.subscripts, subscript)
}

func (e *TableCall) GetSubscript() []*valuestmt.Subscript {
	return e.subscripts
}
