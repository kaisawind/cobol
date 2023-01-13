package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/data/communication"
	"github.com/kaisawind/cobol/asg/util"
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
	if v.GetElement(ctx) == nil {
		e := communication.NewCommunicationSection(ctx)

		// CommunicationDescriptionEntry
		for _, vv := range ctx.AllCommunicationDescriptionEntry() {
			entry := vv.(*cobol85.CommunicationDescriptionEntryContext)
			if f1 := entry.CommunicationDescriptionEntryFormat1(); f1 != nil {
				f1Ctx := f1.(*cobol85.CommunicationDescriptionEntryFormat1Context)
				if v.GetElement(f1Ctx) == nil {
					name := util.DetermineName(f1Ctx.CdName())
					input := communication.NewCommunicationDescriptionEntryInput(f1Ctx, name)

					// symbolic queue
					clauses := f1Ctx.AllSymbolicQueueClause()
					if len(clauses) != 0 {
						clause := clauses[0].(*cobol85.SymbolicQueueClauseContext)
						if v.GetElement(clause) == nil {
							clauseElement := communication.NewSymbolicQueueClause(clause)

							dataDescName := clause.DataDescName()
							if v.GetElement(dataDescName) == nil {
								dataDescCall := model.CreateCall(v.program, dataDescName)
								clauseElement.SetDataDescCall(dataDescCall)
							}

							input.SetSymbolicQueueClause(clauseElement)
							v.AddElement(clauseElement)
						}
					}
					e.AddCommunicationDescriptionEntry(input)
					v.AddElement(input)
				}
			} else if f2 := entry.CommunicationDescriptionEntryFormat2(); f2 != nil {
				f2Ctx := f2.(*cobol85.CommunicationDescriptionEntryFormat2Context)
				if v.GetElement(f2Ctx) == nil {
					name := util.DetermineName(f2Ctx.CdName())
					ioput := communication.NewCommunicationDescriptionEntryInputOutput(f2Ctx, name)
					// TODO
					e.AddCommunicationDescriptionEntry(ioput)
					v.AddElement(ioput)
				}
			} else if f3 := entry.CommunicationDescriptionEntryFormat3(); f3 != nil {
				f3Ctx := f3.(*cobol85.CommunicationDescriptionEntryFormat3Context)
				if v.GetElement(f3Ctx) == nil {
					name := util.DetermineName(f3Ctx.CdName())
					output := communication.NewCommunicationDescriptionEntryOutput(f3Ctx, name)
					// TODO
					e.AddCommunicationDescriptionEntry(output)
					v.AddElement(output)
				}
			} else {
				// wrong
			}
		}
		dataDivision := model.GetParent[*model.DataDivision](v.program, ctx)
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
