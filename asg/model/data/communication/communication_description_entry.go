package communication

type CDEType int

const (
	INPUT CDEType = iota
	INPUT_OUTPUT
	OUTPUT
)

type CommunicationDescriptionEntry interface {
	Type() CDEType
	Name() string
}
