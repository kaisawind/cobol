package valuestmt

import "github.com/kaisawind/cobol/asg/model/element"

type ValueStmt interface {
	element.CobolDivisionElement

	AddValueStmt(value ValueStmt)
	GetValueStmts() []ValueStmt
}
