package datadescription

import "github.com/kaisawind/cobol/asg/model/element"

type DataDescriptionEntry interface {
	element.CobolDivisionElement

	Name() string
}
