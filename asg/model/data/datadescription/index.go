package datadescription

import "github.com/kaisawind/cobol/asg/model"

type Index interface {
	model.CobolDivisionElement

	Name() string
}
