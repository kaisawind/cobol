package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/util"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type EnvironmentDivisionVisitor struct {
	cobol85.BaseCobol85Visitor
	division *pb.EnvironmentDivision
}

func NewEnvironmentDivisionVisitor(division *pb.EnvironmentDivision) *EnvironmentDivisionVisitor {
	return &EnvironmentDivisionVisitor{
		division: division,
	}
}

func (v *EnvironmentDivisionVisitor) VisitSegmentLimitClause(ctx *cobol85.SegmentLimitClauseContext) interface{} {
	var integerLiteral *pb.IntegerLiteral
	if ictx := ctx.IntegerLiteral(); ictx != nil {
		integerLiteral = &pb.IntegerLiteral{
			Value: ictx.GetText(),
		}
	}
	v.division.ConfigurationSection.ObjectComputerParagraph.SegmentLimitClause = &pb.SegmentLimitClause{
		IntegerLiteral: integerLiteral,
	}
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitCharacterSetClause(ctx *cobol85.CharacterSetClauseContext) interface{} {
	v.division.ConfigurationSection.ObjectComputerParagraph.CharacterSetClause = &pb.CharacterSetClause{}
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitCollatingSequenceClause(ctx *cobol85.CollatingSequenceClauseContext) interface{} {
	names := []string{}
	for _, ictx := range ctx.AllAlphabetName() {
		names = append(names, util.DetermineName(ictx))
	}
	alphanumeric := ""
	if ictx := ctx.CollatingSequenceClauseAlphanumeric(); ictx != nil {
		alphanumeric = util.DetermineName(ictx)
	}
	national := ""
	if ictx := ctx.CollatingSequenceClauseNational(); ictx != nil {
		national = util.DetermineName(ictx)
	}
	v.division.ConfigurationSection.ObjectComputerParagraph.CollatingSequenceClause = &pb.CollatingSequenceClause{
		National:      national,
		AlphabetNames: names,
		Alphanumeric:  alphanumeric,
	}
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitDiskSizeClause(ctx *cobol85.DiskSizeClauseContext) interface{} {
	unit := pb.DiskSizeClause_WORDS
	switch {
	case ctx.WORDS() != nil:
		unit = pb.DiskSizeClause_WORDS
	case ctx.MODULES() != nil:
		unit = pb.DiskSizeClause_MODULES
	}
	var value *pb.ValueStmt
	if ictx := ctx.IntegerLiteral(); ictx != nil {
		value = util.CreateValueStmt(ictx)
	} else if ictx := ctx.CobolWord(); ictx != nil {
		value = util.CreateValueStmt(ictx)
	}
	v.division.ConfigurationSection.ObjectComputerParagraph.DiskSizeClause = &pb.DiskSizeClause{
		ValueStmt: value,
		Unit:      unit,
	}
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitMemorySizeClause(ctx *cobol85.MemorySizeClauseContext) interface{} {
	unit := pb.MemorySizeClause_WORDS
	switch {
	case ctx.WORDS() != nil:
		unit = pb.MemorySizeClause_WORDS
	case ctx.CHARACTERS() != nil:
		unit = pb.MemorySizeClause_CHARACTERS
	case ctx.MODULES() != nil:
		unit = pb.MemorySizeClause_MODULES
	}
	var value *pb.ValueStmt
	if ictx := ctx.IntegerLiteral(); ictx != nil {
		value = util.CreateValueStmt(ictx)
	} else if ictx := ctx.CobolWord(); ictx != nil {
		value = util.CreateValueStmt(ictx)
	}
	v.division.ConfigurationSection.ObjectComputerParagraph.MemorySizeClause = &pb.MemorySizeClause{
		ValueStmt: value,
		Unit:      unit,
	}
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitObjectComputerClause(ctx *cobol85.ObjectComputerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitObjectComputerParagraph(ctx *cobol85.ObjectComputerParagraphContext) interface{} {
	v.division.ConfigurationSection.ObjectComputerParagraph = &pb.ObjectComputerParagraph{
		Name: util.DetermineName(ctx),
	}
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitSourceComputerParagraph(ctx *cobol85.SourceComputerParagraphContext) interface{} {
	v.division.ConfigurationSection.SourceComputerParagraph = &pb.SourceComputerParagraph{
		Name:          util.DetermineName(ctx),
		DebuggingMode: ctx.DEBUGGING() != nil,
	}
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitConfigurationSectionParagraph(ctx *cobol85.ConfigurationSectionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitConfigurationSection(ctx *cobol85.ConfigurationSectionContext) interface{} {
	v.division.ConfigurationSection = &pb.ConfigurationSection{}
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitEnvironmentDivision(ctx *cobol85.EnvironmentDivisionContext) any {
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitEnvironmentDivisionBody(ctx *cobol85.EnvironmentDivisionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *EnvironmentDivisionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
