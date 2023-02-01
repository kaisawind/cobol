package environment

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/conv"
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

type ConfigurationSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.ConfigurationSection
}

func NewConfigurationSectionVisitor(section *pb.ConfigurationSection) *ConfigurationSectionVisitor {
	return &ConfigurationSectionVisitor{
		section: section,
	}
}

func (v *ConfigurationSectionVisitor) VisitSegmentLimitClause(ctx *cobol85.SegmentLimitClauseContext) interface{} {
	var integerLiteral *pb.IntegerLiteral
	if ictx := ctx.IntegerLiteral(); ictx != nil {
		integerLiteral = &pb.IntegerLiteral{
			Value: ictx.GetText(),
		}
	}
	v.section.ObjectComputerParagraph.SegmentLimitClause = &pb.SegmentLimitClause{
		SegmentLimit: integerLiteral,
	}
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitCharacterSetClause(ctx *cobol85.CharacterSetClauseContext) interface{} {
	v.section.ObjectComputerParagraph.CharacterSetClause = &pb.CharacterSetClause{}
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitCollatingSequenceClause(ctx *cobol85.CollatingSequenceClauseContext) interface{} {
	names := []*pb.AlphabetName{}
	for _, ictx := range ctx.AllAlphabetName() {
		names = append(names, &pb.AlphabetName{
			CobolWord: &pb.CobolWord{
				Value: conv.DetermineName(ictx),
			},
		})
	}
	alphanumeric := ""
	if ictx := ctx.CollatingSequenceClauseAlphanumeric(); ictx != nil {
		alphanumeric = conv.DetermineName(ictx)
	}
	national := ""
	if ictx := ctx.CollatingSequenceClauseNational(); ictx != nil {
		national = conv.DetermineName(ictx)
	}
	v.section.ObjectComputerParagraph.CollatingSequenceClause = &pb.CollatingSequenceClause{
		National: &pb.AlphabetName{
			CobolWord: &pb.CobolWord{
				Value: national,
			},
		},
		AlphabetNames: names,
		Alphanumeric: &pb.AlphabetName{
			CobolWord: &pb.CobolWord{
				Value: alphanumeric,
			},
		},
	}
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitDiskSizeClause(ctx *cobol85.DiskSizeClauseContext) interface{} {
	unit := pb.DiskSizeClause_WORDS
	switch {
	case ctx.WORDS() != nil:
		unit = pb.DiskSizeClause_WORDS
	case ctx.MODULES() != nil:
		unit = pb.DiskSizeClause_MODULES
	}
	diskSizeClause := &pb.DiskSizeClause{
		Unit: unit,
	}
	if ictx := ctx.IntegerLiteral(); ictx != nil {
		diskSizeClause.DiskSize = &pb.DiskSizeClause_IntegerLiteral{
			IntegerLiteral: &pb.IntegerLiteral{
				Value: ictx.GetText(),
			},
		}
	} else if ictx := ctx.CobolWord(); ictx != nil {
		diskSizeClause.DiskSize = &pb.DiskSizeClause_CobolWord{
			CobolWord: &pb.CobolWord{
				Value: ictx.GetText(),
			},
		}
	}
	v.section.ObjectComputerParagraph.DiskSizeClause = diskSizeClause
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitMemorySizeClause(ctx *cobol85.MemorySizeClauseContext) interface{} {
	unit := pb.MemorySizeClause_WORDS
	switch {
	case ctx.WORDS() != nil:
		unit = pb.MemorySizeClause_WORDS
	case ctx.CHARACTERS() != nil:
		unit = pb.MemorySizeClause_CHARACTERS
	case ctx.MODULES() != nil:
		unit = pb.MemorySizeClause_MODULES
	}

	memorySizeClause := &pb.MemorySizeClause{
		Unit: unit,
	}
	if ictx := ctx.IntegerLiteral(); ictx != nil {
		memorySizeClause.MemorySize = &pb.MemorySizeClause_IntegerLiteral{
			IntegerLiteral: &pb.IntegerLiteral{
				Value: ictx.GetText(),
			},
		}
	} else if ictx := ctx.CobolWord(); ictx != nil {
		memorySizeClause.MemorySize = &pb.MemorySizeClause_CobolWord{
			CobolWord: &pb.CobolWord{
				Value: ictx.GetText(),
			},
		}
	}
	v.section.ObjectComputerParagraph.MemorySizeClause = memorySizeClause
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitObjectComputerClause(ctx *cobol85.ObjectComputerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitObjectComputerParagraph(ctx *cobol85.ObjectComputerParagraphContext) interface{} {
	v.section.ObjectComputerParagraph = &pb.ObjectComputerParagraph{
		ComputerName: &pb.ComputerName{
			SystemName: &pb.SystemName{
				CobolWord: &pb.CobolWord{
					Value: conv.DetermineName(ctx),
				},
			},
		},
	}
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitSourceComputerParagraph(ctx *cobol85.SourceComputerParagraphContext) interface{} {
	v.section.SourceComputerParagraph = &pb.SourceComputerParagraph{
		ComputerName: &pb.ComputerName{
			SystemName: &pb.SystemName{
				CobolWord: &pb.CobolWord{
					Value: conv.DetermineName(ctx),
				},
			},
		},
		DebuggingMode: ctx.DEBUGGING() != nil,
	}
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitConfigurationSectionParagraph(ctx *cobol85.ConfigurationSectionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitConfigurationSection(ctx *cobol85.ConfigurationSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ConfigurationSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
