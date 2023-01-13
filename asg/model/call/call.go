package call

type CallType int

const (
	COMMUNICATION_DESCRIPTION_ENTRY_CALL CallType = iota
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
	Type() CallType
	Name() string
}
