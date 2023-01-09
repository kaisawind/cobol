package object

import "github.com/kaisawind/cobol/asg/model/element"

type CollatingSequenceClause interface {
	element.CobolDivisionElement

	AddAlphabetName(name string)
	GetAlphabetNames() []string
	SetAlphanumeric(alphanumeric string)
	GetAlphanumeric() string
	SetNational(national string)
	GetNational() string
}
