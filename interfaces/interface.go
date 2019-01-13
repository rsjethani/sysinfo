package interfaces

type InfoProvider interface {
	Type() string
	Category() string
	Attributes() *map[string]interface{}
	Attribute(string) (interface{}, bool)
}
