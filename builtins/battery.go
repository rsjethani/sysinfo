package builtins

import (
	csv "encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rsjethani/sysinfo/interfaces"
)

// battery represents the information about the system battery
type battery struct {
	devType     string
	category    string
	atttributes []*map[string]interface{}
}

func (b *battery) Category() string {
	return b.category
}

func (b *battery) Attribute(index uint, attr string) (interface{}, error) {
	count := uint(len(b.atttributes))
	if index >= count {
		return nil, fmt.Errorf("Error while locating index. Given index ('%v') exceeds total no. of available objects ('%v')", index, count)
	}
	val, ok := (*b.atttributes[index])[attr]
	if !ok {
		return nil, fmt.Errorf("Attribute '%v' not found", attr)
	}
	return val, nil
}

func (b *battery) Attributes() []*map[string]interface{} {
	return b.atttributes
}

func (b *battery) Type() string {
	return b.devType
}

func BatteryInit() (interfaces.InfoProvider, error) {
	b := battery{devType: "battery", category: "hardware"}
	files, _ := filepath.Glob("/sys/class/power_supply/BAT*/uevent")
	if len(files) == 0 {
		return nil, fmt.Errorf("Error locating battery info. Either battery not present or not enough permissions")
	}
	for _, file := range files {
		f, err := os.Open(file)
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

		m := make(map[string]interface{})
		for _, rec := range records {
			key := strings.Split(rec[0], "POWER_SUPPLY_")[1]
			// convert values like capacity, charge etc to uint64
			m[key], err = strconv.ParseUint(rec[1], 10, 64)
			// Error means the value is alphanumeric value like model no., mannufacturer etc., keep these as strings
			if err != nil {
				m[key] = rec[1]
			}
		}

		b.atttributes = append(b.atttributes, &m)
	}
	return &b, nil
}
