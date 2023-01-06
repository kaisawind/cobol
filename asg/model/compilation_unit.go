package model

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type CompilationUnit interface {
	Element
	RegisterElement(element Element)
	Name() string
	Context() antlr.ParserRuleContext
}
