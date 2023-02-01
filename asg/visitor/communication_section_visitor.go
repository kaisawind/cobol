package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type CommunicationSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.CommunicationSection
}

func NewCommunicationSectionVisitor(section *pb.CommunicationSection) *CommunicationSectionVisitor {
	return &CommunicationSectionVisitor{
		section: section,
	}
}

func (v *CommunicationSectionVisitor) VisitCommunicationSection(ctx *cobol85.CommunicationSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *CommunicationSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *CommunicationSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
