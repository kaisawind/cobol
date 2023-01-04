package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Element interface {
	Parent() Element
	Children() []Element
	Context() antlr.ParserRuleContext
	Program() *Program
}

type BaseElement struct {
	ctx     antlr.ParserRuleContext
	program *Program
}

func NewBaseElement(ctx antlr.ParserRuleContext, program *Program) *BaseElement {
	return &BaseElement{
		ctx:     ctx,
		program: program,
	}
}

func (e *BaseElement) Parent() (parent Element) {
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

func (e *BaseElement) Children() (elements []Element) {
	registry := e.program.GetRegistry()
	for _, v := range e.ctx.GetChildren() {
		element := registry.GetElement(v)
		if element != nil {
			elements = append(elements, element)
		}
	}
	return
}

func (e *BaseElement) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *BaseElement) Program() *Program {
	return e.program
}
