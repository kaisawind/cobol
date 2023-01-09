package valuestmt

import (
	"github.com/kaisawind/cobol/asg/internal/element"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ArithmeticValueStmt struct {
	valuestmt.ValueStmt
}

func NewArithmeticValueStmt(ctx *cobol85.ArgumentContext, programUnit element.ProgramUnit) valuestmt.ArithmeticValueStmt {
	return &ArithmeticValueStmt{
		ValueStmt: NewValueStmt(ctx, programUnit),
	}
}
