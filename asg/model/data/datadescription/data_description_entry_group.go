package datadescription

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type DataDescriptionEntryGroup struct {
	ctx           cobol85.IDataDescriptionEntryFormat1Context
	name          string
	entries       []DataDescriptionEntry
	occursClauses []*OccursClause
}

func NewDataDescriptionEntryGroup(ctx cobol85.IDataDescriptionEntryFormat1Context, name string) *DataDescriptionEntryGroup {
	return &DataDescriptionEntryGroup{
		ctx:  ctx,
		name: name,
	}
}

func (e *DataDescriptionEntryGroup) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *DataDescriptionEntryGroup) Name() string {
	return e.name
}

func (e *DataDescriptionEntryGroup) Type() DDEType {
	return GROUP
}

func (e *DataDescriptionEntryGroup) AddDataDescriptionEntry(entry DataDescriptionEntry) {
	e.entries = append(e.entries, entry)
}

func (e *DataDescriptionEntryGroup) GetDataDescriptionEntries() []DataDescriptionEntry {
	return e.entries
}

func (e *DataDescriptionEntryGroup) AddOccursClause(clause *OccursClause) {
	e.occursClauses = append(e.occursClauses, clause)
}

func (e *DataDescriptionEntryGroup) GetOccursClauses() []*OccursClause {
	return e.occursClauses
}
