package source

import (
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/source"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type SourceComputerParagraph struct {
	element.CobolDivisionElement
	ctx           *cobol85.SourceComputerParagraphContext
	name          string
	debuggingMode bool
}

func NewSourceComputerParagraph(ctx *cobol85.SourceComputerParagraphContext, name string, programUnit element.ProgramUnit) source.SourceComputerParagraph {
	return &SourceComputerParagraph{
		CobolDivisionElement: instances.NewCobolDivisionElement(ctx, programUnit),
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
