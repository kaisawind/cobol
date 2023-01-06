package object

import "github.com/kaisawind/cobol/asg/model"

type SegmentLimitClause interface {
	model.CobolDivisionElement

	SetIntegerLiteral(integerLiteral model.IntegerLiteral)
	GetIntegerLiteral() model.IntegerLiteral
}
