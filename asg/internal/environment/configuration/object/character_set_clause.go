package object

import (
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/object"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CharacterSetClause struct {
	element.CobolDivisionElement
	ctx *cobol85.CharacterSetClauseContext
}

func NewCharacterSetClause(ctx *cobol85.CharacterSetClauseContext, programUnit element.ProgramUnit) object.CharacterSetClause {
	return &CharacterSetClause{
		CobolDivisionElement: instances.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
	}
}
