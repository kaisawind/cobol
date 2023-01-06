package valuestmt

import "github.com/kaisawind/cobol/asg/model/call"

type CallValueStmt interface {
	ValueStmt

	Call() call.Call
}
