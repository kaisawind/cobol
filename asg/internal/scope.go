package internal

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
)

type Scope struct {
	model.CobolDivisionElement
}

func NewScope(ctx antlr.ParserRuleContext, programUnit model.ProgramUnit) model.Scope {
	return &Scope{
		CobolDivisionElement: NewCobolDivisionElement(ctx, programUnit),
	}
}
