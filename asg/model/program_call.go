package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/asg/model/procedure"
	"github.com/kaisawind/cobol/asg/util"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type programCall struct {
	*program
}

func newProgramCall() *programCall {
	return &programCall{
		program: newProgram(),
	}
}

func (e *programCall) CreateCall(ctx antlr.ParserRuleContext) (ret call.Call) {
	switch t := ctx.(type) {
	case cobol85.IDataDescNameContext,
		cobol85.IConditionNameContext,
		cobol85.IDataNameContext:
		ret = e.CreateDataDescriptionEntryCall(t)
	case cobol85.IClassNameContext,
		cobol85.ICobolWordContext,
		cobol85.IAlphabetNameContext,
		cobol85.IAssignmentNameContext,
		cobol85.ILibraryNameContext,
		cobol85.ILocalNameContext:
		element := e.GetElement(ctx)
		if element != nil {
			ret, _ = element.(call.Call)
		} else {
			ret = e.CreateUndefinedCall(ctx)
		}

	}
	return
}

func (e *programCall) CreateIdentifierCall(ictx cobol85.IIdentifierContext) (ret call.Call) {
	ctx := ictx.(*cobol85.IdentifierContext)

	if cctx := ctx.QualifiedDataName(); cctx != nil {
		ret = e.CreateQualifiedDataNameCall(cctx)
	} else if cctx := ctx.TableCall(); cctx != nil {

	} else if cctx := ctx.FunctionCall(); cctx != nil {

	} else if cctx := ctx.SpecialRegister(); cctx != nil {

	} else {
		ret = e.CreateDataDescriptionEntryCall(ctx)
	}
	return
}

func (e *programCall) CreateQualifiedDataNameCall(ictx cobol85.IQualifiedDataNameContext) (ret call.Call) {
	ctx := ictx.(*cobol85.QualifiedDataNameContext)

	if cctx := ctx.QualifiedDataNameFormat1(); cctx != nil {

	} else if cctx := ctx.QualifiedDataNameFormat2(); cctx != nil {

	} else if cctx := ctx.QualifiedDataNameFormat3(); cctx != nil {

	} else if cctx := ctx.QualifiedDataNameFormat4(); cctx != nil {

	} else {
		ret = e.CreateDataDescriptionEntryCall(ctx)
	}
	return
}

func (e *programCall) QualifiedInDataCompare(ctxs []cobol85.IQualifiedInDataContext, entry datadescription.DataDescriptionEntry) (ret bool) {
	ret = true
	parent := entry.GetDataDescriptionEntryGroup()
	for _, ctx := range ctxs {
		if parent == nil {
			ret = false
			break
		}

		entryName := util.Symbol(parent.Name())
		ctxName := util.Symbol(util.DetermineName(ctx))

		if entryName == "" || entryName != ctxName {
			ret = false
			break
		}
		parent = parent.GetDataDescriptionEntryGroup()
	}
	return
}

func (e *programCall) CreateQualifiedDataNameFormat1Call(ictx cobol85.IQualifiedDataNameFormat1Context) (ret call.Call) {
	ctx := ictx.(*cobol85.QualifiedDataNameFormat1Context)

	datas := ctx.AllQualifiedInData()
	if len(datas) != 0 {
		name := util.DetermineName(ctx)
		var validEntry datadescription.DataDescriptionEntry
		entries := e.GetDataDescriptionEntries(ctx)
		for _, entry := range entries {
			if e.QualifiedInDataCompare(datas, entry) {
				validEntry = entry
				break
			}
		}

		if validEntry == nil {
			ret = e.CreateUndefinedCall(ctx)
		} else {
			ret = e.CreateDataDescriptionEntryCallWithName(ctx, name, validEntry)
		}
	} else if cctx := ctx.DataName(); cctx != nil {
		ret = e.CreateCall(cctx)
	} else if cctx := ctx.ConditionName(); cctx != nil {
		ret = e.CreateCall(cctx)
	} else {
		ret = e.CreateDataDescriptionEntryCall(ctx)
	}
	return
}

func (e *programCall) CreateQualifiedDataNameFormat2Call(ictx cobol85.IQualifiedDataNameFormat2Context) (ret call.Call) {
	ctx := ictx.(*cobol85.QualifiedDataNameFormat2Context)
	paragraphName := util.DetermineName(ctx)
	section := e.GetSection(ctx.InSection())

	if section == nil {
		ret = e.CreateUndefinedCall(ctx)
	} else {
		paragraph := section.GetParagraphByName(paragraphName)
		if paragraph == nil {
			ret = e.CreateUndefinedCall(ctx)
		} else {
			ret = call.NewProcedureCall(ctx, paragraphName)
		}
	}
	return
}

func (e *programCall) GetSection(ctx antlr.ParserRuleContext) (ret *procedure.Section) {
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

func (e *programCall) GetDataDescriptionEntries(ctx antlr.ParserRuleContext) (entries []datadescription.DataDescriptionEntry) {
	name := util.DetermineName(ctx)
	programUnit := e.GetProgramUnit(ctx)
	if programUnit == nil || programUnit.dataDivision == nil {
		return
	}
	dataDivision := programUnit.dataDivision

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

func (e *programCall) CreateDataDescriptionEntryCall(ctx antlr.ParserRuleContext) (ret call.Call) {
	element := e.GetElement(ctx)
	if element != nil {
		ret, _ = element.(call.Call)
		return
	}

	index := e.GetIndex(ctx)

	if index != nil {
		ret = e.CreateIndexCall(ctx)
	} else {
		entries := e.GetDataDescriptionEntries(ctx)
		if len(entries) == 0 {
			ret = e.CreateUndefinedCall(ctx)
		} else {
			name := util.DetermineName(ctx)
			ret = e.CreateDataDescriptionEntryCallWithName(ctx, name, entries[0])
		}
	}
	return
}

func (e *programCall) CreateDataDescriptionEntryCallWithName(ctx antlr.ParserRuleContext, name string, entry datadescription.DataDescriptionEntry) (ret *call.DataDescriptionEntryCall) {
	element := e.GetElement(ctx)
	if element != nil {
		ret, _ = element.(*call.DataDescriptionEntryCall)
		return
	}

	ret = call.NewDataDescriptionEntryCall(ctx, name, entry)
	// link call with entry
	entry.AddCall(ret)
	e.AddElement(ret)
	return
}

func (e *programCall) CreateUndefinedCall(ctx antlr.ParserRuleContext) (ret call.Call) {
	element := e.GetElement(ctx)
	if element != nil {
		ret, _ = element.(call.Call)
		return
	}

	name := util.DetermineName(ctx)
	ret = call.NewUndefinedCall(ctx, name)
	e.program.AddElement(ret)
	return
}

func (e *programCall) CreateIndexCall(ctx antlr.ParserRuleContext) (ret *call.IndexCall) {
	element := e.GetElement(ctx)
	if element != nil {
		ret, _ = element.(*call.IndexCall)
		return
	}

	name := util.DetermineName(ctx)
	index := e.GetIndex(ctx)

	if index == nil {
		e.CreateDataDescriptionEntryCall(ctx)
	} else {
		ret = call.NewIndexCall(ctx, name, index)
		// link call and index
		index.AddIndexCall(ret)
		e.AddElement(ret)
	}
	return
}

func (e *programCall) GetIndex(ctx antlr.ParserRuleContext) (ret *datadescription.Index) {
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
