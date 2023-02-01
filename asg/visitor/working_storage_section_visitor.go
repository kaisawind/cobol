package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type WorkingStorageSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.WorkingStorageSection
}

func NewWorkingStorageSectionVisitor(section *pb.WorkingStorageSection) *WorkingStorageSectionVisitor {
	return &WorkingStorageSectionVisitor{
		section: section,
	}
}

func (v *WorkingStorageSectionVisitor) VisitWorkingStorageSection(ctx *cobol85.WorkingStorageSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *WorkingStorageSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *WorkingStorageSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
