package plugin

import (
	"fmt"
	"log"
	"plugin"

	"github.com/MedzikUser/go-plugins-loader-example/types"
)

// Load plugins
func LoadPlugins(plugins []string) []types.Plugin {
	loaded_plugins := []types.Plugin{}

	for _, plugin_path := range plugins {
		// open the `.so` plugin file
		pl, err := plugin.Open(plugin_path)
		if err != nil {
			log.Fatal(err)
		}

		// lookup for variable `PluginName`
		pName, err := pl.Lookup("PluginName")
		if err != nil {
			log.Fatal(err)
		}
		pluginName := *pName.(*string)

		// lookup for function `OnLoad`
		pLoadF, err := pl.Lookup("OnLoad")
		if err != nil {
			log.Fatal(err)
		}
		pluginLoadF := pLoadF.(func())

		// lookup for function `F`
		pF, err := pl.Lookup("F")
		if err != nil {
			log.Fatal(err)
		}
		pluginF := pF.(func())

		// execute plugin `OnLoad` function
		pluginLoadF()

		plugin := types.Plugin{
			Name:   pluginName,
			OnLoad: pluginLoadF,
			F:      pluginF,
		}

		// append the plugin to the loaded
		loaded_plugins = append(loaded_plugins, plugin)
	}

	fmt.Printf("Loaded plugins (%d):\n", len(loaded_plugins))
	for i, plugin := range loaded_plugins {
		fmt.Printf("(%d) %s\n", i+1, plugin.Name)
	}

	return loaded_plugins
}
