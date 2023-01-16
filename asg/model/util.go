package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func GetParent[T ElementType](program *Program, ctx antlr.Tree) (parent T) {
	return getParent[T](program.program, ctx)
}
