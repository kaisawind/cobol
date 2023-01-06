package call

import "github.com/kaisawind/cobol/asg/model/procedure"

type ProcedureCall interface {
	Call

	Paragraph() procedure.Paragraph
}
