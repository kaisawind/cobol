package valuestmt

import (
	"github.com/kaisawind/cobol/asg/internal/element"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type LiteralValueStmt struct {
	valuestmt.ValueStmt
}

func NewLiteralValueStmt(ctx *cobol85.LiteralContext, programUnit element.ProgramUnit) valuestmt.LiteralValueStmt {
	return &LiteralValueStmt{
		ValueStmt: NewValueStmt(ctx, programUnit),
	}
}
