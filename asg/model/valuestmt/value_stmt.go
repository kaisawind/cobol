package valuestmt

import "github.com/kaisawind/cobol/asg/model"

type ValueStmt interface {
	model.CobolDivisionElement

	AddValueStmt(value ValueStmt)
	GetValueStmts() []ValueStmt
}
