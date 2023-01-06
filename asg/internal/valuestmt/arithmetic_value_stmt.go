package valuestmt

import (
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ArithmeticValueStmt struct {
	valuestmt.ValueStmt
}

func NewArithmeticValueStmt(ctx *cobol85.ArgumentContext, programUnit model.ProgramUnit) valuestmt.ArithmeticValueStmt {
	return &ArithmeticValueStmt{
		ValueStmt: NewValueStmt(ctx, programUnit),
	}
}
