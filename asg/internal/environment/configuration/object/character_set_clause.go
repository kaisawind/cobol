package object

import (
	"github.com/kaisawind/cobol/asg/internal"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/environment/configuration/object"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CharacterSetClause struct {
	model.CobolDivisionElement
	ctx *cobol85.CharacterSetClauseContext
}

func NewCharacterSetClause(ctx *cobol85.CharacterSetClauseContext, programUnit model.ProgramUnit) object.CharacterSetClause {
	return &CharacterSetClause{
		CobolDivisionElement: internal.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
	}
}
