package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/data/communication"
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
	element := v.GetElement(ctx)
	if element == nil {
		e := communication.NewCommunicationSection(ctx)

		// CommunicationDescriptionEntry
		for _, vv := range ctx.AllCommunicationDescriptionEntry() {
			entry := vv.(*cobol85.CommunicationDescriptionEntryContext)
			if f1 := entry.CommunicationDescriptionEntryFormat1(); f1 != nil {
				f1Ctx := f1.(*cobol85.CommunicationDescriptionEntryFormat1Context)
				if v.GetElement(f1Ctx) == nil {
					input := communication.NewCommunicationDescriptionEntryInput(f1Ctx)
					// TODO

					// symbolic queue
					clauses := f1Ctx.AllSymbolicQueueClause()
					if len(clauses) != 0 {

					}
					e.AddCommunicationDescriptionEntry(input)
					v.AddElement(input)
				}
			} else if f2 := entry.CommunicationDescriptionEntryFormat2(); f2 != nil {
				f2Ctx := f2.(*cobol85.CommunicationDescriptionEntryFormat2Context)
				if v.GetElement(f2Ctx) == nil {
					ioput := communication.NewCommunicationDescriptionEntryInputOutput(f2Ctx)
					// TODO
					e.AddCommunicationDescriptionEntry(ioput)
					v.AddElement(ioput)
				}
			} else if f3 := entry.CommunicationDescriptionEntryFormat3(); f3 != nil {
				f3Ctx := f3.(*cobol85.CommunicationDescriptionEntryFormat3Context)
				if v.GetElement(f3Ctx) == nil {
					output := communication.NewCommunicationDescriptionEntryOutput(f3Ctx)
					// TODO
					e.AddCommunicationDescriptionEntry(output)
					v.AddElement(output)
				}
			} else {
				// wrong
			}
		}
		dataDivision := model.GetElement[*model.DataDivision](ctx, v.program)
		dataDivision.SetCommunicationSection(e)
		v.AddElement(e)
	}
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
