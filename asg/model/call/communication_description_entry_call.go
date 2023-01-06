package call

import "github.com/kaisawind/cobol/asg/model/data/communication"

type CommunicationDescriptionEntryCall interface {
	Call

	CommunicationDescriptionEntry() communication.CommunicationDescriptionEntry
}
