package builtins

import "github.com/rsjethani/sysinfo/interfaces"

var BuiltinPlugins = map[string]func() (interfaces.InfoProvider, error){
	"battery": BatteryInit,
}
