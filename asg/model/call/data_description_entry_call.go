package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
)

type DataDescriptionEntryCall struct {
	ctx                  antlr.ParserRuleContext
	name                 string
	dataDescriptionEntry datadescription.DataDescriptionEntry
}

func NewDataDescriptionEntryCall(ctx antlr.ParserRuleContext, name string, dataDescriptionEntry datadescription.DataDescriptionEntry) *DataDescriptionEntryCall {
	return &DataDescriptionEntryCall{
		ctx:                  ctx,
		name:                 name,
		dataDescriptionEntry: dataDescriptionEntry,
	}
}

func (e *DataDescriptionEntryCall) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *DataDescriptionEntryCall) Name() string {
	return e.name
}

func (e *DataDescriptionEntryCall) Type() CallType {
	return DATA_DESCRIPTION_ENTRY_CALL
}
