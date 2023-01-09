package datadescription

import "github.com/kaisawind/cobol/asg/model/element"

type Index interface {
	element.CobolDivisionElement

	Name() string
}
