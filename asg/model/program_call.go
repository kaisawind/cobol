package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/asg/model/procedure"
	"github.com/kaisawind/cobol/asg/model/valuestmt"
	"github.com/kaisawind/cobol/asg/util"
	"github.com/kaisawind/cobol/gen/cobol85"
)

func (e *program) CreateCall(ctxs ...antlr.ParserRuleContext) (ret call.Call) {
	for _, ctx := range ctxs {
		if ret != nil {
			break
		}
		if ctx == nil {
			continue
		}
		switch t := ctx.(type) {
		case cobol85.IIdentifierContext:
			ret = e.CreateIdentifierCall(t)
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
			element := e.GetElement(t)
			if element != nil {
				ret, _ = element.(call.Call)
			} else {
				ret = e.CreateUndefinedCall(t)
			}

		}
	}
	return
}

func (e *program) CreateIdentifierCall(ictx cobol85.IIdentifierContext) (ret call.Call) {
	ctx := ictx.(*cobol85.IdentifierContext)

	if cctx := ctx.QualifiedDataName(); cctx != nil {
		c := e.CreateQualifiedDataNameCall(cctx)
		ret = call.NewCallDelegate(cctx, c)
		e.AddElement(ret)
	} else if cctx := ctx.TableCall(); cctx != nil {
		c := e.CreateTableCall(cctx)
		ret = call.NewCallDelegate(cctx, c)
		e.AddElement(ret)
	} else if cctx := ctx.FunctionCall(); cctx != nil {
		c := e.CreateFunctionCall(cctx)
		ret = call.NewCallDelegate(cctx, c)
		e.AddElement(ret)
	} else if cctx := ctx.SpecialRegister(); cctx != nil {
		c := e.CreateSpecialRegisterCall(cctx)
		ret = call.NewCallDelegate(cctx, c)
		e.AddElement(ret)
	} else {
		ret = e.CreateDataDescriptionEntryCall(ctx)
	}
	return
}

func (e *program) CreateSpecialRegisterCall(ictx cobol85.ISpecialRegisterContext) (ret call.Call) {
	element := e.GetElement(ictx)
	if element != nil {
		ret, _ = element.(call.Call)
		return
	}
	ctx := ictx.(*cobol85.SpecialRegisterContext)
	var typ call.SpecialRegisterType
	switch {
	case ctx.ADDRESS() != nil:
		typ = call.ADDRESS_OF
	case ctx.DATE() != nil:
		typ = call.DATE
	case ctx.DAY() != nil:
		typ = call.DAY
	case ctx.DAY_OF_WEEK() != nil:
		typ = call.DAY_OF_WEEK
	case ctx.DEBUG_CONTENTS() != nil:
		typ = call.DEBUG_CONTENTS
	case ctx.DEBUG_ITEM() != nil:
		typ = call.DEBUG_ITEM
	case ctx.DEBUG_LINE() != nil:
		typ = call.DEBUG_LINE
	case ctx.DEBUG_NAME() != nil:
		typ = call.DEBUG_NAME
	case ctx.DEBUG_SUB_1() != nil:
		typ = call.DEBUG_SUB_1
	case ctx.DEBUG_SUB_2() != nil:
		typ = call.DEBUG_SUB_2
	case ctx.DEBUG_SUB_3() != nil:
		typ = call.DEBUG_SUB_3
	case ctx.LENGTH() != nil:
		typ = call.LENGTH_OF
	case ctx.LINAGE_COUNTER() != nil:
		typ = call.LINAGE_COUNTER
	case ctx.LINE_COUNTER() != nil:
		typ = call.LINE_COUNTER
	case ctx.PAGE_COUNTER() != nil:
		typ = call.ADDRESS_OF
	case ctx.RETURN_CODE() != nil:
		typ = call.RETURN_CODE
	case ctx.SHIFT_IN() != nil:
		typ = call.SHIFT_IN
	case ctx.SHIFT_OUT() != nil:
		typ = call.SHIFT_OUT
	case ctx.SORT_CONTROL() != nil:
		typ = call.SORT_CONTROL
	case ctx.SORT_CORE_SIZE() != nil:
		typ = call.SORT_CORE_SIZE
	case ctx.SORT_FILE_SIZE() != nil:
		typ = call.SORT_FILE_SIZE
	case ctx.SORT_MESSAGE() != nil:
		typ = call.SORT_MESSAGE
	case ctx.SORT_MODE_SIZE() != nil:
		typ = call.SORT_MODE_SIZE
	case ctx.SORT_RETURN() != nil:
		typ = call.SORT_RETURN
	case ctx.TALLY() != nil:
		typ = call.TALLY
	case ctx.TIME() != nil:
		typ = call.TIME
	case ctx.WHEN_COMPILED() != nil:
		typ = call.WHEN_COMPILED
	default:
	}

	specialRegisterCall := call.NewSpecialRegisterCall(ctx, typ)

	if cctx := ctx.Identifier(); cctx != nil {
		c := e.CreateIdentifierCall(cctx)
		specialRegisterCall.SetIdentifierCall(c)
	}
	ret = specialRegisterCall
	e.AddElement(ret)
	return
}

