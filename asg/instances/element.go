package instances

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/element"
)

type NewCobolDivisionFunc func(ctx antlr.ParserRuleContext, programUnit element.ProgramUnit) element.CobolDivision
type NewCobolDivisionElementFunc func(ctx antlr.ParserRuleContext, programUnit element.ProgramUnit) element.CobolDivisionElement

var (
	NewCobolDivision        NewCobolDivisionFunc
	NewCobolDivisionElement NewCobolDivisionElementFunc
)

func RegisterNewCobolDivisionFunc(f NewCobolDivisionFunc) {
	NewCobolDivision = f
}

func RegisterNewCobolDivisionElementFunc(f NewCobolDivisionElementFunc) {
	NewCobolDivisionElement = f
}
