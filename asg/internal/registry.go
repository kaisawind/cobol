package internal

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
)

type Registry struct {
	elements map[antlr.Tree]model.Element
}

func NewRegistry() model.Registry {
	return &Registry{
		elements: map[antlr.Tree]model.Element{},
	}
}

func (e *Registry) AddElement(element model.Element) {
	if element == nil {
		return
	}
	ctx := element.Context()
	if ctx == nil {
		return
	}
	if pt, ok := ctx.(antlr.Tree); ok {
		if _, ok = e.elements[pt]; !ok {
			e.elements[pt] = element
		}
	}
}

func (e *Registry) GetElement(ctx antlr.Tree) model.Element {
	return e.elements[ctx]
}
