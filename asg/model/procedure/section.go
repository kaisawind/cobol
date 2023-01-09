package procedure

import "github.com/kaisawind/cobol/asg/model/element"

type Section interface {
	element.Scope

	Name() string
}
