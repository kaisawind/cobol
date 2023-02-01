package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type ProgramLibrarySectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.ProgramLibrarySection
}

func NewProgramLibrarySectionVisitor(section *pb.ProgramLibrarySection) *ProgramLibrarySectionVisitor {
	return &ProgramLibrarySectionVisitor{
		section: section,
	}
}

func (v *ProgramLibrarySectionVisitor) VisitProgramLibrarySection(ctx *cobol85.ProgramLibrarySectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *ProgramLibrarySectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ProgramLibrarySectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
