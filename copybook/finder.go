package copybook

import (
	"regexp"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/preprocessor"
	"github.com/kaisawind/cobol/options"
)

type Finder struct {
}

func NewFinder() *Finder {
	return &Finder{}
}

func TrimQuotes(s string) string {
	return regexp.MustCompile("^[\"']|[\"']$").ReplaceAllString(s, "")
}

func GetCopyBook(ctx antlr.ParserRuleContext, opts *options.Options) (ret string) {
	if cctx, ok := ctx.(*preprocessor.CopySourceContext); ok {
		if icctx := cctx.CobolWord(); icctx != nil {
			ret = NewCobolWordFinder().GetCopyBook(icctx, opts)
		}
		if icctx := cctx.Filename(); icctx != nil {
			ret = NewFilenameFinder().GetCopyBook(icctx, opts)
		}
		if icctx := cctx.Literal(); icctx != nil {
			ret = NewLiteralFinder().GetCopyBook(icctx, opts)
		}
	}
	if cctx, ok := ctx.(preprocessor.ICobolWordContext); ok {
		ret = NewCobolWordFinder().GetCopyBook(cctx, opts)
	}
	if cctx, ok := ctx.(preprocessor.IFilenameContext); ok {
		ret = NewFilenameFinder().GetCopyBook(cctx, opts)
	}
	if cctx, ok := ctx.(preprocessor.ILiteralContext); ok {
		ret = NewLiteralFinder().GetCopyBook(cctx, opts)
	}
	return
}
