package data

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type ScreenSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.ScreenSection
}

func NewScreenSectionVisitor(section *pb.ScreenSection) *ScreenSectionVisitor {
	return &ScreenSectionVisitor{
		section: section,
	}
}

func (v *ScreenSectionVisitor) VisitScreenSection(ctx *cobol85.ScreenSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *ScreenSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ScreenSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
