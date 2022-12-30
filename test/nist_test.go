package test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/document"
	"github.com/kaisawind/cobol/format"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/options"
)

type ErrorListener struct {
	antlr.DefaultErrorListener
	t *testing.T
}

func NewErrorListener(t *testing.T) *ErrorListener {
	return &ErrorListener{t: t}
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err := fmt.Sprintf("syntax error in line %d : %d %s", line, column, msg)
	l.t.Error(err)
	l.t.FailNow()
}

func TestParse(tt *testing.T) {
	rootdir := "./testdata/nist"
	infos, err := os.ReadDir(rootdir)
	if err != nil {
		tt.Error(err)
		tt.FailNow()
	}

	opts := options.NewOptions().AddCopyBookDirectory(rootdir).SetFormat(format.FIXED)
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		tt.Run(info.Name(), func(t *testing.T) {
			listener := NewErrorListener(t)

			filepath := path.Join(rootdir, info.Name())
			processed := document.ParseFile(filepath, opts)
			is := antlr.NewInputStream(processed)
			lexer := cobol85.NewCobol85Lexer(is)
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(listener)

			cts := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
			cpp := cobol85.NewCobol85Parser(cts)
			cpp.RemoveErrorListeners()
			cpp.AddErrorListener(listener)

			_ = antlr.TreesStringTree(cpp.StartRule(), []string{}, cpp)
		})
	}
}
