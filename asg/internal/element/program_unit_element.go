package element

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/asg/model/data/workingstorage"
	"github.com/kaisawind/cobol/asg/model/element"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/asg/util"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ProgramUnitElement struct {
	element.CompilationUnitElement
	programUnit element.ProgramUnit
}

func NewProgramUnitElement(ctx antlr.ParserRuleContext, programUnit element.ProgramUnit) element.ProgramUnitElement {
	return &ProgramUnitElement{
		CompilationUnitElement: NewCompilationUnitElement(ctx, programUnit.CompilationUnit()),
		programUnit:            programUnit,
	}
}

func (e *ProgramUnitElement) ProgramUnit() element.ProgramUnit {
	return e.programUnit
}

func (e *ProgramUnitElement) GetElement(ctx antlr.Tree) element.Element {
	return e.Program().GetRegistry().GetElement(ctx)
}

func (e *ProgramUnitElement) AddElement(element element.Element) {
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
			dataCall := e.CreateCall(t.QualifiedDataNameFormat1())
			ret = NewCallDelegate(ctx, dataCall, e.ProgramUnit())
			e.AddElement(ret)
		} else if t.QualifiedDataNameFormat2() != nil {
			dataCall := e.CreateCall(t.QualifiedDataNameFormat2())
			ret = NewCallDelegate(ctx, dataCall, e.ProgramUnit())
			e.AddElement(ret)
		} else if t.QualifiedDataNameFormat3() != nil {
			dataCall := e.CreateCall(t.QualifiedDataNameFormat3())
			ret = NewCallDelegate(ctx, dataCall, e.ProgramUnit())
			e.AddElement(ret)
		} else if t.QualifiedDataNameFormat4() != nil {
			dataCall := e.CreateCall(t.QualifiedDataNameFormat4())
			ret = NewCallDelegate(ctx, dataCall, e.ProgramUnit())
			e.AddElement(ret)
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

func (e *ProgramUnitElement) CreateDataDescriptionEntryCall(ctx antlr.ParserRuleContext) (ret call.Call) {
	element := e.GetElement(ctx)
	if element != nil {
		ret = element.(call.Call)
	} else {
		name := util.DetermineName(ctx)
	}
	return
}

func (e *ProgramUnitElement) GetIndex(name string) datadescription.Index {

}

func (e *ProgramUnitElement) WorkingStorageSection() workingstorage.WorkingStorageSection {
	programUnit := e.ProgramUnit().(*ProgramUnit)
}
