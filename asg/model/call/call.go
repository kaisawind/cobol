package call

import "github.com/kaisawind/cobol/asg/model"

type CallType int

const (
	DEFAULT_CALL CallType = iota
	COMMUNICATION_DESCRIPTION_ENTRY_CALL
	DATA_DESCRIPTION_ENTRY_CALL
	ENVIRONMENT_CALL
	FILE_CONTROL_ENTRY_CALL
	FUNCTION_CALL
	INDEX_CALL
	MNEMONIC_CALL
	PROCEDURE_CALL
	REPORT_DESCRIPTION_CALL
	REPORT_DESCRIPTION_ENTRY_CALL
	SCREEN_DESCRIPTION_ENTRY_CALL
	SECTION_CALL
	SPECIAL_REGISTER_CALL
	TABLE_CALL
	UNDEFINED_CALL
)

type Call interface {
	model.CobolDivisionElement

	Name() string
	Type() CallType
	Unwrap() Call
}
