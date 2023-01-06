package procedure

import "github.com/kaisawind/cobol/asg/model"

type Paragraph interface {
	model.Scope

	Name() string
}
