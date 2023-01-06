package object

import "github.com/kaisawind/cobol/asg/model"

type CollatingSequenceClause interface {
	model.CobolDivisionElement

	AddAlphabetName(name string)
	GetAlphabetNames() []string
	SetAlphanumeric(alphanumeric string)
	GetAlphanumeric() string
	SetNational(national string)
	GetNational() string
}
