package object

import (
	"github.com/kaisawind/cobol/asg/internal"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/object"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type SegmentLimitClause struct {
	model.CobolDivisionElement
	ctx *cobol85.SegmentLimitClauseContext

	integerLiteral model.IntegerLiteral
}

func NewSegmentLimitClause(ctx *cobol85.SegmentLimitClauseContext, programUnit model.ProgramUnit) object.SegmentLimitClause {
	return &SegmentLimitClause{
		CobolDivisionElement: internal.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
	}
}

func (e *SegmentLimitClause) SetIntegerLiteral(integerLiteral model.IntegerLiteral) {
	e.integerLiteral = integerLiteral
}

func (e *SegmentLimitClause) GetIntegerLiteral() model.IntegerLiteral {
	return e.integerLiteral
}
