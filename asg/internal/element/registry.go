package element

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/element"
)

type Registry struct {
	elements map[antlr.Tree]element.Element
}

func NewRegistry() element.Registry {
	return &Registry{
		elements: map[antlr.Tree]element.Element{},
	}
}

func (e *Registry) AddElement(element element.Element) {
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

func (e *Registry) GetElement(ctx antlr.Tree) element.Element {
	return e.elements[ctx]
}
