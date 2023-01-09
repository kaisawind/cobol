package report

import "github.com/kaisawind/cobol/asg/model/element"

type ReportDescription interface {
	element.CobolDivisionElement

	Name() string
}
