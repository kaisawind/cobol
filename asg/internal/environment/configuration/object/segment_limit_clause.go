package object

import (
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/object"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type SegmentLimitClause struct {
	element.CobolDivisionElement
	ctx *cobol85.SegmentLimitClauseContext

	integerLiteral element.IntegerLiteral
}

func NewSegmentLimitClause(ctx *cobol85.SegmentLimitClauseContext, programUnit element.ProgramUnit) object.SegmentLimitClause {
	return &SegmentLimitClause{
		CobolDivisionElement: instances.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
	}
}

func (e *SegmentLimitClause) SetIntegerLiteral(integerLiteral element.IntegerLiteral) {
	e.integerLiteral = integerLiteral
}

func (e *SegmentLimitClause) GetIntegerLiteral() element.IntegerLiteral {
	return e.integerLiteral
}
