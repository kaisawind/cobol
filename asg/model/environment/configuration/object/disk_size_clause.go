package object

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
)

type DiskSizeClauseUnit int

const (
	DiskSizeClause_MODULES DiskSizeClauseUnit = iota
	DiskSizeClause_WORDS
)

type DiskSizeClause interface {
	model.CobolDivisionElement

	SetUnit(unit DiskSizeClauseUnit)
	GetUnit() DiskSizeClauseUnit
	SetValueStmt(valueStmt valuestmt.ValueStmt)
	GetValueStmt() valuestmt.ValueStmt
}
