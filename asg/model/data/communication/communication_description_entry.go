package communication

import "github.com/kaisawind/cobol/asg/model/element"

type CommunicationDescriptionEntry interface {
	element.CobolDivisionElement

	Name() string
}
