package main

import (
	"github.com/gorilla/websocket"
	"github.com/medzikuser/wserver/utils"
)

// Name of the plugin
var PluginName = "test"

// Plugin Command
var Command = "/test"

// Help message of the command
var HelpMessage = "Test command from plugin"

// Main function of the plugin must be named `F`
func F(args []string, msgType int, conn *websocket.Conn) {
	// write output message to browser
	if err := conn.WriteMessage(msgType, []byte("test message")); err != nil {
		utils.Log.Error(err)
		return
	}
}
