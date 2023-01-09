package call

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/asg/model/element"
)

type IndexCall struct {
	call.Call

	index    datadescription.Index
	callType call.CallType
}

func NewIndexCall(
	ctx antlr.ParserRuleContext,
	name string,
	index datadescription.Index,
	programUnit element.ProgramUnit,
) call.IndexCall {
	return &IndexCall{
		Call:     NewCall(ctx, name, programUnit),
		index:    index,
		callType: call.INDEX_CALL,
	}
}

func (e *IndexCall) Index() datadescription.Index {
	return e.index
}

func (e *IndexCall) Type() call.CallType {
	return e.callType
}
