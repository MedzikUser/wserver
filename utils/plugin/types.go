package plugin

import "github.com/gorilla/websocket"

type Plugin struct {
	Name        string
	Command     string
	HelpMessage string
	F           func(args []string, msgType int, conn *websocket.Conn)
}
