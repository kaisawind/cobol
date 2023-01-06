package internal

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ProgramUnitElement struct {
	model.CompilationUnitElement
	programUnit model.ProgramUnit
}

func NewProgramUnitElement(ctx antlr.ParserRuleContext, programUnit model.ProgramUnit) model.ProgramUnitElement {
	return &ProgramUnitElement{
		CompilationUnitElement: NewCompilationUnitElement(ctx, programUnit.CompilationUnit()),
		programUnit:            programUnit,
	}
}

func (e *ProgramUnitElement) ProgramUnit() model.ProgramUnit {
	return e.programUnit
}

func (e *ProgramUnitElement) GetElement(ctx antlr.Tree) model.Element {
	return e.Program().GetRegistry().GetElement(ctx)
}

func (e *ProgramUnitElement) AddElement(element model.Element) {
	e.Program().GetRegistry().AddElement(element)
}

func (e *ProgramUnitElement) CreateValueStmt(ctxs ...antlr.ParserRuleContext) (ret valuestmt.ValueStmt) {
	for _, ctx := range ctxs {
		if ret != nil {
			break
		}
		if ctx == nil {
			continue
		}

		switch t := ctx.(type) {
		case *cobol85.IdentifierContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.AlphabetNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.ArithmeticExpressionContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.AssignmentNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.ClassNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.CobolWordContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.ConditionContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.DataDescNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.DataNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.FileNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.IndexNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.IntegerLiteralContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.LibraryNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.LiteralContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.LocalNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.ProgramNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.QualifiedDataNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.RelationConditionContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.ReportNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.SystemNameContext:
			ret = e.CreateCallValueStmt(ctx)
		case *cobol85.TableCallContext:
			ret = e.CreateCallValueStmt(ctx)
		default:
		}
	}
	return
}

func (e *ProgramUnitElement) CreateCallValueStmt(ctx antlr.ParserRuleContext) (ret valuestmt.CallValueStmt) {
	switch t := ctx.(type) {
	case *cobol85.IdentifierContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.AlphabetNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.ArithmeticExpressionContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.AssignmentNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.ClassNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.CobolWordContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.ConditionContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.DataDescNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.DataNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.FileNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.IndexNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.IntegerLiteralContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.LibraryNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.LiteralContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.LocalNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.ProgramNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.QualifiedDataNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.RelationConditionContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.ReportNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.SystemNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.TableCallContext:
		ret = e.CreateCallValueStmt(ctx)
	default:
	}
	return
}

func (e *ProgramUnitElement) CreateCall(ctx antlr.ParserRuleContext) (ret call.Call) {
	switch t := ctx.(type) {
	case *cobol85.IdentifierContext:
		if t.QualifiedDataName() != nil {

		} else if t.TableCall() != nil {

		} else if t.FunctionCall() != nil {

		} else if t.SpecialRegister() != nil {

		} else {

		}
	case *cobol85.AlphabetNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.ArithmeticExpressionContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.AssignmentNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.ClassNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.CobolWordContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.ConditionContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.DataDescNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.DataNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.FileNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.IndexNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.IntegerLiteralContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.LibraryNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.LiteralContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.LocalNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.ProgramNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.QualifiedDataNameContext:
		if t.QualifiedDataNameFormat1() != nil {

		} else if t.QualifiedDataNameFormat2() != nil {

		} else if t.QualifiedDataNameFormat3() != nil {

		} else if t.QualifiedDataNameFormat4() != nil {

		} else {

		}
	case *cobol85.QualifiedDataNameFormat1Context:
	case *cobol85.QualifiedDataNameFormat2Context:
	case *cobol85.QualifiedDataNameFormat3Context:
	case *cobol85.QualifiedDataNameFormat4Context:
	case *cobol85.RelationConditionContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.ReportNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.SystemNameContext:
		ret = e.CreateCallValueStmt(ctx)
	case *cobol85.TableCallContext:
		ret = e.CreateCallValueStmt(ctx)
	default:
	}
	return
}
