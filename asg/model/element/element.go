package element

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type Element interface {
	Parent() Element
	Children() []Element
	Context() antlr.ParserRuleContext
	Program() Program
}
