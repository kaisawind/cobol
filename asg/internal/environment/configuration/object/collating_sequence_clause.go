package object

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/object"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CollatingSequenceClause struct {
	model.CobolDivisionElement
	ctx           *cobol85.CollatingSequenceClauseContext
	alphabetNames []string
	alphanumeric  string
	national      string
}

func NewCollatingSequenceClause(ctx *cobol85.CollatingSequenceClauseContext, cobolDivisionElement model.CobolDivisionElement) object.CollatingSequenceClause {
	return &CollatingSequenceClause{
		CobolDivisionElement: cobolDivisionElement,
		ctx:                  ctx,
	}
}

func (e *CollatingSequenceClause) AddAlphabetName(name string) {
	e.alphabetNames = append(e.alphabetNames, name)
}

func (e *CollatingSequenceClause) GetAlphabetNames() []string {
	return e.alphabetNames
}

func (e *CollatingSequenceClause) SetAlphanumeric(alphanumeric string) {
	e.alphanumeric = alphanumeric
}

func (e *CollatingSequenceClause) GetAlphanumeric() string {
	return e.alphanumeric
}

func (e *CollatingSequenceClause) SetNational(national string) {
	e.national = national
}

func (e *CollatingSequenceClause) GetNational() string {
	return e.national
}
