package communication

import "github.com/kaisawind/cobol/asg/model"

type CommunicationDescriptionEntry interface {
	model.CobolDivisionElement

	Name() string
}
