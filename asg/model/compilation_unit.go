package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type CompilationUnit struct {
	ctx          cobol85.ICompilationUnitContext
	name         string
	programUnits []*ProgramUnit
}

func NewCompilationUnit(ctx cobol85.ICompilationUnitContext, name string) *CompilationUnit {
	compilationUnit := &CompilationUnit{
		ctx:          ctx,
		name:         name,
		programUnits: []*ProgramUnit{},
	}
	return compilationUnit
}

func (e *CompilationUnit) Name() string {
	return e.name
}

func (e *CompilationUnit) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *CompilationUnit) CompilationUnitContext() *cobol85.CompilationUnitContext {
	return e.ctx.(*cobol85.CompilationUnitContext)
}

func (e *CompilationUnit) AddProgramUnit(unit *ProgramUnit) {
	e.programUnits = append(e.programUnits, unit)
}

func (e *CompilationUnit) GetProgramUnit() *ProgramUnit {
	if len(e.programUnits) != 0 {
		return e.programUnits[0]
	}
	return nil
}

func (e *CompilationUnit) GetProgramUnits() []*ProgramUnit {
	return e.programUnits
}
