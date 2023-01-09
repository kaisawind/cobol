package procedure

import "github.com/kaisawind/cobol/asg/model/element"

type Paragraph interface {
	element.Scope

	Name() string
}
