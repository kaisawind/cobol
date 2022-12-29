package document

import (
	"bufio"
	"regexp"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/constant"
	"github.com/kaisawind/cobol/copybook"
	"github.com/kaisawind/cobol/gen/preprocessor"
	"github.com/kaisawind/cobol/line"
	"github.com/kaisawind/cobol/options"
)

type PreprocessorListener struct {
	*preprocessor.BaseCobol85PreprocessorListener

	contexts []*Context
	cts      *antlr.CommonTokenStream
	opts     *options.Options
}

func NewPreprocessorListener(cts *antlr.CommonTokenStream, opts ...options.Option) *PreprocessorListener {
	o := options.NewOptions()
	for _, opt := range opts {
		opt.Apply(o)
	}
	return &PreprocessorListener{
		contexts: []*Context{NewContext()},
		cts:      cts,
		opts:     o,
	}
}

func (s *PreprocessorListener) push() {
	s.contexts = append([]*Context{NewContext()}, s.contexts...)
}

func (s *PreprocessorListener) pop() {
	if len(s.contexts) > 0 {
		s.contexts = s.contexts[1:]
	}
}

func (s *PreprocessorListener) peek() *Context {
	if len(s.contexts) == 0 {
		return nil
	}
	return s.contexts[0]
}

func (s *PreprocessorListener) context() *Context {
	return s.peek()
}

func (s *PreprocessorListener) getCopyBook(ctx preprocessor.ICopySourceContext, opts *options.Options) string {
	filename := copybook.GetCopyBook(ctx, opts)
	return ParseFile(filename, opts)
}

func (s *PreprocessorListener) buildLines(prefix, text string) (ret string) {
	scan := bufio.NewScanner(strings.NewReader(text))
	first := true
	re := regexp.MustCompile("(?i)(end-exec)")
	for scan.Scan() {
		if !first {
			ret += constant.CHAR_NEWLINE
		}
		line := scan.Text()
		line = strings.TrimSpace(line)
		line = prefix + constant.CHAR_WHITESPACE + line
		line = re.ReplaceAllString(line, "$1 "+constant.EXEC_END_TAG)
		ret += line
		first = false
	}
	return
}

// EnterCompilerOptions is called when production compilerOptions is entered.
func (s *PreprocessorListener) EnterCompilerOptions(ctx *preprocessor.CompilerOptionsContext) {
	s.push()
}

// ExitCompilerOptions is called when production compilerOptions is exited.
func (s *PreprocessorListener) ExitCompilerOptions(ctx *preprocessor.CompilerOptionContext) {
	s.pop()
}

// EnterCopyStatement is called when production copyStatement is entered.
func (s *PreprocessorListener) EnterCopyStatement(ctx *preprocessor.CopyStatementContext) {
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
	content := s.getCopyBook(ctx.CopySource(), s.opts)
	if content != "" {
		s.context().Write(content + constant.CHAR_NEWLINE)
		s.context().Replace(s.cts)
	}

	content = s.context().Read()
	s.pop()
	s.context().Write(content)
}

// EnterEjectStatement is called when production ejectStatement is entered.
func (s *PreprocessorListener) EnterEjectStatement(ctx *preprocessor.EjectStatementContext) {
	s.push()
}

// ExitEjectStatement is called when production ejectStatement is exited.
func (s *PreprocessorListener) ExitEjectStatement(ctx *preprocessor.EjectStatementContext) {
	s.pop()
}

// EnterExecCicsStatement is called when production execCicsStatement is entered.
func (s *PreprocessorListener) EnterExecCicsStatement(ctx *preprocessor.ExecCicsStatementContext) {
	s.push()
}

// ExitExecCicsStatement is called when production execCicsStatement is exited.
func (s *PreprocessorListener) ExitExecCicsStatement(ctx *preprocessor.ExecCicsStatementContext) {
	s.pop()
	s.push()

	text := GetTextWithHiddenTokens(ctx, s.cts)
	linePrefix := line.LinePrefix(s.opts.Format) + constant.EXEC_CICS_TAG
	content := s.buildLines(linePrefix, text)
	s.context().Write(content)

	content = s.context().Read()
	s.pop()
	s.context().Write(content)
}

