package valuestmt

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type IntegerLiteralValueStmt struct {
	valuestmt.ValueStmt
}

func NewIntegerLiteralValueStmt(ctx *cobol85.IntegerLiteralContext, programUnit model.ProgramUnit) valuestmt.IntegerLiteralValueStmt {
	return &IntegerLiteralValueStmt{
		ValueStmt: NewValueStmt(ctx, programUnit),
	}
}
