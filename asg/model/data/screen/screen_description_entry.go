package screen

import "github.com/kaisawind/cobol/asg/model/element"

type ScreenDescriptionEntry interface {
	element.CobolDivisionElement

	Name() string
}