// EnterExecSqlImsStatement is called when production execSqlImsStatement is entered.
func (s *PreprocessorListener) EnterExecSqlImsStatement(ctx *preprocessor.ExecSqlImsStatementContext) {
	s.push()
}

// ExitExecSqlImsStatement is called when production execSqlImsStatement is exited.
func (s *PreprocessorListener) ExitExecSqlImsStatement(ctx *preprocessor.ExecSqlImsStatementContext) {
	s.pop()
	s.push()

	text := GetTextWithHiddenTokens(ctx, s.cts)
	linePrefix := line.LinePrefix(s.opts.Format) + constant.EXEC_SQLIMS_TAG
	content := s.buildLines(linePrefix, text)
	s.context().Write(content)

	content = s.context().Read()
	s.pop()
	s.context().Write(content)

}

// EnterExecSqlStatement is called when production execSqlStatement is entered.
func (s *PreprocessorListener) EnterExecSqlStatement(ctx *preprocessor.ExecSqlStatementContext) {
	s.push()
}

// ExitExecSqlStatement is called when production execSqlStatement is exited.
func (s *PreprocessorListener) ExitExecSqlStatement(ctx *preprocessor.ExecSqlStatementContext) {
	s.pop()
	s.push()

	text := GetTextWithHiddenTokens(ctx, s.cts)
	linePrefix := line.LinePrefix(s.opts.Format) + constant.EXEC_SQL_TAG
	content := s.buildLines(linePrefix, text)
	s.context().Write(content)

	content = s.context().Read()
	s.pop()
	s.context().Write(content)
}

// EnterReplaceArea is called when production replaceArea is entered.
func (s *PreprocessorListener) EnterReplaceArea(ctx *preprocessor.ReplaceAreaContext) {
	s.push()
}

// ExitReplaceArea is called when production replaceArea is exited.
func (s *PreprocessorListener) ExitReplaceArea(ctx *preprocessor.ReplaceAreaContext) {
	if cctx, ok := ctx.ReplaceByStatement().(*preprocessor.ReplaceByStatementContext); ok {
		s.context().Store(cctx.AllReplaceClause())
		s.context().Replace(s.cts)

		content := s.context().Read()
		s.pop()
		s.context().Write(content)
	}
}

// EnterReplaceByStatement is called when production replaceByStatement is entered.
func (s *PreprocessorListener) EnterReplaceByStatement(ctx *preprocessor.ReplaceByStatementContext) {
	s.push()
}

// ExitReplaceByStatement is called when production replaceByStatement is exited.
func (s *PreprocessorListener) ExitReplaceByStatement(ctx *preprocessor.ReplaceByStatementContext) {
	s.pop()
}

// EnterReplaceOffStatement is called when production replaceOffStatement is entered.
func (s *PreprocessorListener) EnterReplaceOffStatement(ctx *preprocessor.ReplaceOffStatementContext) {
	s.push()
}

// ExitReplaceOffStatement is called when production replaceOffStatement is exited.
func (s *PreprocessorListener) ExitReplaceOffStatement(ctx *preprocessor.ReplaceOffStatementContext) {
	s.pop()
}

// EnterSkipStatement is called when production skipStatement is entered.
func (s *PreprocessorListener) EnterSkipStatement(ctx *preprocessor.SkipStatementContext) {
	s.push()
}

// ExitSkipStatement is called when production skipStatement is exited.
func (s *PreprocessorListener) ExitSkipStatement(ctx *preprocessor.SkipStatementContext) {
	s.pop()
}

// EnterTitleStatement is called when production titleStatement is entered.
func (s *PreprocessorListener) EnterTitleStatement(ctx *preprocessor.TitleStatementContext) {
	s.push()
}

// ExitTitleStatement is called when production titleStatement is exited.
func (s *PreprocessorListener) ExitTitleStatement(ctx *preprocessor.TitleStatementContext) {
	s.pop()
}

// VisitTerminal is called when a terminal node is visited.
func (s *PreprocessorListener) VisitTerminal(node antlr.TerminalNode) {
	pos := node.GetSourceInterval().Start
	s.context().Write(GetHiddenTokensToLeft(s.cts, pos))

	if node.GetSymbol().GetTokenType() != antlr.TokenEOF {
		s.context().Write(node.GetText())
	}
}
