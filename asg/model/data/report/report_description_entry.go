package report

import "github.com/kaisawind/cobol/asg/model/element"

type ReportDescriptionEntry interface {
	element.CobolDivisionElement

	Name() string
}
