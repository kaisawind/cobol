package valuestmt

import (
	"github.com/kaisawind/cobol/asg/internal/element"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type RelationConditionValueStmt struct {
	valuestmt.ValueStmt
	ctx *cobol85.RelationConditionContext
}

func NewRelationConditionValueStmt(ctx *cobol85.RelationConditionContext, programUnit element.ProgramUnit) valuestmt.RelationConditionValueStmt {
	return &RelationConditionValueStmt{
		ValueStmt: NewValueStmt(ctx, programUnit),
		ctx:       ctx,
	}
}
