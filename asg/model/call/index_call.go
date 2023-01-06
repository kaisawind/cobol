package call

import "github.com/kaisawind/cobol/asg/model/data/datadescription"

type IndexCall interface {
	Call

	Index() datadescription.Index
}
