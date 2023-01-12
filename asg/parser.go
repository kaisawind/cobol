package asg

import (
	"path"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/visitor"
	"github.com/kaisawind/cobol/document"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/options"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func AnalyzeFile(filename string, opts ...options.Option) *model.Program {
	if filename == "" {
		return nil
	}
	program := model.NewProgram()
	AnalyzeCompilationUnit(filename, program, opts...)
	AnalyzeProgramUnits(program)
	return program
}

func GetCompilationUnitName(filename string) string {
	return cases.Title(language.English).String(strings.TrimSuffix(path.Base(filename), path.Ext(filename)))
}

func AnalyzeCompilationUnit(filename string, program *model.Program, opts ...options.Option) {
	name := GetCompilationUnitName(filename)
	processed := document.ParseFile(filename, opts...)

	is := antlr.NewInputStream(processed)
	lexer := cobol85.NewCobol85Lexer(is)
	cts := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	cpp := cobol85.NewCobol85Parser(cts)

	ctx := cpp.StartRule()

	vr := visitor.NewCompilationUnitVisitor(name, program)

	vr.Visit(ctx)
}

func AnalyzeProgramUnits(program *model.Program) {
	for _, unit := range program.GetCompilationUnits() {
		vr := visitor.NewProgramUnitVisitor(unit, program)
		vr.Visit(unit.Context())
	}
}
