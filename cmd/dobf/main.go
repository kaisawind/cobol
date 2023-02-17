package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/document"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/options"
)

func main() {
	buff := make([]byte, 1024*1024*2)
	n, err := os.Stdin.Read(buff)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}
	processed := string(buff[:n])
	is := antlr.NewInputStream(processed)
	lexer := cobol85.NewCobol85Lexer(is)
	cts := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	cpp := cobol85.NewCobol85Parser(cts)

	listener := NewDobfListener(cts)
	antlr.ParseTreeWalkerDefault.Walk(listener, cpp.StartRule())
	output := map[string]any{
		"vars":   listener.GetVars(),
		"tokens": listener.GetTokens(),
	}
	buff, err = json.Marshal(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "%s", string(buff))
}

type Tok struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Context struct {
	buffer string
	tokens []*Tok
}

func NewContext() *Context {
	return &Context{
		buffer: "",
	}
}

func (ctx *Context) Read() string {
	return ctx.buffer
}

func (ctx *Context) Write(s string) {
	ctx.buffer += s
}

func (ctx *Context) GetTokens() []*Tok {
	return ctx.tokens
}

func (ctx *Context) AppendTokens(tokens []*Tok) {
	ctx.tokens = append(ctx.tokens, tokens...)
}

type DobfListener struct {
	cobol85.BaseCobol85Listener
	contexts []*Context
	cts      *antlr.CommonTokenStream
	opts     *options.Options
	vars     map[string]int
	index    int
}

func NewDobfListener(cts *antlr.CommonTokenStream, opts ...options.Option) *DobfListener {
	o := options.NewOptions()
	for _, opt := range opts {
		opt.Apply(o)
	}
	return &DobfListener{
		contexts: []*Context{NewContext()},
		opts:     o,
		cts:      cts,
		vars:     map[string]int{},
	}
}

func (s *DobfListener) push() {
	s.contexts = append([]*Context{NewContext()}, s.contexts...)
}

func (s *DobfListener) pop() {
	if len(s.contexts) > 0 {
		s.contexts = s.contexts[1:]
	}
}

func (s *DobfListener) peek() *Context {
	if len(s.contexts) == 0 {
		return nil
	}
	return s.contexts[0]
}

func (s *DobfListener) Context() *Context {
	return s.peek()
}

func (s *DobfListener) GetText() string {
	return s.Context().Read()
}

func (s *DobfListener) GetVars() (out map[string]string) {
	out = map[string]string{}
	for k, v := range s.vars {
		out[fmt.Sprintf("VAR_%d", v)] = k
	}
	return
}

func (s *DobfListener) GetTokens() []*Tok {
	return s.Context().GetTokens()
}

// EnterDataName is called when production dataName is entered.
func (s *DobfListener) EnterDataName(ctx *cobol85.DataNameContext) {
	s.push()
}

// ExitDataName is called when production dataName is exited.
func (s *DobfListener) ExitDataName(ctx *cobol85.DataNameContext) {
	popped := s.Context().Read()
	s.pop()
	tree := ctx.GetParent()
	oldVariable := ctx.GetText()
	newVariable := ""
	switch tree.(type) {
	case cobol85.IDataDescriptionEntryFormat1Context:
		i, ok := s.vars[oldVariable]
		if !ok {
			s.index++
			s.vars[oldVariable] = s.index
			newVariable = fmt.Sprintf("VAR_%d", s.index)
		} else {
			newVariable = fmt.Sprintf("VAR_%d", i)
		}
	case cobol85.IDataDescriptionEntryFormat2Context:
		i, ok := s.vars[oldVariable]
		if !ok {
			s.index++
			s.vars[oldVariable] = s.index
			newVariable = fmt.Sprintf("VAR_%d", s.index)
		} else {
			newVariable = fmt.Sprintf("VAR_%d", i)
		}
	case cobol85.IQualifiedDataNameFormat1Context:
		i, ok := s.vars[oldVariable]
		if ok {
			newVariable = fmt.Sprintf("VAR_%d", i)
		} else {
			newVariable = oldVariable
		}
	default:
		newVariable = oldVariable
	}

	content := strings.ReplaceAll(popped, oldVariable, newVariable)
	s.Context().AppendTokens([]*Tok{
		{
			Type:  "RULE",
			Value: newVariable,
		},
	})
	s.Context().Write(content)
}

// VisitTerminal is called when a terminal node is visited.
func (s *DobfListener) VisitTerminal(node antlr.TerminalNode) {
	pos := node.GetSourceInterval().Start
	tok := node.GetSymbol()
	tt := tok.GetTokenType()
	if tt != cobol85.Cobol85ParserCOMMENTENTRYLINE &&
		tt != cobol85.Cobol85ParserCOMMENTLINE {
		text := strings.TrimSpace(tok.GetText())
		if tt == cobol85.Cobol85ParserDOT_FS ||
			tt == cobol85.Cobol85ParserDOT {
			s.Context().AppendTokens([]*Tok{
				{
					Type:  "DOT",
					Value: text,
				},
				{
					Type:  "NEWLINE",
					Value: "\n",
				},
			})
		} else if tt == cobol85.Cobol85ParserNONNUMERICLITERAL {
			s.Context().AppendTokens([]*Tok{
				{
					Type:  "NONNUMERICLITERAL",
					Value: text,
				},
			})
		} else {
			s.Context().AppendTokens([]*Tok{
				{
					Type:  "RULE",
					Value: text,
				},
			})
		}
	}

	left := document.GetHiddenTokensToLeft(s.cts, pos)
	s.Context().Write(left)
	if node.GetSymbol().GetTokenType() != antlr.TokenEOF {
		s.Context().Write(node.GetText())
	}
}
