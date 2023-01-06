package datadescription

import "github.com/kaisawind/cobol/asg/model"

type DataDescriptionEntry interface {
	model.CobolDivisionElement

	Name() string
}
