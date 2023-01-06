package report

import "github.com/kaisawind/cobol/asg/model"

type ReportDescription interface {
	model.CobolDivisionElement

	Name() string
}
