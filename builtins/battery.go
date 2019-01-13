package builtins

import "github.com/rsjethani/sysinfo/interfaces"

// battery represents the information about the system battery
type battery struct {
	initialized bool
	category    string
	devType     string
	atttributes *map[string]interface{}
	// Name             string
	// Status           string
	// Present          string
	// Technology       string
	// CycleCount       string
	// VoltageMinDesign uint
	// VoltageNow       uint
	// PowerNow         uint
	// EnergyFullDesign uint
	// EnergyFull       uint
	// EnergyNow        uint
	// Capacity         uint
	// CapacityLevel    string
	// ModelName        string
	// Manufacturer     string
	// SerialNumber     string
}

func (*battery) Type() string {
	return "battery"
}

func (*battery) Category() string {
	return "hardware"
}

func (b *battery) Attribute(attr string) (interface{}, error) {
	return (*b.atttributes)[attr], nil
}

func (b *battery) Attributes() *map[string]interface{} {
	return b.atttributes
}

func (b *battery) String() string {
	return "To be implemented"
}

func BatteryInit() (interfaces.InfoProvider, error) {
	b := battery{}
	b.initialized = true
	x := make(map[string]interface{})
	b.atttributes = &x
	(*b.atttributes)["name"] = "sfsfsdf"
	return &b, nil
}
