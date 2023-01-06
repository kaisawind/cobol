package call

import "github.com/kaisawind/cobol/asg/model/data/report"

type ReportDescriptionEntryCall interface {
	Call

	ReportDescriptionEntry() report.ReportDescriptionEntry
}
