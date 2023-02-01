package data

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type ReportSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.ReportSection
}

func NewReportSectionVisitor(section *pb.ReportSection) *ReportSectionVisitor {
	return &ReportSectionVisitor{
		section: section,
	}
}

func (v *ReportSectionVisitor) VisitReportSection(ctx *cobol85.ReportSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *ReportSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ReportSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
