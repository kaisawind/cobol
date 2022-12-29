package parser

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/preprocessor"
)

type PreprocessorListener struct {
	*preprocessor.BaseCobol85PreprocessorListener

	contexts []*PreprocessorContext
	cts      *antlr.CommonTokenStream
}

func NewPreprocessorListener(cts *antlr.CommonTokenStream) *PreprocessorListener {
	return &PreprocessorListener{
		contexts: []*PreprocessorContext{},
		cts:      cts,
	}
}

func (s *PreprocessorListener) push() {
	s.contexts = append([]*PreprocessorContext{NewPreprocessorContext()}, s.contexts...)
}

func (s *PreprocessorListener) pop() {
	if len(s.contexts) > 0 {
		s.contexts = s.contexts[1:]
	}
}

func (s *PreprocessorListener) peek() *PreprocessorContext {
	if len(s.contexts) == 0 {
		return nil
	}
	return s.contexts[0]
}

func (s *PreprocessorListener) context() *PreprocessorContext {
	return s.peek()
}

// EnterCompilerOptions is called when production compilerOptions is entered.
func (s *PreprocessorListener) EnterCompilerOptions(ctx *preprocessor.CompilerOptionsContext) {
	fmt.Println("EnterCompilerOptions", ctx)
}

// EnterCompilerOption is called when production compilerOption is entered.
func (s *PreprocessorListener) EnterCompilerOption(ctx *preprocessor.CompilerOptionContext) {}

// EnterCopyStatement is called when production copyStatement is entered.
func (s *PreprocessorListener) EnterCopyStatement(ctx *preprocessor.CopyStatementContext) {
	fmt.Println("EnterCopyStatement", ctx)
	s.push()
}

// ExitCopyStatement is called when production copyStatement is exited.
func (s *PreprocessorListener) ExitCopyStatement(ctx *preprocessor.CopyStatementContext) {
	s.pop()

	s.push()

	// replace phrase
	for _, iPhrase := range ctx.AllReplacingPhrase() {
		phrase, ok := iPhrase.(*preprocessor.ReplacingPhraseContext)
		if ok {
			s.context().Store(phrase.AllReplaceClause())
		}
	}

	// copy book
	copySource := ctx.CopySource()
}

// EnterEjectStatement is called when production ejectStatement is entered.
func (s *PreprocessorListener) EnterEjectStatement(ctx *preprocessor.EjectStatementContext) {}

// ExitEjectStatement is called when production ejectStatement is exited.
func (s *PreprocessorListener) ExitEjectStatement(ctx *preprocessor.EjectStatementContext) {}

// EnterExecCicsStatement is called when production execCicsStatement is entered.
func (s *PreprocessorListener) EnterExecCicsStatement(ctx *preprocessor.ExecCicsStatementContext) {}

// ExitExecCicsStatement is called when production execCicsStatement is exited.
func (s *PreprocessorListener) ExitExecCicsStatement(ctx *preprocessor.ExecCicsStatementContext) {}

// EnterExecSqlImsStatement is called when production execSqlImsStatement is entered.
func (s *PreprocessorListener) EnterExecSqlImsStatement(ctx *preprocessor.ExecSqlImsStatementContext) {
}

// ExitExecSqlImsStatement is called when production execSqlImsStatement is exited.
func (s *PreprocessorListener) ExitExecSqlImsStatement(ctx *preprocessor.ExecSqlImsStatementContext) {
}

// EnterExecSqlStatement is called when production execSqlStatement is entered.
func (s *PreprocessorListener) EnterExecSqlStatement(ctx *preprocessor.ExecSqlStatementContext) {}

// ExitExecSqlStatement is called when production execSqlStatement is exited.
func (s *PreprocessorListener) ExitExecSqlStatement(ctx *preprocessor.ExecSqlStatementContext) {}

// EnterReplaceArea is called when production replaceArea is entered.
func (s *PreprocessorListener) EnterReplaceArea(ctx *preprocessor.ReplaceAreaContext) {}

// ExitReplaceArea is called when production replaceArea is exited.
func (s *PreprocessorListener) ExitReplaceArea(ctx *preprocessor.ReplaceAreaContext) {}

// EnterReplaceByStatement is called when production replaceByStatement is entered.
func (s *PreprocessorListener) EnterReplaceByStatement(ctx *preprocessor.ReplaceByStatementContext) {
	fmt.Println("EnterReplaceByStatement", ctx)
}

// ExitReplaceByStatement is called when production replaceByStatement is exited.
func (s *PreprocessorListener) ExitReplaceByStatement(ctx *preprocessor.ReplaceByStatementContext) {
	fmt.Println("ExitReplaceByStatement", ctx)
}

// EnterReplaceOffStatement is called when production replaceOffStatement is entered.
func (s *PreprocessorListener) EnterReplaceOffStatement(ctx *preprocessor.ReplaceOffStatementContext) {
}

// ExitReplaceOffStatement is called when production replaceOffStatement is exited.
func (s *PreprocessorListener) ExitReplaceOffStatement(ctx *preprocessor.ReplaceOffStatementContext) {
}

// EnterSkipStatement is called when production skipStatement is entered.
func (s *PreprocessorListener) EnterSkipStatement(ctx *preprocessor.SkipStatementContext) {}

// ExitSkipStatement is called when production skipStatement is exited.
func (s *PreprocessorListener) ExitSkipStatement(ctx *preprocessor.SkipStatementContext) {}

// EnterTitleStatement is called when production titleStatement is entered.
func (s *PreprocessorListener) EnterTitleStatement(ctx *preprocessor.TitleStatementContext) {}

// ExitTitleStatement is called when production titleStatement is exited.
func (s *PreprocessorListener) ExitTitleStatement(ctx *preprocessor.TitleStatementContext) {}
