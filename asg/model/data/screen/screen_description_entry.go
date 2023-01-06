package screen

import "github.com/kaisawind/cobol/asg/model"

type ScreenDescriptionEntry interface {
	model.CobolDivisionElement

	Name() string
}
