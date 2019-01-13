package sysinfo

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"plugin"

	"github.com/rsjethani/sysinfo/builtins"
	"github.com/rsjethani/sysinfo/interfaces"
)

const configDir = ".sysinfo"

var pluginDir = filepath.Join(configDir, "plugins")

func loadExternalPlugin(category string, name string) (func() (interfaces.InfoProvider, error), error) {
	u, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("Error while locating plugins: %v", err)
	}

	pluginFile := fmt.Sprintf("sysinfo-%v-%v", category, name)
	pluginPath := filepath.Join(u.HomeDir, pluginDir, pluginFile)
	if _, e := os.Stat(pluginPath); os.IsNotExist(e) || os.IsPermission(e) {
		return nil, fmt.Errorf("The plugin file '%v' either does not exists (not implemented?) or we do not have read permission on it", pluginPath)
	}

	pluginHandle, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, fmt.Errorf("Error while loading plugin binary '%v': %v", pluginPath, err)
	}

	init, err := pluginHandle.Lookup("Init")
	if err != nil {
		return nil, fmt.Errorf("Error while looking up 'Init()' in plugin '%v': %v", name, err)
	}

	initFunc, ok := init.(func() (interfaces.InfoProvider, error))
	if !ok {
		return nil, fmt.Errorf("Error while parsing '%v' plugin's Init(), bad function signature. Required signature: %T", name, initFunc)
	}

	return initFunc, nil
}

func initPlugin(category string, name string) (interfaces.InfoProvider, error) {
	initFunc, ok := builtins.BuiltinPlugins[name]
	// If not a builtin plugin, try finding and loading an external plugin
	if !ok {
		f, err := loadExternalPlugin(category, name)
		if err != nil {
			return nil, err
		}
		initFunc = f
	}
	provider, err := initFunc()
	if err != nil {
		return nil, fmt.Errorf("Error while initializing plugin '%v': %v", name, err)
	}

	return provider, nil

}

func GetInfo(category string, name string) (interfaces.InfoProvider, error) {
	provider, err := initPlugin(category, name)
	if err != nil {
		return nil, fmt.Errorf("Cannot get information about '%v': %v", name, err)
	}
	return provider, nil
}
