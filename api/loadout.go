package api

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/PlayBlockiro/CraftWorks/utils"
)

type Plugin struct {
	Name string
	Path string
}

func LoadPlugins() ([]Plugin, error) {
	pluginDir := "plugins"
	entries, err := os.ReadDir(pluginDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read plugin directory: %v", err)
	}

	var plugins []Plugin
	for _, entry := range entries {
		if entry.IsDir() {
			pluginPath := filepath.Join(pluginDir, entry.Name(), "main.go")
			if _, err := os.Stat(pluginPath); err == nil {
				plugins = append(plugins, Plugin{Name: entry.Name(), Path: pluginPath})
			}
		}
	}

	return plugins, nil
}

func RunPlugin(plugin Plugin) error {
	utils.LogInfo(fmt.Sprintf("Running plugin: %s", plugin.Name))

	cmd := exec.Command("go", "run", plugin.Path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
