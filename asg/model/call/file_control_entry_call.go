package call

import "github.com/kaisawind/cobol/asg/model/environment/inputoutput/filecontrol"

type FileControlEntryCall interface {
	Call

	FileControlEntry() filecontrol.FileControlEntry
}
