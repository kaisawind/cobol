package procedure

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/util"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type Section struct {
	ctx  cobol85.IProcedureSectionContext
	name string

	paragraphs      []*Paragraph
	paragraphsTable map[string]*Paragraph
}

func NewSection(ctx cobol85.IProcedureSectionContext, name string) *Section {
	return &Section{
		ctx:             ctx,
		name:            name,
		paragraphs:      []*Paragraph{},
		paragraphsTable: map[string]*Paragraph{},
	}
}

func (e *Section) Name() string {
	return e.name
}

func (e *Section) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *Section) AddParagraph(paragraph *Paragraph) {
	e.paragraphs = append(e.paragraphs, paragraph)
	name := util.Symbol(paragraph.Name())
	e.paragraphsTable[name] = paragraph
}

func (e *Section) GetParagraphByName(name string) *Paragraph {
	return e.paragraphsTable[name]
}

func (e *Section) GetParagraphs() []*Paragraph {
	return e.paragraphs
}

func (e *Section) GetParagraphsByName(name string) []*Paragraph {
	name = util.Symbol(name)
	paragraphs := []*Paragraph{}
	for k, v := range e.paragraphsTable {
		if k == name {
			paragraphs = append(paragraphs, v)
		}
	}
	return paragraphs
}
