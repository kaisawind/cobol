package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/procedure"
	"github.com/kaisawind/cobol/asg/util"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ProcedureDivision struct {
	ctx           cobol85.IProcedureDivisionContext
	sections      []*procedure.Section
	sectionsTable map[string]*procedure.Section
}

func NewProcedureDivision(ctx cobol85.IProcedureDivisionContext) *ProcedureDivision {
	return &ProcedureDivision{
		ctx:           ctx,
		sections:      []*procedure.Section{},
		sectionsTable: map[string]*procedure.Section{},
	}
}

func (e *ProcedureDivision) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *ProcedureDivision) AddSection(section *procedure.Section) {
	e.sections = append(e.sections, section)
	name := util.Symbol(section.Name())
	e.sectionsTable[name] = section
}

func (e *ProcedureDivision) GetSection(name string) *procedure.Section {
	name = util.Symbol(name)
	return e.sectionsTable[name]
}

func (e *ProcedureDivision) GetSections() []*procedure.Section {
	return e.sections
}

func (e *ProcedureDivision) GetSectionsByName(name string) []*procedure.Section {
	name = util.Symbol(name)
	sections := []*procedure.Section{}
	for k, v := range e.sectionsTable {
		if k == name {
			sections = append(sections, v)
		}
	}
	return sections
}
