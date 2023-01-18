package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

func (e *program) CreateValueStmt(ctxs ...antlr.ParserRuleContext) (ret valuestmt.ValueStmt) {
	for _, ctx := range ctxs {
		if ret != nil {
			break
		}
		if ctx == nil {
			continue
		}
		switch t := ctx.(type) {
		case cobol85.IIdentifierContext:
			ret = e.CreateCallValueStmt(t)
		case cobol85.IQualifiedDataNameContext:
			ret = e.CreateCallValueStmt(t)
		case cobol85.IIndexNameContext:
			ret = e.CreateCallValueStmt(t)
		case cobol85.IArithmeticExpressionContext:
			// TODO:
		}
	}
	return
}

func (e *program) CreateCallValueStmt(ctx antlr.ParserRuleContext) *valuestmt.CallValueStmt {
	delegatedCall := e.CreateCall(ctx)
	return valuestmt.NewCallValueStmt(ctx, delegatedCall)
}
