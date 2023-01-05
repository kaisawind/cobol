package model

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type Registry interface {
	AddElement(element Element)
	GetElement(ctx antlr.Tree) Element
}
