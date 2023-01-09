package valuestmt

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/instances"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
)

type ValueStmt struct {
	element.CobolDivisionElement
	ctx        antlr.ParserRuleContext
	valueStmts []valuestmt.ValueStmt
}

func NewValueStmt(ctx antlr.ParserRuleContext, programUnit element.ProgramUnit) valuestmt.ValueStmt {
	return &ValueStmt{
		CobolDivisionElement: instances.NewCobolDivisionElement(ctx, programUnit),
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
