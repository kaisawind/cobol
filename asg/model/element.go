package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Element interface {
	Context() antlr.ParserRuleContext
}

type BaseElement struct {
	ctx antlr.ParserRuleContext
}

func NewBaseElement(ctx antlr.ParserRuleContext) *BaseElement {
	return &BaseElement{
		ctx: ctx,
	}
}

func (e *BaseElement) Context() antlr.ParserRuleContext {
	return e.ctx
}
