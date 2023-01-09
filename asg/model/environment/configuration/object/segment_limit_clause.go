package object

import (
	"github.com/kaisawind/cobol/asg/model/element"
)

type SegmentLimitClause interface {
	element.CobolDivisionElement

	SetIntegerLiteral(integerLiteral element.IntegerLiteral)
	GetIntegerLiteral() element.IntegerLiteral
}
