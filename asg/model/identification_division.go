package model

import "github.com/kaisawind/cobol/gen/cobol85"

type IdentificationDivision struct {
	Element
}

func NewIdentificationDivision(ctx *cobol85.IdentificationDivisionContext) *IdentificationDivision {
	return &IdentificationDivision{
		Element: NewBaseElement(ctx),
	}
}
