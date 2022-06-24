package ws

import (
	"github.com/gorilla/websocket"
	"github.com/medzikuser/wserver/utils/plugin"
)

func handleCommand(msgType int, conn *websocket.Conn, cmd string, args []string, plugins []plugin.Plugin) {
	for _, plugin := range plugins {
		if cmd == plugin.Command {
			plugin.F(args, msgType, conn)

			// don't look for this command in others plugins
			break
		}
	}
}
