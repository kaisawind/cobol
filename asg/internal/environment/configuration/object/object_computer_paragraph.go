package object

import (
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/object"
	"github.com/kaisawind/cobol/asg/util"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ObjectComputerParagraph struct {
	element.CobolDivisionElement
	ctx  *cobol85.ObjectComputerParagraphContext
	name string

	characterSetClause      object.CharacterSetClause
	collatingSequenceClause object.CollatingSequenceClause
	diskSizeClause          object.DiskSizeClause
	memorySizeClause        object.MemorySizeClause
	segmentLimitClause      object.SegmentLimitClause
}

func NewObjectComputerParagraph(ctx *cobol85.ObjectComputerParagraphContext, name string, programUnit element.ProgramUnit) object.ObjectComputerParagraph {
	return &ObjectComputerParagraph{
		CobolDivisionElement: instances.NewCobolDivisionElement(ctx, programUnit),
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

func (e *ObjectComputerParagraph) GetCharacterSetClause() object.CharacterSetClause {
	return e.characterSetClause
}

func (e *ObjectComputerParagraph) SetCollatingSequenceClause(ctx *cobol85.CollatingSequenceClauseContext) (ret object.CollatingSequenceClause) {
	element := e.GetElement(ctx)
	if element != nil {
		ret = element.(object.CollatingSequenceClause)
	} else {
		ret = NewCollatingSequenceClause(ctx, e.ProgramUnit())

		for _, v := range ctx.AllAlphabetName() {
			ret.AddAlphabetName(util.DetermineName(v))
		}

		if cctx := ctx.CollatingSequenceClauseAlphanumeric(); cctx != nil {
			if ictx, ok := cctx.(*cobol85.CollatingSequenceClauseAlphanumericContext); ok {
				ret.SetAlphanumeric(util.DetermineName(ictx.AlphabetName()))
			}
		}

		if cctx := ctx.CollatingSequenceClauseNational(); cctx != nil {
			if ictx, ok := cctx.(*cobol85.CollatingSequenceClauseNationalContext); ok {
				ret.SetNational(util.DetermineName(ictx.AlphabetName()))
			}
		}

		e.collatingSequenceClause = ret
		e.AddElement(ret)
	}
	return
}

func (e *ObjectComputerParagraph) GetCollatingSequenceClause() object.CollatingSequenceClause {
	return e.collatingSequenceClause
}

func (e *ObjectComputerParagraph) SetDiskSizeClause(ctx *cobol85.DiskSizeClauseContext) (ret object.DiskSizeClause) {
	element := e.GetElement(ctx)
	if element != nil {
		ret = element.(object.DiskSizeClause)
	} else {
		ret = NewDiskSizeClause(ctx, e.ProgramUnit())

		// TODO

		e.characterSetClause = ret
		e.AddElement(ret)
	}
	return
}

func (e *ObjectComputerParagraph) GetDiskSizeClause() object.DiskSizeClause {
	return e.diskSizeClause
}

func (e *ObjectComputerParagraph) SetMemorySizeClause(ctx *cobol85.MemorySizeClauseContext) (ret object.MemorySizeClause) {
	return
}

func (e *ObjectComputerParagraph) GetMemorySizeClause() object.MemorySizeClause {
	return e.memorySizeClause
}

func (e *ObjectComputerParagraph) SetSegmentLimitClause(ctx *cobol85.SegmentLimitClauseContext) (ret object.SegmentLimitClause) {
	return
}

func (e *ObjectComputerParagraph) GetSegmentLimitClause() object.SegmentLimitClause {
	return e.segmentLimitClause
}
