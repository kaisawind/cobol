package datadescription

type DDEType int

const (
	CONDITION DDEType = iota
	EXEC_SQL
	GROUP
	RENAME
	SCALAR
)

type DataDescriptionEntry interface {
	Type() DDEType
	Name() string
}
