package document

import (
	"sort"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/preprocessor"
)

type Context struct {
	stores Stores
	buffer string
}

func NewContext() *Context {
	return &Context{
		stores: Stores{},
		buffer: "",
	}
}

func (ctx *Context) Store(ctxs []preprocessor.IReplaceClauseContext) {
	for _, v := range ctxs {
		rcc, ok := v.(*preprocessor.ReplaceClauseContext)
		if ok {
			ctx.stores = append(ctx.stores, NewStore(rcc.Replaceable(), rcc.Replacement()))
		}
	}
}

func (ctx *Context) Replace(cts *antlr.CommonTokenStream) {
	if len(ctx.stores) == 0 {
		return
	}
	sort.Sort(ctx.stores)
	for _, store := range ctx.stores {
		ctx.buffer = store.Replace(ctx.buffer, cts)
	}
}

func (ctx *Context) Read() string {
	return ctx.buffer
}

func (ctx *Context) Write(s string) {
	ctx.buffer += s
}
