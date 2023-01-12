package model

import "github.com/kaisawind/cobol/gen/cobol85"

type EnvironmentDivision struct {
	Element
}

func NewEnvironmentDivision(ctx *cobol85.EnvironmentDivisionContext) *EnvironmentDivision {
	return &EnvironmentDivision{
		Element: NewBaseElement(ctx),
	}
}
