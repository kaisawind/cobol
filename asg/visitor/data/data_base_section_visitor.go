package data

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type DataBaseSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.DataBaseSection
}

func NewDataBaseSectionVisitor(section *pb.DataBaseSection) *DataBaseSectionVisitor {
	return &DataBaseSectionVisitor{
		section: section,
	}
}

func (v *DataBaseSectionVisitor) VisitDataBaseSection(ctx *cobol85.DataBaseSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *DataBaseSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *DataBaseSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
