package call

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type SpecialRegisterType int

const (
	ADDRESS_OF SpecialRegisterType = iota
	DATE
	DAY
	DAY_OF_WEEK
	DEBUG_CONTENTS
	DEBUG_ITEM
	DEBUG_LINE
	DEBUG_NAME
	DEBUG_SUB_1
	DEBUG_SUB_2
	DEBUG_SUB_3
	LENGTH_OF
	LINAGE_COUNTER
	LINE_COUNTER
	PAGE_COUNTER
	RETURN_CODE
	SHIFT_IN
	SHIFT_OUT
	SORT_CONTROL
	SORT_CORE_SIZE
	SORT_FILE_SIZE
	SORT_MESSAGE
	SORT_MODE_SIZE
	SORT_RETURN
	TALLY
	TIME
	WHEN_COMPILED
)

type SpecialRegisterCall struct {
	ctx            antlr.ParserRuleContext
	typ            SpecialRegisterType
	identifierCall Call
}

func NewSpecialRegisterCall(ctx antlr.ParserRuleContext, typ SpecialRegisterType) *SpecialRegisterCall {
	return &SpecialRegisterCall{
		ctx: ctx,
		typ: typ,
	}
}

func (e *SpecialRegisterCall) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *SpecialRegisterCall) Name() string {
	return "" // WARN: ?
}

func (e *SpecialRegisterCall) Type() CallType {
	return SPECIAL_REGISTER_CALL
}

func (e *SpecialRegisterCall) SpecialRegisterType() SpecialRegisterType {
	return e.typ
}

func (e *SpecialRegisterCall) SetIdentifierCall(call Call) {
	e.identifierCall = call
}

func (e *SpecialRegisterCall) GetIdentifierCall() Call {
	return e.identifierCall
}