func (e *program) CreateFunctionCall(ictx cobol85.IFunctionCallContext) (ret call.Call) {
	element := e.GetElement(ictx)
	if element != nil {
		ret, _ = element.(call.Call)
		return
	}

	name := util.DetermineName(ictx)
	ret = call.NewFunctionCall(ictx, name)
	e.AddElement(ret)
	return
}

func (e *program) CreateTableCall(ictx cobol85.ITableCallContext) (ret call.Call) {
	element := e.GetElement(ictx)
	if element != nil {
		ret, _ = element.(call.Call)
		return
	}

	ctx := ictx.(*cobol85.TableCallContext)

	entries := e.GetDataDescriptionEntries(ctx)

	if len(entries) == 0 {
		ret = e.CreateUndefinedCall(ctx)
	} else {
		entry := entries[0]
		name := util.DetermineName(ctx)
		tableCall := call.NewTableCall(ctx, name, entry)
		for _, v := range ctx.AllSubscript() {
			var subscript *valuestmt.Subscript
			element1 := e.GetElement(v)
			if element1 == nil {
				cctx := v.(*cobol85.SubscriptContext)
				subscript = valuestmt.NewSubscript(v)
				subscriptValueStmt := e.CreateValueStmt(cctx.IntegerLiteral(), cctx.QualifiedDataName(), cctx.IndexName(), cctx.ArithmeticExpression())
				subscript.SetSubscriptValueStmt(subscriptValueStmt)
				tableCall.AddSubscript(subscript)
				e.AddElement(subscript)
			}
		}
		ret = tableCall
		entry.AddCall(ret)
		e.AddElement(ret)
	}

	return
}

func (e *program) CreateQualifiedDataNameCall(ictx cobol85.IQualifiedDataNameContext) (ret call.Call) {
	ctx := ictx.(*cobol85.QualifiedDataNameContext)

	if cctx := ctx.QualifiedDataNameFormat1(); cctx != nil {
		c := e.CreateQualifiedDataNameFormat1Call(cctx)
		ret = call.NewCallDelegate(cctx, c)
		e.AddElement(ret)
	} else if cctx := ctx.QualifiedDataNameFormat2(); cctx != nil {
		c := e.CreateQualifiedDataNameFormat2Call(cctx)
		ret = call.NewCallDelegate(cctx, c)
		e.AddElement(ret)
	} else if cctx := ctx.QualifiedDataNameFormat3(); cctx != nil {
		c := e.CreateQualifiedDataNameFormat3Call(cctx)
		ret = call.NewCallDelegate(cctx, c)
		e.AddElement(ret)
	} else if cctx := ctx.QualifiedDataNameFormat4(); cctx != nil {
		c := e.CreateQualifiedDataNameFormat4Call(cctx)
		ret = call.NewCallDelegate(cctx, c)
		e.AddElement(ret)
	} else {
		ret = e.CreateDataDescriptionEntryCall(ctx)
	}
	return
}

func (e *program) CompareQualifiedInData(ctxs []cobol85.IQualifiedInDataContext, entry datadescription.DataDescriptionEntry) (ret bool) {
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

func (e *program) CreateQualifiedDataNameFormat1Call(ictx cobol85.IQualifiedDataNameFormat1Context) (ret call.Call) {
	ctx := ictx.(*cobol85.QualifiedDataNameFormat1Context)

	datas := ctx.AllQualifiedInData()
	if len(datas) != 0 {
		name := util.DetermineName(ctx)
		var validEntry datadescription.DataDescriptionEntry
		entries := e.GetDataDescriptionEntries(ctx)
		for _, entry := range entries {
			if e.CompareQualifiedInData(datas, entry) {
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

func (e *program) CreateQualifiedDataNameFormat2Call(ictx cobol85.IQualifiedDataNameFormat2Context) (ret call.Call) {
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

func (e *program) CreateQualifiedDataNameFormat3Call(ictx cobol85.IQualifiedDataNameFormat3Context) (ret call.Call) {
	ctx := ictx.(*cobol85.QualifiedDataNameFormat3Context)
	return e.CreateCall(ctx.TextName())
}

func (e *program) CreateQualifiedDataNameFormat4Call(ictx cobol85.IQualifiedDataNameFormat4Context) (ret call.Call) {
	ctx := ictx.(*cobol85.QualifiedDataNameFormat4Context)
	return e.CreateUndefinedCall(ctx)
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

func (e *program) GetDataDescriptionEntries(ctx antlr.ParserRuleContext) (entries []datadescription.DataDescriptionEntry) {
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

func (e *program) CreateDataDescriptionEntryCall(ctx antlr.ParserRuleContext) (ret call.Call) {
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

func (e *program) CreateDataDescriptionEntryCallWithName(ctx antlr.ParserRuleContext, name string, entry datadescription.DataDescriptionEntry) (ret *call.DataDescriptionEntryCall) {
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

func (e *program) CreateUndefinedCall(ctx antlr.ParserRuleContext) (ret call.Call) {
	element := e.GetElement(ctx)
	if element != nil {
		ret, _ = element.(call.Call)
		return
	}

	name := util.DetermineName(ctx)
	ret = call.NewUndefinedCall(ctx, name)
	e.AddElement(ret)
	return
}

func (e *program) CreateIndexCall(ctx antlr.ParserRuleContext) (ret *call.IndexCall) {
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
