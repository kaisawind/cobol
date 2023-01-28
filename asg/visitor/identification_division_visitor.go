package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/util"
	"github.com/kaisawind/cobol/constant"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type IdentificationDivisionVisitor struct {
	cobol85.BaseCobol85Visitor
	division *pb.IdentificationDivision
}

func NewIdentificationDivisionVisitor(division *pb.IdentificationDivision) *IdentificationDivisionVisitor {
	return &IdentificationDivisionVisitor{
		division: division,
	}
}

func (v *IdentificationDivisionVisitor) VisitProgramIdParagraph(ctx *cobol85.ProgramIdParagraphContext) interface{} {
	attribute := pb.ProgramIdParagraph_COMMON
	switch {
	case ctx.COMMON() != nil:
		attribute = pb.ProgramIdParagraph_COMMON
	case ctx.INITIAL() != nil:
		attribute = pb.ProgramIdParagraph_INITIAL
	case ctx.LIBRARY() != nil:
		attribute = pb.ProgramIdParagraph_LIBRARY
	case ctx.DEFINITION() != nil:
		attribute = pb.ProgramIdParagraph_DEFINITION
	case ctx.RECURSIVE() != nil:
		attribute = pb.ProgramIdParagraph_RECURSIVE
	default:
		attribute = pb.ProgramIdParagraph_COMMON
	}
	v.division.ProgramIdParagraph = &pb.ProgramIdParagraph{
		Name:      ctx.ProgramName().GetText(),
		Attribute: attribute,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitAuthorParagraph(ctx *cobol85.AuthorParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = util.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.AuthorParagraph = &pb.AuthorParagraph{
		Author: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitInstallationParagraph(ctx *cobol85.InstallationParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = util.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.InstallationParagraph = &pb.InstallationParagraph{
		Installation: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitDateWrittenParagraph(ctx *cobol85.DateWrittenParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = util.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.DateWrittenParagraph = &pb.DateWrittenParagraph{
		DateWritten: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitDateCompiledParagraph(ctx *cobol85.DateCompiledParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = util.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.DateCompiledParagraph = &pb.DateCompiledParagraph{
		DateCompiled: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitSecurityParagraph(ctx *cobol85.SecurityParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = util.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.SecurityParagraph = &pb.SecurityParagraph{
		Security: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitRemarksParagraph(ctx *cobol85.RemarksParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = util.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.RemarksParagraph = &pb.RemarksParagraph{
		Remarks: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitIdentificationDivision(ctx *cobol85.IdentificationDivisionContext) any {
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *IdentificationDivisionVisitor) VisitIdentificationDivisionBody(ctx *cobol85.IdentificationDivisionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
