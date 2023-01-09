package object

import (
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/object"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type MemorySizeClause struct {
	element.CobolDivisionElement
	ctx       *cobol85.MemorySizeClauseContext
	unit      object.MemorySizeClauseUnit
	valueStmt valuestmt.ValueStmt
}

func NewMemorySizeClause(ctx *cobol85.MemorySizeClauseContext, programUnit element.ProgramUnit) object.MemorySizeClause {
	return &MemorySizeClause{
		CobolDivisionElement: instances.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
	}
}

func (e *MemorySizeClause) SetUnit(unit object.MemorySizeClauseUnit) {
	e.unit = unit
}

func (e *MemorySizeClause) GetUnit() object.MemorySizeClauseUnit {
	return e.unit
}

func (e *MemorySizeClause) SetValueStmt(valueStmt valuestmt.ValueStmt) {
	e.valueStmt = valueStmt
}

func (e *MemorySizeClause) GetValueStmt() valuestmt.ValueStmt {
	return e.valueStmt
}
