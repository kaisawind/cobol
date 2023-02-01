package environment

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/conv"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type InputOutputSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.InputOutputSection
}

func NewInputOutputSectionVisitor(section *pb.InputOutputSection) *InputOutputSectionVisitor {
	return &InputOutputSectionVisitor{
		section: section,
	}
}

func (v *InputOutputSectionVisitor) VisitFileControlParagraph(ctx *cobol85.FileControlParagraphContext) interface{} {
	fcp := &pb.FileControlParagraph{}
	for _, ictx := range ctx.AllFileControlEntry() {
		cctx := ictx.(*cobol85.FileControlEntryContext)
		fce := &pb.FileControlEntry{}
		if icctx := cctx.SelectClause(); icctx != nil {
			cc, ok := icctx.(*cobol85.SelectClauseContext)
			if ok {
				fce.SelectClause = &pb.SelectClause{
					FileName: conv.FileName(cc.FileName()),
					Optional: cc.OPTIONAL() != nil,
				}
				fce.FileName = conv.FileName(cc.FileName())
			}
		}
		for _, icctx := range cctx.AllFileControlClause() {
			cc := icctx.(*cobol85.FileControlClauseContext)
			if icc := cc.ReserveClause(); icc != nil {
				iccc := icc.(*cobol85.ReserveClauseContext)
				fce.ReserveClause = &pb.ReserveClause{
					Value: conv.IntegerLiteral(iccc.IntegerLiteral()),
				}
			}
			if icc := cc.AccessModeClause(); icc != nil {
				iccc := icc.(*cobol85.AccessModeClauseContext)
				mode := pb.AccessModeClause_DYNAMIC
				switch {
				case iccc.DYNAMIC() != nil:
					mode = pb.AccessModeClause_DYNAMIC
				case iccc.SEQUENTIAL() != nil:
					mode = pb.AccessModeClause_SEQUENTIAL
				case iccc.RANDOM() != nil:
					mode = pb.AccessModeClause_RANDOM
				case iccc.EXCLUSIVE() != nil:
					mode = pb.AccessModeClause_EXCLUSIVE
				}
				fce.AccessModeClause = &pb.AccessModeClause{
					Mode: mode,
				}
			}
			if icc := cc.AlternateRecordKeyClause(); icc != nil {
				iccc := icc.(*cobol85.AlternateRecordKeyClauseContext)
				pcctx := iccc.PasswordClause().(*cobol85.PasswordClauseContext)
				fce.AlternateRecordKeyClause = &pb.AlternateRecordKeyClause{
					PasswordClause: &pb.PasswordClause{
						DataName: conv.DataName(pcctx.DataName()),
					},
					QualifiedDataName: conv.QualifiedDataName(iccc.QualifiedDataName()),
				}
			}
			if icc := cc.AssignClause(); icc != nil {
				iccc := icc.(*cobol85.AssignClauseContext)
				fce.AssignClause = &pb.AssignClause{}
				switch {
				case iccc.DISK() != nil:
					fce.AssignClause.ToValue = &pb.AssignClause_Type_{
						Type: pb.AssignClause_DISK,
					}
				case iccc.DISPLAY() != nil:
					fce.AssignClause.ToValue = &pb.AssignClause_Type_{
						Type: pb.AssignClause_DISPLAY,
					}
				case iccc.KEYBOARD() != nil:
					fce.AssignClause.ToValue = &pb.AssignClause_Type_{
						Type: pb.AssignClause_KEYBOARD,
					}
				case iccc.PORT() != nil:
					fce.AssignClause.ToValue = &pb.AssignClause_Type_{
						Type: pb.AssignClause_PORT,
					}
				case iccc.PRINTER() != nil:
					fce.AssignClause.ToValue = &pb.AssignClause_Type_{
						Type: pb.AssignClause_PRINTER,
					}
				case iccc.READER() != nil:
					fce.AssignClause.ToValue = &pb.AssignClause_Type_{
						Type: pb.AssignClause_READER,
					}
				case iccc.REMOTE() != nil:
					fce.AssignClause.ToValue = &pb.AssignClause_Type_{
						Type: pb.AssignClause_REMOTE,
					}
				case iccc.TAPE() != nil:
					fce.AssignClause.ToValue = &pb.AssignClause_Type_{
						Type: pb.AssignClause_TAPE,
					}
				case iccc.VIRTUAL() != nil:
					fce.AssignClause.ToValue = &pb.AssignClause_Type_{
						Type: pb.AssignClause_VIRTUAL,
					}
				case iccc.AssignmentName() != nil:
					fce.AssignClause.ToValue = &pb.AssignClause_AssignmentName{
						AssignmentName: conv.AssignmentName(iccc.AssignmentName()),
					}
				case iccc.Literal() != nil:
					fce.AssignClause.ToValue = &pb.AssignClause_Literal{
						Literal: conv.Literal(iccc.Literal()),
					}
				}
			}
			if icc := cc.FileStatusClause(); icc != nil {
				iccc := icc.(*cobol85.FileStatusClauseContext)
				fce.FileStatusClause = &pb.FileStatusClause{}
				for _, v := range iccc.AllQualifiedDataName() {
					if fce.FileStatusClause.QualifiedDataName_1 == nil {
						fce.FileStatusClause.QualifiedDataName_1 = conv.QualifiedDataName(v)
					} else {
						fce.FileStatusClause.QualifiedDataName_1 = conv.QualifiedDataName(v)
					}
				}
			}
			if icc := cc.OrganizationClause(); icc != nil {
				iccc := icc.(*cobol85.OrganizationClauseContext)
				var mode pb.OrganizationClause_Mode
				var typ pb.OrganizationClause_Type
				switch {
				case iccc.INDEXED() != nil:
					mode = pb.OrganizationClause_INDEXED
				case iccc.RELATIVE() != nil:
					mode = pb.OrganizationClause_RELATIVE
				case iccc.SEQUENTIAL() != nil:
					mode = pb.OrganizationClause_SEQUENTIAL
				}
				switch {
				case iccc.LINE() != nil:
					typ = pb.OrganizationClause_LINE
				case iccc.RECORD() != nil:
					if iccc.BINARY() != nil {
						typ = pb.OrganizationClause_RECORD_BINARY
					} else {
						typ = pb.OrganizationClause_RECORD
					}
				case iccc.BINARY() != nil:
					typ = pb.OrganizationClause_BINARY
				}
				fce.OrganizationClause = &pb.OrganizationClause{
					Type: typ,
					Mode: mode,
				}
			}
			if icc := cc.PaddingCharacterClause(); icc != nil {
				iccc := icc.(*cobol85.PaddingCharacterClauseContext)
				fce.PaddingCharacterClause = &pb.PaddingCharacterClause{}
				if qdn := iccc.QualifiedDataName(); qdn != nil {
					fce.PaddingCharacterClause.PaddingCharacter = &pb.PaddingCharacterClause_QualifiedDataName{
						QualifiedDataName: conv.QualifiedDataName(qdn),
					}
				} else if literal := iccc.Literal(); literal != nil {
					fce.PaddingCharacterClause.PaddingCharacter = &pb.PaddingCharacterClause_Literal{
						Literal: conv.Literal(literal),
					}
				}
			}
			if icc := cc.PasswordClause(); icc != nil {
				iccc := icc.(*cobol85.PasswordClauseContext)
				fce.PasswordClause = &pb.PasswordClause{
					DataName: conv.DataName(iccc.DataName()),
				}
			}
			if icc := cc.RecordDelimiterClause(); icc != nil {
				iccc := icc.(*cobol85.RecordDelimiterClauseContext)
				switch {
				case iccc.AssignmentName() != nil:
					fce.RecordDelimiterClause = &pb.RecordDelimiterClause{
						Type:           pb.RecordDelimiterClause_ASSIGNMENT,
						AssignmentName: conv.AssignmentName(iccc.AssignmentName()),
					}
				case iccc.IMPLICIT() != nil:
					fce.RecordDelimiterClause = &pb.RecordDelimiterClause{
						Type: pb.RecordDelimiterClause_IMPLICIT,
					}
				case iccc.STANDARD_1() != nil:
					fce.RecordDelimiterClause = &pb.RecordDelimiterClause{
						Type: pb.RecordDelimiterClause_STANDARD_1,
					}
				}
			}
			if icc := cc.RecordKeyClause(); icc != nil {
				iccc := icc.(*cobol85.RecordKeyClauseContext)
				pcctx := iccc.PasswordClause().(*cobol85.PasswordClauseContext)
				fce.RecordKeyClause = &pb.RecordKeyClause{
					QualifiedDataName: conv.QualifiedDataName(iccc.QualifiedDataName()),
					PasswordClause: &pb.PasswordClause{
						DataName: conv.DataName(pcctx.DataName()),
					},
				}
			}
			if icc := cc.RelativeKeyClause(); icc != nil {
				iccc := icc.(*cobol85.RelativeKeyClauseContext)
				fce.RelativeKeyClause = &pb.RelativeKeyClause{
					QualifiedDataName: conv.QualifiedDataName(iccc.QualifiedDataName()),
				}
			}
			if icc := cc.ReserveClause(); icc != nil {
				iccc := icc.(*cobol85.ReserveClauseContext)
				fce.ReserveClause = &pb.ReserveClause{
					Value: conv.IntegerLiteral(iccc.IntegerLiteral()),
				}
			}
			if icc := cc.AccessModeClause(); icc != nil {
				iccc := icc.(*cobol85.AccessModeClauseContext)
				mode := pb.AccessModeClause_DYNAMIC
				switch {
				case iccc.DYNAMIC() != nil:
					mode = pb.AccessModeClause_DYNAMIC
				case iccc.EXCLUSIVE() != nil:
					mode = pb.AccessModeClause_EXCLUSIVE
				case iccc.RANDOM() != nil:
					mode = pb.AccessModeClause_RANDOM
				case iccc.SEQUENTIAL() != nil:
					mode = pb.AccessModeClause_SEQUENTIAL
				}
				fce.AccessModeClause = &pb.AccessModeClause{
					Mode: mode,
				}
			}
		}
		fcp.FileControlEntries = append(fcp.FileControlEntries, fce)
	}
	v.section.FileControlParagraph = fcp
	return v.VisitChildren(ctx)
}

func (v *InputOutputSectionVisitor) VisitIoControlParagraph(ctx *cobol85.IoControlParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *InputOutputSectionVisitor) VisitInputOutputSectionParagraph(ctx *cobol85.InputOutputSectionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *InputOutputSectionVisitor) VisitInputOutputSection(ctx *cobol85.InputOutputSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *InputOutputSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *InputOutputSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
