package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data/datadescription"
	"github.com/kaisawind/cobol/asg/util"
	"github.com/kaisawind/cobol/gen/cobol85"
)

func GetParent[T ElementType](program *Program, ctx antlr.Tree) (parent T) {
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

func CreateCall(program *Program, ctx antlr.ParserRuleContext) (ret call.Call) {
	switch t := ctx.(type) {
	case cobol85.IDataDescNameContext:
		CreateDataDescriptionEntryCall(program, t)
	}
	return
}

func GetDataDescriptionEntries(program *Program, ctx antlr.ParserRuleContext) (entries []datadescription.DataDescriptionEntry) {
	name := util.DetermineName(ctx)
	programUnit := GetParent[*ProgramUnit](program, ctx)
	if programUnit == nil || programUnit.dataDivision == nil {
		return
	}
	dataDivision := programUnit.dataDivision

	workingStorageSection := dataDivision.GetWorkingStorageSection()
	communicationSection := dataDivision.GetCommunicationSection()
	return
}

func CreateDataDescriptionEntryCall(program *Program, ctx antlr.ParserRuleContext) (ret call.Call) {
	element := program.GetElement(ctx)
	if element != nil {
		ret, _ = element.(call.Call)
		return
	}

	index := GetIndex(program, ctx)

	if index != nil {
		ret = CreateIndexCall(program, ctx)
	} else {

	}
	return
}

func CreateIndexCall(program *Program, ctx antlr.ParserRuleContext) (ret *call.IndexCall) {
	element := program.GetElement(ctx)
	if element != nil {
		ret, _ = element.(*call.IndexCall)
		return
	}

	name := util.DetermineName(ctx)
	index := GetIndex(program, ctx)

	if index == nil {
		CreateDataDescriptionEntryCall(program, ctx)
	} else {
		ret = call.NewIndexCall(ctx, name, index)
		// link call and index
		index.AddIndexCall(ret)
		program.AddElement(ret)
	}
	return
}

func GetIndex(program *Program, ctx antlr.ParserRuleContext) (ret *datadescription.Index) {
	name := util.DetermineName(ctx)
	programUnit := GetParent[*ProgramUnit](program, ctx)
	if programUnit == nil || programUnit.dataDivision == nil {
		return
	}
	workingStorageSection := programUnit.dataDivision.GetWorkingStorageSection()
	if workingStorageSection == nil {
		return
	}

	for _, entry := range workingStorageSection.GetCommunicationDescriptionEntries() {
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
