package plugin

import (
	"plugin"

	"github.com/gorilla/websocket"
	"github.com/medzikuser/wserver/utils"
)

// Function that loads plugins from `.so` files
func Load(files []string) []Plugin {
	plugins := []Plugin{}

	for _, plugin_file := range files {
		// open the `.so` plugin file
		plugin, err := plugin.Open(plugin_file)
		if err != nil {
			utils.Log.Fatal(err)
		}

		// lookup `PluginName`
		symbol, err := plugin.Lookup("PluginName")
		if err != nil {
			utils.Log.Fatal(err)
		}

		name := *symbol.(*string)

		// lookup `Command`
		symbol, err = plugin.Lookup("Command")
		if err != nil {
			utils.Log.Fatal(err)
		}

		command := *symbol.(*string)

		// lookup `HelpMessage`
		symbol, err = plugin.Lookup("HelpMessage")
		if err != nil {
			utils.Log.Fatal(err)
		}

		help := *symbol.(*string)

		// lookup `F` (main function)
		symbol, err = plugin.Lookup("F")
		if err != nil {
			utils.Log.Fatal(err)
		}

		F := symbol.(func(args []string, msgType int, conn *websocket.Conn))

		plugins = append(plugins, Plugin{
			Name:        name,
			Command:     command,
			HelpMessage: help,
			F:           F,
		})
	}

	if len(plugins) == 0 {
		utils.Log.Info("No plugins has been loaded.")
	} else {
		utils.Log.Infof("Loaded plugins (%d):", len(plugins))

		for i, plugin := range plugins {
			utils.Log.Infof("(%d) %s", i+1, plugin.Name)
		}
	}

	return plugins
}
