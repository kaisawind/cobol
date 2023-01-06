package valuestmt

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type RelationConditionValueStmt struct {
	valuestmt.ValueStmt
	ctx *cobol85.RelationConditionContext
}

func NewRelationConditionValueStmt(ctx *cobol85.RelationConditionContext, programUnit model.ProgramUnit) valuestmt.RelationConditionValueStmt {
	return &RelationConditionValueStmt{
		ValueStmt: NewValueStmt(ctx, programUnit),
		ctx:       ctx,
	}
}
