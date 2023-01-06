package call

import "github.com/kaisawind/cobol/asg/model/data/datadescription"

type DataDescriptionEntryCall interface {
	Call

	DataDescriptionEntry() datadescription.DataDescriptionEntry
}
