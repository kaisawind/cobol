package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/communication"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type DataDivision struct {
	ctx                  *cobol85.DataDivisionContext
	communicationSection *communication.CommunicationSection
}

func NewDataDivision(ctx *cobol85.DataDivisionContext) *DataDivision {
	return &DataDivision{
		ctx: ctx,
	}
}

func (e *DataDivision) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *DataDivision) SetCommunicationSection(communicationSection *communication.CommunicationSection) {
	e.communicationSection = communicationSection
}

func (e *DataDivision) GetCommunicationSection() *communication.CommunicationSection {
	return e.communicationSection
}
