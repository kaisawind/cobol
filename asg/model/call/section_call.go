package call

import "github.com/kaisawind/cobol/asg/model/procedure"

type SectionCall interface {
	Call

	Section() procedure.Section
}
