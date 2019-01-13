package interfaces

type InfoProvider interface {
	Type() string
	Category() string
	Attributes() []*map[string]interface{}
	Attribute(index uint, attr string) (interface{}, error)
}
