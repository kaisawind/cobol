package valuestmt

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/internal"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
)

type ValueStmt struct {
	model.CobolDivisionElement
	ctx        antlr.ParserRuleContext
	valueStmts []valuestmt.ValueStmt
}

func NewValueStmt(ctx antlr.ParserRuleContext, programUnit model.ProgramUnit) valuestmt.ValueStmt {
	return &ValueStmt{
		CobolDivisionElement: internal.NewCobolDivisionElement(ctx, programUnit),
		ctx:                  ctx,
		valueStmts:           []valuestmt.ValueStmt{},
	}
}

func (e *ValueStmt) AddValueStmt(value valuestmt.ValueStmt) {
	e.valueStmts = append(e.valueStmts, value)
}

func (e *ValueStmt) GetValueStmts() []valuestmt.ValueStmt {
	return e.valueStmts
}
