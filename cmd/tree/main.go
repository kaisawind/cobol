package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/conv"
	"github.com/kaisawind/cobol/document"
	"github.com/kaisawind/cobol/format"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/options"
)

var (
	formatFlag = flag.String("format", "FIXED", "cobol file format, FIXED, TANDEM, VARIABLE")
	pathFlag   = flag.String("path", ".", "cobol files path")
	suffixFlag = flag.String("suffix", ".CBL", "cobol cbl")
)

func main() {
	flag.Parse()

	f := format.FIXED
	switch strings.ToUpper(*formatFlag) {
	case "FIXED":
		f = format.FIXED
	case "TANDEM":
		f = format.TANDEM
	case "VARIABLE":
		f = format.VARIABLE
	}
	rootPath := *pathFlag
	fi, err := os.Stat(rootPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		return
	}
	if !fi.IsDir() {
		TreesStringTree(rootPath, f)
	} else {
		filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				TreesStringTree(path, f)
			}
			return err
		})
	}
}

func TreesStringTree(path string, f format.Format) {
	if strings.HasSuffix(path, ".tree") {
		return
	}
	suffix := *suffixFlag
	if suffix != "" && !strings.HasSuffix(path, suffix) {
		return
	}
	fmt.Fprintln(os.Stdout, path)
	buff, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		return
	}

	opts := options.NewOptions().SetFormat(f)
	processed := document.Parse(string(buff), opts)

	is := antlr.NewInputStream(processed)
	lexer := cobol85.NewCobol85Lexer(is)
	cts := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	cpp := cobol85.NewCobol85Parser(cts)

	l := NewErrorListener()
	lexer.AddErrorListener(l)

	errs := l.GetErrors()
	if len(errs) != 0 {
		os.WriteFile(path+".error", []byte(strings.Join(errs, "\n")), os.ModePerm)
	}

	ctx := cpp.StartRule()

	tree := conv.TreesStringTree(ctx, cpp.GetRuleNames(), 0)
	os.WriteFile(path+".tree", []byte(tree), os.ModePerm)
}

type ErrorListener struct {
	antlr.DefaultErrorListener
	errs []string
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		errs: []string{},
	}
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err := fmt.Sprintf("syntax error in line %d : %d %s", line, column, msg)
	l.errs = append(l.errs, err)
}

func (l *ErrorListener) GetErrors() []string {
	return l.errs
}
