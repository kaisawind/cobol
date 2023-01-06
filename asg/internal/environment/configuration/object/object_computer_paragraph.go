package object

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/object"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ObjectComputerParagraph struct {
	model.CobolDivisionElement
	ctx  *cobol85.ObjectComputerParagraphContext
	name string

	characterSetClause      object.CharacterSetClause
	collatingSequenceClause object.CollatingSequenceClause
	diskSizeClause          object.DiskSizeClause
	memorySizeClause        object.MemorySizeClause
	segmentLimitClause      object.SegmentLimitClause
}

func NewObjectComputerParagraph(ctx *cobol85.ObjectComputerParagraphContext, name string, cobolDivisionElement model.CobolDivisionElement) object.ObjectComputerParagraph {
	return &ObjectComputerParagraph{
		CobolDivisionElement: cobolDivisionElement,
		ctx:                  ctx,
		name:                 name,
	}
}

func (e *ObjectComputerParagraph) Name() string {
	return e.name
}

func (e *ObjectComputerParagraph) SetCharacterSetClause(ctx *cobol85.CharacterSetClauseContext) (ret object.CharacterSetClause) {
	element := e.GetElement(ctx)
	if element != nil {

	}
}

func (e *ObjectComputerParagraph) GetCharacterSetClause() (ret object.CharacterSetClause) {
	return
}

func (e *ObjectComputerParagraph) SetCollatingSequenceClause(ctx *cobol85.CollatingSequenceClauseContext) (ret object.CollatingSequenceClause) {

}

func (e *ObjectComputerParagraph) GetCollatingSequenceClause() (ret object.CollatingSequenceClause) {
	return
}

func (e *ObjectComputerParagraph) SetDiskSizeClause(ctx *cobol85.DiskSizeClauseContext) (ret object.DiskSizeClause) {

}

func (e *ObjectComputerParagraph) GetDiskSizeClause() (ret object.DiskSizeClause) {
	return
}

func (e *ObjectComputerParagraph) SetMemorySizeClause(ctx *cobol85.MemorySizeClauseContext) (ret object.MemorySizeClause) {

}

func (e *ObjectComputerParagraph) GetMemorySizeClause() (ret object.MemorySizeClause) {
	return
}

func (e *ObjectComputerParagraph) SetSegmentLimitClause(ctx *cobol85.SegmentLimitClauseContext) (ret object.SegmentLimitClause) {

}

func (e *ObjectComputerParagraph) GetSegmentLimitClause() (ret object.SegmentLimitClause) {
	return
}
