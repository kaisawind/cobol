package procedure

import "github.com/kaisawind/cobol/asg/model"

type Section interface {
	model.Scope

	Name() string
}
