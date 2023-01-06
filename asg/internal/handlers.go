package internal

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/asg/model"
	"github.com/kaisawind/cobol/asg/model/call"
	"github.com/kaisawind/cobol/asg/model/data"
	"github.com/kaisawind/cobol/asg/model/environment"
	"github.com/kaisawind/cobol/gen/cobol85"
)

type NewDataDivisionFunc func(ctx *cobol85.DataDivisionContext, programUnit model.ProgramUnit) data.DataDivision
type NewEnvironmentDivisionFunc func(ctx *cobol85.EnvironmentDivisionContext, programUnit model.ProgramUnit) environment.EnvironmentDivision
type NewCallDelegateFunc func(ctx antlr.ParserRuleContext, delegate call.Call, programUnit model.ProgramUnit) call.Call

var (
	NewDataDivision        NewDataDivisionFunc
	NewEnvironmentDivision NewEnvironmentDivisionFunc
	NewCallDelegate        NewCallDelegateFunc
)

func RegisterNewDataDivisionFunc(f NewDataDivisionFunc) {
	NewDataDivision = f
}

func RegisterNewEnvironmentDivisionFunc(f NewEnvironmentDivisionFunc) {
	NewEnvironmentDivision = f
}

func RegisterNewCallDelegateFunc(f NewCallDelegateFunc) {
	NewCallDelegate = f
}
