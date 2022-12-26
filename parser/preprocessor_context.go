package parser

import (
	"sort"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/preprocessor"
)

type PreprocessorContext struct {
	stores Stores
	buffer string
}

func NewPreprocessorContext() *PreprocessorContext {
	return &PreprocessorContext{
		stores: Stores{},
		buffer: "",
	}
}

func (ctx *PreprocessorContext) Store(ctxs []preprocessor.IReplaceClauseContext) {
	for _, v := range ctxs {
		rcc, ok := v.(*preprocessor.ReplaceClauseContext)
		if ok {
			ctx.stores = append(ctx.stores, NewStore(rcc.Replaceable(), rcc.Replacement()))
		}
	}
}

func (ctx *PreprocessorContext) Replace(cts *antlr.CommonTokenStream) {
	if len(ctx.stores) == 0 {
		return
	}
	sort.Sort(ctx.stores)
	for _, store := range ctx.stores {
		ctx.buffer = store.Replace(ctx.buffer, cts)
	}
}
