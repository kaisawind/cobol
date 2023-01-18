package model

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data/communication"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/asg/model/procedure"
	"github.com/kaisawind/cobol/asg/util"
)

type ElementType interface {
	*DataDivision | *ProgramUnit | *CompilationUnit | *communication.CommunicationDescriptionEntryInput
}

func getParent[T ElementType](program *program, ctx antlr.Tree) (parent T) {
	registry := program.Registry()
	current := ctx
	for {
		if parent != nil || current == nil {
			break
		}
		current = current.GetParent()
		element := registry.GetElement(current)
		if element != nil {
			if parent, ok := element.(T); ok {
				return parent
			}
		}
	}
	return
}

type program struct {
	registry         *Registry
	compilationUnits []*CompilationUnit
	names            []string
}

func newProgram() *program {
	return &program{
		registry:         NewRegistry(),
		compilationUnits: []*CompilationUnit{},
		names:            []string{},
	}
}

type Program struct {
	*program
}

func NewProgram() *Program {
	return &Program{
		program: newProgram(),
	}
}

func (e *program) Registry() *Registry {
	return e.registry
}

func (e *program) AddElement(element Element) {
	e.Registry().AddElement(element)
}

func (e *program) GetElement(ctx antlr.ParserRuleContext) Element {
	return e.Registry().GetElement(ctx)
}

func (e *program) GetFirstCompilationUnit() *CompilationUnit {
	if len(e.compilationUnits) == 0 {
		return nil
	}
	return e.compilationUnits[0]
}

func (e *program) GetCompilationUnit(name string) *CompilationUnit {
	name = strings.ToLower(name)
	for k, v := range e.names {
		if v == name {
			if len(e.compilationUnits) > k {
				return e.compilationUnits[k]
			} else {
				return nil
			}
		}
	}
	return nil
}

func (e *program) GetCompilationUnits() []*CompilationUnit {
	return e.compilationUnits
}

func (e *program) AddCompilationUnit(element *CompilationUnit) {
	name := strings.ToLower(element.Name())
	e.compilationUnits = append(e.compilationUnits, element)
	e.names = append(e.names, name)
}

func (e *program) GetProgramUnit(ctx antlr.Tree) *ProgramUnit {
	return getParent[*ProgramUnit](e, ctx)
}

func (e *program) GetSection(ctx antlr.ParserRuleContext) (ret *procedure.Section) {
	programUnit := e.GetProgramUnit(ctx)
	if programUnit == nil {
		return
	}
	procedureDivision := programUnit.GetProcedureDivision()
	if procedureDivision == nil {
		return
	}
	ret = procedureDivision.GetSection(util.DetermineName(ctx))
	return
}

func (e *program) GetCommunicationDescriptionEntry(ctx antlr.ParserRuleContext) (ret communication.CommunicationDescriptionEntry) {
	name := util.DetermineName(ctx)
	programUnit := e.GetProgramUnit(ctx)
	if programUnit == nil {
		return
	}
	dataDivision := programUnit.GetDataDivision()
	if dataDivision == nil {
		return
	}
	communicationSection := dataDivision.GetCommunicationSection()
	ret = communicationSection.GetCommunicationDescriptionEntryByName(name)
	return
}

func (e *program) GetDataDescriptionEntries(ctx antlr.ParserRuleContext) (entries []datadescription.DataDescriptionEntry) {
	name := util.DetermineName(ctx)
	programUnit := e.GetProgramUnit(ctx)
	if programUnit == nil {
		return
	}
	dataDivision := programUnit.GetDataDivision()
	if dataDivision == nil {
		return
	}

	workingStorageSection := dataDivision.GetWorkingStorageSection()
	communicationSection := dataDivision.GetCommunicationSection()
	fileSection := dataDivision.GetFileSection()
	localStorageSection := dataDivision.GetLocalStorageSection()
	linkageSection := dataDivision.GetLinkageSection()

	if workingStorageSection != nil {
		entries = append(entries, workingStorageSection.GetDataDescriptionEntriesByName(name)...)
	}

	if communicationSection != nil {
		entries = append(entries, communicationSection.GetDataDescriptionEntriesByName(name)...)
	}

	if fileSection != nil {
		entries = append(entries, fileSection.GetDataDescriptionEntriesByName(name)...)
	}

	if localStorageSection != nil {
		entries = append(entries, localStorageSection.GetDataDescriptionEntriesByName(name)...)
	}

	if linkageSection != nil {
		entries = append(entries, linkageSection.GetDataDescriptionEntriesByName(name)...)
	}
	return
}

func (e *program) GetIndex(ctx antlr.ParserRuleContext) (ret *datadescription.Index) {
	name := util.DetermineName(ctx)
	programUnit := e.GetProgramUnit(ctx)
	if programUnit == nil || programUnit.dataDivision == nil {
		return
	}
	workingStorageSection := programUnit.dataDivision.GetWorkingStorageSection()
	if workingStorageSection == nil {
		return
	}

	for _, entry := range workingStorageSection.GetDataDescriptionEntries() {
		if entry.Type() == datadescription.GROUP {
			dataDescriptionEntryGroup := entry.(*datadescription.DataDescriptionEntryGroup)
			for _, clause := range dataDescriptionEntryGroup.GetOccursClauses() {
				occursIndexed := clause.GetOccursIndexed()
				if occursIndexed != nil {
					index := occursIndexed.GetIndex(name)
					if index != nil {
						ret = index
						return
					}
				}
			}
		}
	}
	return
}
