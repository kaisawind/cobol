package impl

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
)

type Element struct {
	ctx     antlr.ParserRuleContext
	program model.Program
}

func NewElement(ctx antlr.ParserRuleContext, program model.Program) model.Element {
	return &Element{
		ctx:     ctx,
		program: program,
	}
}

func (e *Element) Parent() (parent model.Element) {
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

func (e *Element) Children() (elements []model.Element) {
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

func (e *Element) Program() model.Program {
	return e.program
}
