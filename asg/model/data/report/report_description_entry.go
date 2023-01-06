package report

import "github.com/kaisawind/cobol/asg/model"

type ReportDescriptionEntry interface {
	model.CobolDivisionElement

	Name() string
}
