package model

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type CompilationUnit interface {
	Element
	RegistryElement(element Element)
	Name() string
	Context() antlr.ParserRuleContext
}
