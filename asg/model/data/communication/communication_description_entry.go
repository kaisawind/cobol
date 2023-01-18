package communication

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type CDEType int

const (
	INPUT CDEType = iota
	INPUT_OUTPUT
	OUTPUT
)

type CommunicationDescriptionEntry interface {
	Type() CDEType
	Name() string

	// AddCall add *call.CommunicationDescriptionEntryCall
	AddCall(call any)
	// GetCalls []*call.CommunicationDescriptionEntryCall
	GetCalls() []any
}

type BaseCommunicationDescriptionEntry struct {
	ctx   antlr.ParserRuleContext
	name  string
	calls []any
}

func NewBaseCommunicationDescriptionEntry(ctx antlr.ParserRuleContext, name string) *BaseCommunicationDescriptionEntry {
	return &BaseCommunicationDescriptionEntry{
		ctx:  ctx,
		name: name,
	}
}

func (e *BaseCommunicationDescriptionEntry) Context() antlr.ParserRuleContext {
	return e.ctx
}

func (e *BaseCommunicationDescriptionEntry) Name() string {
	return e.name
}

func (e *BaseCommunicationDescriptionEntry) AddCall(call any) {
	e.calls = append(e.calls, call)
}

func (e *BaseCommunicationDescriptionEntry) GetCalls() []any {
	return e.calls
}
