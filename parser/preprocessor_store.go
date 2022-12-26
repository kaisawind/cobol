package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/preprocessor"
)

type Store struct {
	replaceable preprocessor.IReplaceableContext
	replacement preprocessor.IReplacementContext
}

func (s *Store) Replace(cts *antlr.CommonTokenStream) {

}

func (s *Store) getReplaceableText() {

}

type Stores []*Store

func (ss Stores) Len() int {
	return len(ss)
}

func (ss Stores) Less(i, j int) bool {
	iText := ss[i].replaceable.GetText()
	jText := ss[j].replaceable.GetText()
	return len(iText) < len(jText)
}

func (ss Stores) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
