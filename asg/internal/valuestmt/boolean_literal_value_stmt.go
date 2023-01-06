package valuestmt

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type BooleanLiteralValueStmt struct {
	valuestmt.ValueStmt
}

func NewBooleanLiteralValueStmt(ctx *cobol85.BooleanLiteralContext, programUnit model.ProgramUnit) valuestmt.BooleanLiteralValueStmt {
	return &BooleanLiteralValueStmt{
		ValueStmt: NewValueStmt(ctx, programUnit),
	}
}
