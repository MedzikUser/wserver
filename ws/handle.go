package ws

import (
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/medzikuser/wserver/utils"
	"github.com/medzikuser/wserver/utils/plugin"
)

const MAX_PACKET_LEN int = 65536

var upgrader = websocket.Upgrader{
	ReadBufferSize:  MAX_PACKET_LEN,
	WriteBufferSize: MAX_PACKET_LEN,
}

func Handle(w http.ResponseWriter, r *http.Request, plugins []plugin.Plugin) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.Log.Error(err)
		w.Write([]byte(err.Error()))
		return
	}

	for {
		// read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			utils.Log.Error(err)
			return
		}

		msg_str := string(msg)
		msg_str = strings.ReplaceAll(msg_str, "\n", "")
		msg_str = strings.ReplaceAll(msg_str, "\r", "")

		utils.Log.Infof("%s sent: %s", conn.RemoteAddr(), msg_str)

		msg_split := strings.Split(msg_str, " ")

		handleCommand(msgType, conn, msg_split[0], msg_split[1:], plugins)
	}
}
