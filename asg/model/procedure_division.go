package model

import "github.com/kaisawind/cobol/gen/cobol85"

type ProcedureDivision struct {
	Element
}

func NewProcedureDivision(ctx *cobol85.ProcedureDivisionContext) *ProcedureDivision {
	return &ProcedureDivision{
		Element: NewBaseElement(ctx),
	}
}
