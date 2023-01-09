package element

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/element"
)

type Scope struct {
	element.CobolDivisionElement
}

func NewScope(ctx antlr.ParserRuleContext, programUnit element.ProgramUnit) element.Scope {
	return &Scope{
		CobolDivisionElement: NewCobolDivisionElement(ctx, programUnit),
	}
}
