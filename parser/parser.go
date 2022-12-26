package parser

import (
	"bufio"
	"fmt"
	"io"
	"regexp"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/constant"
	"github.com/kaisawind/cobol/gen/preprocessor"
)

func Parser() {
	_, err := antlr.NewFileStream("./testdata/lbli0420.src")
	if err != nil {
		return
	}
}

func ReadLines(r io.Reader) {
	scan := bufio.NewScanner(r)
	scan.Scan()
	scan.Text()
}

var (
	executionReg = regexp.MustCompile(
		fmt.Sprintf(
			`.*(%s|%s|%s|%s|%s|%s|%s|%s|%s|%s\s+%s|%s\s+%s|%s\s+%s).*`,
			constant.CBL,
			constant.COPY,
			constant.PROCESS,
			constant.REPLACE,
			constant.EJECT,
			constant.SKIP1,
			constant.SKIP2,
			constant.SKIP3,
			constant.TITLE,
			constant.EXEC, constant.SQL,
			constant.EXEC, constant.SQLIMS,
			constant.EXEC, constant.CICS,
		))
)

func ExecutionParser(code string) string {
	if executionReg.MatchString(code) {
		is := antlr.NewInputStream(code)
		lexer := preprocessor.NewCobol85PreprocessorLexer(is)
		cts := antlr.NewCommonTokenStream(lexer, 0)
		cpp := preprocessor.NewCobol85PreprocessorParser(cts)
		listener := NewPreprocessorListener(cts)
		antlr.ParseTreeWalkerDefault.Walk(listener, cpp.StartRule())
	}
	return code
}
