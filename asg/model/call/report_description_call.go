package call

import "github.com/kaisawind/cobol/asg/model/data/report"

type ReportDescriptionCall interface {
	Call

	ReportDescription() report.ReportDescription
}
