package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type DataDivisionVisitor struct {
	*Visitor
}

func NewDataDivisionVisitor(program *model.Program) *DataDivisionVisitor {
	return &DataDivisionVisitor{
		Visitor: NewVisitor(program),
	}
}

func (v *DataDivisionVisitor) VisitCommunicationSection(ctx *cobol85.CommunicationSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitWorkingStorageSection(ctx *cobol85.WorkingStorageSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitProgramLibrarySection(ctx *cobol85.ProgramLibrarySectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitLocalStorageSection(ctx *cobol85.LocalStorageSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitLinkageSection(ctx *cobol85.LinkageSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitFileSection(ctx *cobol85.FileSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitProgramUnit(ctx *cobol85.ProgramUnitContext) any {
	return v.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *DataDivisionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
