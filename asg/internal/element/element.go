package element

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/element"
)

type Element struct {
	ctx     antlr.ParserRuleContext
	program element.Program
}

func NewElement(ctx antlr.ParserRuleContext, program element.Program) element.Element {
	return &Element{
		ctx:     ctx,
		program: program,
	}
}

func (e *Element) Parent() (parent element.Element) {
	registry := e.program.GetRegistry()
	current := e.ctx.(antlr.Tree)
	for {
		if parent != nil || current == nil {
			break
		}
		current = current.GetParent()
		element := registry.GetElement(current)
		if element != nil {
			parent = element
		}
	}
	return parent
}

func (e *Element) Children() (elements []element.Element) {
	registry := e.program.GetRegistry()
	for _, v := range e.ctx.GetChildren() {
		element := registry.GetElement(v)
		if element != nil {
			elements = append(elements, element)
		}
	}
	return
}

func (e *Element) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *Element) Program() element.Program {
	return e.program
}
