package call

import "github.com/kaisawind/cobol/asg/model/data/screen"

type ScreenDescriptionEntryCall interface {
	Call

	ScreenDescriptionEntry() screen.ScreenDescriptionEntry
}
