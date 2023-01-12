package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CompilationUnitVisitor struct {
	*Visitor
	name string
}

func NewCompilationUnitVisitor(name string, program *model.Program) *CompilationUnitVisitor {
	return &CompilationUnitVisitor{
		Visitor: NewVisitor(program),
		name:    name,
	}
}

func (v *CompilationUnitVisitor) VisitCompilationUnit(ctx *cobol85.CompilationUnitContext) any {
	unit := model.NewCompilationUnit(ctx, v.name)
	v.AddCompilationUnit(unit)
	v.AddElement(unit)
	return v.VisitChildren(ctx)
}

func (v *CompilationUnitVisitor) VisitStartRule(ctx *cobol85.StartRuleContext) any {
	return v.VisitChildren(ctx)
}

func (v *CompilationUnitVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *CompilationUnitVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
