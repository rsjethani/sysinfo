package sysinfo

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"plugin"
)

const configDir = ".sysinfo"

var pluginDir = filepath.Join(configDir, "plugins")

func initPlugin(category string, name string) (InfoProvider, error) {
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

	provider, err := init.(func() (InfoProvider, error))()
	if err != nil {
		return nil, fmt.Errorf("Error while initializing plugin '%v': %v", name, err)
	}

	return provider, nil
}

func HwInfo(name string) (InfoProvider, error) {
	provider, err := initPlugin("hardware", name)
	if err != nil {
		return nil, fmt.Errorf("Cannot get information about '%v': %v", name, err)
	}
	return provider, nil
}

func SwInfo(name string) (InfoProvider, error) {
	provider, err := initPlugin("software", name)
	if err != nil {
		return nil, fmt.Errorf("Cannot get information about '%v': %v", name, err)
	}
	return provider, nil
}
