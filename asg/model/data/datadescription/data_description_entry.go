package datadescription

import "github.com/kaisawind/cobol/asg/util"

type DDEType int

const (
	CONDITION DDEType = iota
	EXEC_SQL
	GROUP
	RENAME
	SCALAR
)

type DataDescriptionEntry interface {
	Type() DDEType
	Name() string

	// AddCall add *call.DataDescriptionEntryCall
	AddCall(call any)
	// GetCalls []*call.DataDescriptionEntryCall
	GetCalls() []any
	SetDataDescriptionEntryGroup(dataDescriptionEntryGroup *DataDescriptionEntryGroup)
	GetDataDescriptionEntryGroup() *DataDescriptionEntryGroup
}

type BaseDataDescriptionEntry struct {
	dataDescriptionEntryGroup   *DataDescriptionEntryGroup
	dataDescriptionEntries      []DataDescriptionEntry
	dataDescriptionEntriesTable map[string]DataDescriptionEntry
	calls                       []any
}

func NewBaseDataDescriptionEntry() *BaseDataDescriptionEntry {
	return &BaseDataDescriptionEntry{
		dataDescriptionEntries:      []DataDescriptionEntry{},
		dataDescriptionEntriesTable: map[string]DataDescriptionEntry{},
		calls:                       []any{},
	}
}

func (e *BaseDataDescriptionEntry) AddCall(call any) {
	e.calls = append(e.calls, call)
}

func (e *BaseDataDescriptionEntry) GetCalls() []any {
	return e.calls
}

func (e *BaseDataDescriptionEntry) SetDataDescriptionEntryGroup(dataDescriptionEntryGroup *DataDescriptionEntryGroup) {
	e.dataDescriptionEntryGroup = dataDescriptionEntryGroup
}

func (e *BaseDataDescriptionEntry) GetDataDescriptionEntryGroup() *DataDescriptionEntryGroup {
	return e.dataDescriptionEntryGroup
}

func (e *BaseDataDescriptionEntry) AddDataDescriptionEntry(entry DataDescriptionEntry) {
	e.dataDescriptionEntries = append(e.dataDescriptionEntries, entry)
	name := util.Symbol(entry.Name())
	e.dataDescriptionEntriesTable[name] = entry
}

func (e *BaseDataDescriptionEntry) GetDataDescriptionEntries() []DataDescriptionEntry {
	return e.dataDescriptionEntries
}

func (e *BaseDataDescriptionEntry) GetDataDescriptionEntriesByName(name string) []DataDescriptionEntry {
	entries := []DataDescriptionEntry{}
	for _, v := range e.dataDescriptionEntriesTable {
		if v.Name() == util.Symbol(name) {
			entries = append(entries, v)
		}
	}
	return entries
}
