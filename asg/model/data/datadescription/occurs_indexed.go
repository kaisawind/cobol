package datadescription

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type OccursIndexed struct {
	ctx          cobol85.IDataOccursIndexedContext
	indexes      []*Index
	indexesTable map[string]*Index
}

func NewOccursIndexed(ctx cobol85.IDataOccursIndexedContext) *OccursIndexed {
	return &OccursIndexed{
		ctx:          ctx,
		indexes:      []*Index{},
		indexesTable: map[string]*Index{},
	}
}

func (e *OccursIndexed) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *OccursIndexed) AddIndex(name string, index *Index) {
	e.indexes = append(e.indexes, index)
	e.indexesTable[name] = index
}

func (e *OccursIndexed) GetIndex(name string) *Index {
	return e.indexesTable[name]
}

func (e *OccursIndexed) GetIndexes() []*Index {
	return e.indexes
}
