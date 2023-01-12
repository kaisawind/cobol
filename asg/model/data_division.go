package model

import "github.com/kaisawind/cobol/gen/cobol85"

type DataDivision struct {
	Element
}

func NewDataDivision(ctx *cobol85.DataDivisionContext) *DataDivision {
	return &DataDivision{
		Element: NewBaseElement(ctx),
	}
}
