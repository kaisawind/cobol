package valuestmt

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ConditionValueStmt struct {
	valuestmt.ValueStmt
}

func NewConditionValueStmt(ctx *cobol85.ConditionContext, programUnit model.ProgramUnit) valuestmt.ConditionValueStmt {
	return &ConditionValueStmt{
		ValueStmt: NewValueStmt(ctx, programUnit),
	}
}
