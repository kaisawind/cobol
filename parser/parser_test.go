package parser

import (
	"fmt"
	"os"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/format"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/line"
)

type TreeShapeListener struct {
	*cobol85.BaseCobol85Listener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func TestParser(t *testing.T) {
	fs, err := antlr.NewFileStream("./testdata/example.cbl")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	lexer := cobol85.NewCobol85Lexer(fs)
	ts := antlr.NewCommonTokenStream(lexer, 0)
	p := cobol85.NewCobol85Parser(ts)
	tree := p.StartRule()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

func TestExecutionParser(t *testing.T) {
	f, err := os.Open("./testdata/CopyReplace.cbl")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer f.Close()
	code := line.CombineLinkedLine(line.NewLinkedLine(f, format.FIXED))
	code = ExecutionParser(code)
	t.Log("\n" + code)
}
