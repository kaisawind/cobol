package procedure

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type Paragraph struct {
	ctx  cobol85.IParagraphContext
	name string
}

func NewParagraph(ctx cobol85.IParagraphContext, name string) *Paragraph {
	return &Paragraph{
		ctx:  ctx,
		name: name,
	}
}

func (e *Paragraph) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *Paragraph) Name() string {
	return e.name
}
