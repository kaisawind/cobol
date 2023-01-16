package util

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

func DetermineName(ctx antlr.Tree) (name string) {
	switch t := ctx.(type) {
	case *cobol85.AlphabetNameContext:
		name = t.GetText()
	case *cobol85.AssignmentNameContext:
		name = t.GetText()
	case *cobol85.CdNameContext:
		name = t.GetText()
	case *cobol85.ClassNameContext:
		name = t.GetText()
	case *cobol85.CobolWordContext:
		name = t.GetText()
	case *cobol85.ComputerNameContext:
		name = t.GetText()
	case *cobol85.ConditionNameContext:
		name = t.GetText()
	case *cobol85.DataDescriptionEntryFormat1Context:
		name = DetermineName(t.DataName())
	case *cobol85.DataDescriptionEntryFormat2Context:
		name = DetermineName(t.DataName())
	case *cobol85.DataDescriptionEntryFormat3Context:
		name = DetermineName(t.ConditionName())
	case *cobol85.DataDescNameContext:
		name = t.GetText()
	case *cobol85.DataNameContext:
		name = t.GetText()
	case *cobol85.FileControlEntryContext:
		name = DetermineName(t.SelectClause())
	case *cobol85.FileDescriptionEntryContext:
		name = DetermineName(t.FileName())
	case *cobol85.FileNameContext:
		name = t.GetText()
	case *cobol85.FunctionCallContext:
		name = DetermineName(t.FunctionName())
	case *cobol85.FunctionNameContext:
		name = t.GetText()
	case *cobol85.IdentifierContext:
		if t.QualifiedDataName() != nil {
			name = DetermineName(t.QualifiedDataName())
		} else {
			name = t.GetText()
		}
	case *cobol85.IndexNameContext:
		name = t.GetText()
	case *cobol85.InDataContext:
		name = DetermineName(t.DataName())
	case *cobol85.InSectionContext:
		name = DetermineName(t.SectionName())
	case *cobol85.InTableContext:
		name = DetermineName(t.TableCall())
	case *cobol85.LibraryNameContext:
		name = t.GetText()
	case *cobol85.LocalNameContext:
		name = t.GetText()
	case *cobol85.ObjectComputerParagraphContext:
		name = DetermineName(t.ComputerName())
	case *cobol85.ParagraphContext:
		name = DetermineName(t.ParagraphName())
	case *cobol85.ParagraphNameContext:
		name = t.GetText()
	case *cobol85.ProcedureNameContext:
		name = DetermineName(t.ParagraphName())
	case *cobol85.ProcedureSectionContext:
		name = DetermineName(t.ProcedureSectionHeader())
	case *cobol85.ProcedureSectionHeaderContext:
		name = DetermineName(t.SectionName())
	case *cobol85.ProgramIdParagraphContext:
		name = DetermineName(t.ProgramName())
	case *cobol85.ProgramNameContext:
		name = t.GetText()
	case *cobol85.QualifiedDataNameContext:
		if t.QualifiedDataNameFormat1() != nil {
			name = DetermineName(t.QualifiedDataNameFormat1())
		} else if t.QualifiedDataNameFormat2() != nil {
			name = DetermineName(t.QualifiedDataNameFormat2())
		} else if t.QualifiedDataNameFormat3() != nil {
			name = DetermineName(t.QualifiedDataNameFormat3())
		} else if t.QualifiedDataNameFormat4() != nil {
			name = DetermineName(t.QualifiedDataNameFormat4())
		}
	case *cobol85.QualifiedDataNameFormat1Context:
		if t.DataName() != nil {
			name = DetermineName(t.DataName())
		} else if t.ConditionName() != nil {
			name = DetermineName(t.ConditionName())
		}
	case *cobol85.QualifiedDataNameFormat2Context:
		name = DetermineName(t.ParagraphName())
	case *cobol85.QualifiedDataNameFormat3Context:
		name = DetermineName(t.TextName())
	case *cobol85.QualifiedInDataContext:
		if t.InData() != nil {
			name = DetermineName(t.InData())
		} else if t.InTable() != nil {
			name = DetermineName(t.InTable())
		}
	case *cobol85.RecordNameContext:
		name = DetermineName(t.QualifiedDataName())
	case *cobol85.ReportDescriptionContext:
		name = DetermineName(t.ReportDescriptionEntry())
	case *cobol85.ReportDescriptionEntryContext:
		name = DetermineName(t.ReportName())
	case *cobol85.ReportGroupDescriptionEntryFormat1Context:
		name = DetermineName(t.DataName())
	case *cobol85.ReportGroupDescriptionEntryFormat2Context:
		name = DetermineName(t.DataName())
	case *cobol85.ReportGroupDescriptionEntryFormat3Context:
		name = DetermineName(t.DataName())
	case *cobol85.ReportNameContext:
		name = t.GetText()
	case *cobol85.SectionNameContext:
		name = t.GetText()
	case *cobol85.SelectClauseContext:
		name = DetermineName(t.FileName())
	case *cobol85.ScreenNameContext:
		name = t.GetText()
	case *cobol85.ScreenDescriptionEntryContext:
		name = DetermineName(t.ScreenName())
	case *cobol85.SourceComputerParagraphContext:
		name = DetermineName(t.ComputerName())
	case *cobol85.SpecialRegisterContext:
		name = t.GetText()
	case *cobol85.SystemNameContext:
		name = t.GetText()
	case *cobol85.TableCallContext:
		name = t.QualifiedDataName().GetText()
	default:
	}
	return
}

func Symbol(name string) string {
	return strings.ToUpper(name)
}
