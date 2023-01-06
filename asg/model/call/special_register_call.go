package call

import "github.com/kaisawind/cobol/gen/cobol85"

type SpecialRegisterCall interface {
	Call

	SpecialRegisterContext() *cobol85.SpecialRegisterContext
}
