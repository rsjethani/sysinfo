package  interfaces

type InfoProvider interface {
	Type() string
	String() string
	Category() string
	Attributes() *map[string]interface{}
	Attribute(string) (interface{}, error)
}

