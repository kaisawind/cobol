package object

import (
	"github.com/kaisawind/cobol/asg/internal"
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

func NewObjectComputerParagraph(ctx *cobol85.ObjectComputerParagraphContext, name string, programUnit model.ProgramUnit) object.ObjectComputerParagraph {
	return &ObjectComputerParagraph{
		CobolDivisionElement: internal.NewCobolDivisionElement(ctx, programUnit),
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
		ret = element.(object.CharacterSetClause)
	} else {
		ret = NewCharacterSetClause(ctx, e.ProgramUnit())
		e.characterSetClause = ret
		e.AddElement(ret)
	}
	return
}

func (e *ObjectComputerParagraph) GetCharacterSetClause() (ret object.CharacterSetClause) {
	return
}

func (e *ObjectComputerParagraph) SetCollatingSequenceClause(ctx *cobol85.CollatingSequenceClauseContext) (ret object.CollatingSequenceClause) {
	return
}

func (e *ObjectComputerParagraph) GetCollatingSequenceClause() (ret object.CollatingSequenceClause) {
	return
}

func (e *ObjectComputerParagraph) SetDiskSizeClause(ctx *cobol85.DiskSizeClauseContext) (ret object.DiskSizeClause) {
	return
}

func (e *ObjectComputerParagraph) GetDiskSizeClause() (ret object.DiskSizeClause) {
	return
}

func (e *ObjectComputerParagraph) SetMemorySizeClause(ctx *cobol85.MemorySizeClauseContext) (ret object.MemorySizeClause) {
	return
}

func (e *ObjectComputerParagraph) GetMemorySizeClause() (ret object.MemorySizeClause) {
	return
}

func (e *ObjectComputerParagraph) SetSegmentLimitClause(ctx *cobol85.SegmentLimitClauseContext) (ret object.SegmentLimitClause) {
	return
}

func (e *ObjectComputerParagraph) GetSegmentLimitClause() (ret object.SegmentLimitClause) {
	return
}
