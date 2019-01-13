package builtins

import (
	csv "encoding/csv"
	"os"
	"strings"

	"github.com/rsjethani/sysinfo/interfaces"
)

// battery represents the information about the system battery
type battery struct {
	devType     string
	category    string
	atttributes *map[string]interface{}
}

func (b *battery) Category() string {
	return b.category
}

func (b *battery) Attribute(attr string) (interface{}, bool) {
	val, ok := (*b.atttributes)[attr]
	return val, ok
}

func (b *battery) Attributes() *map[string]interface{} {
	return b.atttributes
}

func (b *battery) Type() string {
	return b.devType
}

func BatteryInit() (interfaces.InfoProvider, error) {
	items := make(map[string]interface{})
	b := battery{devType: "battery", category: "hardware", atttributes: &items}

	f, err := os.Open("/sys/class/power_supply/BAT0/uevent")
	defer f.Close()
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)
	r.Comma = 61 //ASCII value of "="
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, rec := range records {
		items[strings.Split(rec[0], "POWER_SUPPLY_")[1]] = rec[1]
	}
	b.atttributes = &items
	return &b, nil
}
