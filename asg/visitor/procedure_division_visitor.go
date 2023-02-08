package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/visitor/procedure"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type ProcedureDivisionVisitor struct {
	cobol85.BaseCobol85Visitor
	division *pb.ProcedureDivision
}

func NewProcedureDivisionVisitor(division *pb.ProcedureDivision) *ProcedureDivisionVisitor {
	return &ProcedureDivisionVisitor{
		division: division,
	}
}

func (v *ProcedureDivisionVisitor) VisitProcedureDeclaratives(ctx *cobol85.ProcedureDeclarativesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ProcedureDivisionVisitor) VisitProcedureDivisionGivingClause(ctx *cobol85.ProcedureDivisionGivingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ProcedureDivisionVisitor) VisitProcedureDivisionUsingClause(ctx *cobol85.ProcedureDivisionUsingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ProcedureDivisionVisitor) VisitParagraphs(ctx *cobol85.ParagraphsContext) interface{} {
	v.division.Paragraphs = &pb.Paragraphs{}
	vr := procedure.NewParagraphsVisitor(v.division.Paragraphs)
	return vr.Visit(ctx)
}

func (v *ProcedureDivisionVisitor) VisitProcedureSection(ctx *cobol85.ProcedureSectionContext) interface{} {
	section := &pb.ProcedureSection{}
	v.division.ProcedureSections = append(v.division.ProcedureSections, section)
	vr := procedure.NewProcedureSectionVisitor(section)
	return vr.Visit(ctx)
}

func (v *ProcedureDivisionVisitor) VisitProcedureDivisionBody(ctx *cobol85.ProcedureDivisionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ProcedureDivisionVisitor) VisitProcedureDivision(ctx *cobol85.ProcedureDivisionContext) any {
	return v.VisitChildren(ctx)
}

func (v *ProcedureDivisionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ProcedureDivisionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
