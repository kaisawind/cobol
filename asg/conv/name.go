package conv

import (
	"github.com/kaisawind/cobol/gen/cobol85"
	"github.com/kaisawind/cobol/pb"
)

func FunctionName(in cobol85.IFunctionNameContext) *pb.FunctionName {
	return &pb.FunctionName{
		Value: in.GetText(),
	}
}

func ClassName(in cobol85.IClassNameContext) *pb.ClassName {
	ctx := in.(*cobol85.ClassNameContext)
	return &pb.ClassName{
		CobolWord: CobolWord(ctx.CobolWord()),
	}
}

func MnemonicName(in cobol85.IMnemonicNameContext) *pb.MnemonicName {
	ctx := in.(*cobol85.MnemonicNameContext)
	return &pb.MnemonicName{
		CobolWord: CobolWord(ctx.CobolWord()),
	}
}

func AlphabetName(in cobol85.IAlphabetNameContext) *pb.AlphabetName {
	ctx := in.(*cobol85.AlphabetNameContext)
	return &pb.AlphabetName{
		CobolWord: CobolWord(ctx.CobolWord()),
	}
}

func ComputerName(in cobol85.IComputerNameContext) *pb.ComputerName {
	ctx := in.(*cobol85.ComputerNameContext)
	return &pb.ComputerName{
		SystemName: SystemName(ctx.SystemName()),
	}
}

func SystemName(in cobol85.ISystemNameContext) *pb.SystemName {
	ctx := in.(*cobol85.SystemNameContext)
	return &pb.SystemName{
		CobolWord: CobolWord(ctx.CobolWord()),
	}
}

func AssignmentName(in cobol85.IAssignmentNameContext) *pb.AssignmentName {
	ctx := in.(*cobol85.AssignmentNameContext)
	return &pb.AssignmentName{
		SystemName: SystemName(ctx.SystemName()),
	}
}

func CobolWord(in cobol85.ICobolWordContext) *pb.CobolWord {
	return &pb.CobolWord{
		Value: in.GetText(),
	}
}

func DataName(in cobol85.IDataNameContext) *pb.DataName {
	ctx := in.(*cobol85.DataNameContext)
	return &pb.DataName{
		CobolWord: CobolWord(ctx.CobolWord()),
	}
}

func FileName(in cobol85.IFileNameContext) *pb.FileName {
	ctx := in.(*cobol85.FileNameContext)
	return &pb.FileName{
		CobolWord: CobolWord(ctx.CobolWord()),
	}
}

func ConditionName(in cobol85.IConditionNameContext) *pb.ConditionName {
	ctx := in.(*cobol85.ConditionNameContext)
	return &pb.ConditionName{
		CobolWord: CobolWord(ctx.CobolWord()),
	}
}

func IndexName(in cobol85.IIndexNameContext) *pb.IndexName {
	ctx := in.(*cobol85.IndexNameContext)
	return &pb.IndexName{
		CobolWord: CobolWord(ctx.CobolWord()),
	}
}

func ParagraphName(in cobol85.IParagraphNameContext) (out *pb.ParagraphName) {
	ctx := in.(*cobol85.ParagraphNameContext)
	out = &pb.ParagraphName{}
	if ictx := ctx.CobolWord(); ictx != nil {
		out.OneOf = &pb.ParagraphName_CobolWord{
			CobolWord: CobolWord(ictx),
		}
	} else if ictx := ctx.IntegerLiteral(); ictx != nil {
		out.OneOf = &pb.ParagraphName_IntegerLiteral{
			IntegerLiteral: IntegerLiteral(ictx),
		}
	}
	return
}

func SectionName(in cobol85.ISectionNameContext) (out *pb.SectionName) {
	ctx := in.(*cobol85.SectionNameContext)
	out = &pb.SectionName{}
	if ictx := ctx.CobolWord(); ictx != nil {
		out.OneOf = &pb.SectionName_CobolWord{
			CobolWord: CobolWord(ictx),
		}
	} else if ictx := ctx.IntegerLiteral(); ictx != nil {
		out.OneOf = &pb.SectionName_IntegerLiteral{
			IntegerLiteral: IntegerLiteral(ictx),
		}
	}
	return
}

func TextName(in cobol85.ITextNameContext) *pb.TextName {
	ctx := in.(*cobol85.TextNameContext)
	return &pb.TextName{
		CobolWord: CobolWord(ctx.CobolWord()),
	}
}

func LibraryName(in cobol85.ILibraryNameContext) *pb.LibraryName {
	ctx := in.(*cobol85.LibraryNameContext)
	return &pb.LibraryName{
		CobolWord: CobolWord(ctx.CobolWord()),
	}
}