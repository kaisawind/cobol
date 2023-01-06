package source

import (
	"github.com/kaisawind/cobol/asg/internal"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/source"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type SourceComputerParagraph struct {
	model.CobolDivisionElement
	ctx           *cobol85.SourceComputerParagraphContext
	name          string
	debuggingMode bool
}

func NewSourceComputerParagraph(ctx *cobol85.SourceComputerParagraphContext, name string, programUnit model.ProgramUnit) source.SourceComputerParagraph {
	return &SourceComputerParagraph{
		CobolDivisionElement: internal.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
		name:                 name,
	}
}

func (e *SourceComputerParagraph) Name() string {
	return e.name
}

func (e *SourceComputerParagraph) IsDebuggingMode() bool {
	return e.debuggingMode
}

func (e *SourceComputerParagraph) SetDebuggingMode(debuggingMode bool) {
	e.debuggingMode = debuggingMode
}
