package object

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
)

type MemorySizeClauseUnit int

const (
	MemorySizeClause_CHARACTERS MemorySizeClauseUnit = iota
	MemorySizeClause_MODULES
	MemorySizeClause_WORDS
)

type MemorySizeClause interface {
	model.CobolDivisionElement

	SetUnit(unit MemorySizeClauseUnit)
	GetUnit() MemorySizeClauseUnit
	SetValueStmt(valueStmt valuestmt.ValueStmt)
	GetValueStmt() valuestmt.ValueStmt
}
