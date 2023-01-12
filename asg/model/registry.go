package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Registry 元组件仓库
type Registry struct {
	elements map[antlr.Tree]Element
}

func NewRegistry() *Registry {
	return &Registry{
		elements: map[antlr.Tree]Element{},
	}
}

// AddElement 注册元组件
func (e *Registry) AddElement(element Element) {
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

// GetElement 获取元组件
func (e *Registry) GetElement(ctx antlr.Tree) Element {
	return e.elements[ctx]
}
