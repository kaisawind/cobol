package model

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type Scope struct {
	*DivisionElement
}

func NewScope(ctx antlr.ParserRuleContext, programUnit *ProgramUnit) *Scope {
	return &Scope{
		DivisionElement: NewDivisionElement(ctx, programUnit),
	}
}
