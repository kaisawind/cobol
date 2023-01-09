package object

import (
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/object"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type DiskSizeClause struct {
	element.CobolDivisionElement
	ctx       *cobol85.DiskSizeClauseContext
	unit      object.DiskSizeClauseUnit
	valueStmt valuestmt.ValueStmt
}

func NewDiskSizeClause(ctx *cobol85.DiskSizeClauseContext, programUnit element.ProgramUnit) object.DiskSizeClause {
	return &DiskSizeClause{
		CobolDivisionElement: instances.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
	}
}

func (e *DiskSizeClause) SetUnit(unit object.DiskSizeClauseUnit) {
	e.unit = unit
}

func (e *DiskSizeClause) GetUnit() object.DiskSizeClauseUnit {
	return e.unit
}

func (e *DiskSizeClause) SetValueStmt(valueStmt valuestmt.ValueStmt) {
	e.valueStmt = valueStmt
}

func (e *DiskSizeClause) GetValueStmt() valuestmt.ValueStmt {
	return e.valueStmt
}
