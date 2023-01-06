package object

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/object"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type SegmentLimitClause struct {
	model.CobolDivisionElement
	ctx *cobol85.SegmentLimitClauseContext

	integerLiteral model.IntegerLiteral
}

func NewSegmentLimitClause(ctx *cobol85.SegmentLimitClauseContext, cobolDivisionElement model.CobolDivisionElement) object.SegmentLimitClause {
	return &SegmentLimitClause{
		CobolDivisionElement: cobolDivisionElement,
		ctx:                  ctx,
	}
}

func (e *SegmentLimitClause) SetIntegerLiteral(integerLiteral model.IntegerLiteral) {
	e.integerLiteral = integerLiteral
}

func (e *SegmentLimitClause) GetIntegerLiteral() model.IntegerLiteral {
	return e.integerLiteral
}
