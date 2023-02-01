package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/visitor/data"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type DataDivisionVisitor struct {
	cobol85.BaseCobol85Visitor
	division *pb.DataDivision
}

func NewDataDivisionVisitor(division *pb.DataDivision) *DataDivisionVisitor {
	return &DataDivisionVisitor{
		division: division,
	}
}

func (v *DataDivisionVisitor) VisitLinkageSection(ctx *cobol85.LinkageSectionContext) interface{} {
	section := &pb.LinkageSection{}
	vr := data.NewLinkageSectionVisitor(section)
	return vr.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitCommunicationSection(ctx *cobol85.CommunicationSectionContext) interface{} {
	section := &pb.CommunicationSection{}
	vr := data.NewCommunicationSectionVisitor(section)
	return vr.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitDataBaseSection(ctx *cobol85.DataBaseSectionContext) interface{} {
	section := &pb.DataBaseSection{}
	vr := data.NewDataBaseSectionVisitor(section)
	return vr.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitLocalStorageSection(ctx *cobol85.LocalStorageSectionContext) interface{} {
	section := &pb.LocalStorageSection{}
	vr := data.NewLocalStorageSectionVisitor(section)
	return vr.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitProgramLibrarySection(ctx *cobol85.ProgramLibrarySectionContext) interface{} {
	section := &pb.ProgramLibrarySection{}
	vr := data.NewProgramLibrarySectionVisitor(section)
	return vr.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitReportSection(ctx *cobol85.ReportSectionContext) interface{} {
	section := &pb.ReportSection{}
	vr := data.NewReportSectionVisitor(section)
	return vr.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitScreenSection(ctx *cobol85.ScreenSectionContext) interface{} {
	section := &pb.ScreenSection{}
	vr := data.NewScreenSectionVisitor(section)
	return vr.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitWorkingStorageSection(ctx *cobol85.WorkingStorageSectionContext) interface{} {
	section := &pb.WorkingStorageSection{}
	vr := data.NewWorkingStorageSectionVisitor(section)
	return vr.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitFileSection(ctx *cobol85.FileSectionContext) interface{} {
	section := &pb.FileSection{}
	vr := data.NewFileSectionVisitor(section)
	return vr.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitDataDivisionSection(ctx *cobol85.DataDivisionSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DataDivisionVisitor) VisitDataDivision(ctx *cobol85.DataDivisionContext) any {
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
