package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type ProgramUnitVisitor struct {
	*Visitor
	compilationUnit *model.CompilationUnit
}

func NewProgramUnitVisitor(compilationUnit *model.CompilationUnit, program *model.Program) *ProgramUnitVisitor {
	return &ProgramUnitVisitor{
		Visitor:         NewVisitor(program),
		compilationUnit: compilationUnit,
	}
}

func (v *ProgramUnitVisitor) VisitDataDivision(ctx *cobol85.DataDivisionContext) any {
	element := v.GetElement(ctx)
	if element == nil {
		unit := model.NewDataDivision(ctx)
		programUnit := v.GetProgramUnit(ctx)
		v.AddElement(unit)
		programUnit.SetDataDivision(unit)
	}
	return v.VisitChildren(ctx)
}

func (v *ProgramUnitVisitor) VisitEnvironmentDivision(ctx *cobol85.EnvironmentDivisionContext) any {
	element := v.GetElement(ctx)
	if element == nil {
		unit := model.NewEnvironmentDivision(ctx)
		programUnit := v.GetProgramUnit(ctx)
		v.AddElement(unit)
		programUnit.SetEnvironmentDivision(unit)
	}
	return v.VisitChildren(ctx)
}

func (v *ProgramUnitVisitor) VisitIdentificationDivision(ctx *cobol85.IdentificationDivisionContext) any {
	element := v.GetElement(ctx)
	if element == nil {
		unit := model.NewIdentificationDivision(ctx)
		programUnit := v.GetProgramUnit(ctx)
		v.AddElement(unit)
		programUnit.SetIdentificationDivision(unit)
	}
	return v.VisitChildren(ctx)
}

func (v *ProgramUnitVisitor) VisitProcedureDivision(ctx *cobol85.ProcedureDivisionContext) any {
	element := v.GetElement(ctx)
	if element == nil {
		unit := model.NewProcedureDivision(ctx)
		programUnit := v.GetProgramUnit(ctx)
		v.AddElement(unit)
		programUnit.SetProcedureDivision(unit)
	}
	return v.VisitChildren(ctx)
}

func (v *ProgramUnitVisitor) VisitProgramUnit(ctx *cobol85.ProgramUnitContext) any {
	element := v.GetElement(ctx)
	if element == nil {
		unit := model.NewProgramUnit(ctx)
		v.AddElement(unit)
		v.compilationUnit.AddProgramUnit(unit)
	}
	return v.VisitChildren(ctx)
}

func (v *ProgramUnitVisitor) VisitCompilationUnit(ctx *cobol85.CompilationUnitContext) any {
	return v.VisitChildren(ctx)
}

func (v *ProgramUnitVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ProgramUnitVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
