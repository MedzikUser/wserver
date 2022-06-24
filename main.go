package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/medzikuser/wserver/utils/plugin"
	"github.com/medzikuser/wserver/ws"
)

const listenPort = ":7567"

var pluginsDir = "plugins"

func main() {
	plugins := plugin.Load(pluginFiles(pluginsDir))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// handle websocket connections
		ws.Handle(w, r, plugins)
	})

	// listen server
	http.ListenAndServe(listenPort, nil)
}

func pluginFiles(dir string) []string {
	outputDirRead, err := os.Open(dir)
	if err != nil {
		panic(err)
	}

	outputDirFiles, err := outputDirRead.Readdir(0)
	if err != nil {
		panic(err)
	}

	var plugin_files []string

	// one by one, add the plugin path to the `plugin_files` variable
	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]
		plugin_files = append(plugin_files, fmt.Sprintf("%s/%s", dir, outputFileHere.Name()))
	}

	return plugin_files
}
