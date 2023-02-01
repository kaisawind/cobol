package data

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type FileSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.FileSection
}

func NewFileSectionVisitor(section *pb.FileSection) *FileSectionVisitor {
	return &FileSectionVisitor{
		section: section,
	}
}

func (v *FileSectionVisitor) VisitFileSection(ctx *cobol85.FileSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *FileSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *FileSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
