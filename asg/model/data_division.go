package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model/data"
	"github.com/kaisawind/cobol/asg/model/data/communication"
	"github.com/kaisawind/cobol/asg/model/data/file"
	"github.com/kaisawind/cobol/asg/model/data/linkage"
	"github.com/kaisawind/cobol/asg/model/data/localstorage"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type DataDivision struct {
	ctx                   *cobol85.DataDivisionContext
	communicationSection  *communication.CommunicationSection
	workingStorageSection *data.WorkingStorageSection
	fileSection           *file.FileSection
	localStorageSection   *localstorage.LocalStorageSection
	linkageSection        *linkage.LinkageSection
}

func NewDataDivision(ctx *cobol85.DataDivisionContext) *DataDivision {
	return &DataDivision{
		ctx: ctx,
	}
}

func (e *DataDivision) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *DataDivision) SetCommunicationSection(communicationSection *communication.CommunicationSection) {
	e.communicationSection = communicationSection
}

func (e *DataDivision) GetCommunicationSection() *communication.CommunicationSection {
	return e.communicationSection
}

func (e *DataDivision) SetWorkingStorageSection(workingStorageSection *data.WorkingStorageSection) {
	e.workingStorageSection = workingStorageSection
}

func (e *DataDivision) GetWorkingStorageSection() *data.WorkingStorageSection {
	return e.workingStorageSection
}

func (e *DataDivision) SetFileSection(fileSection *file.FileSection) {
	e.fileSection = fileSection
}

func (e *DataDivision) GetFileSection() *file.FileSection {
	return e.fileSection
}

func (e *DataDivision) SetLocalStorageSection(localStorageSection *localstorage.LocalStorageSection) {
	e.localStorageSection = localStorageSection
}

func (e *DataDivision) GetLocalStorageSection() *localstorage.LocalStorageSection {
	return e.localStorageSection
}

func (e *DataDivision) SetLinkageSection(linkageSection *linkage.LinkageSection) {
	e.linkageSection = linkageSection
}

func (e *DataDivision) GetLinkageSection() *linkage.LinkageSection {
	return e.linkageSection
}
