package ws

import (
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/medzikuser/wserver/utils"
)

const MAX_PACKET_LEN int = 65536

var upgrader = websocket.Upgrader{
	ReadBufferSize:  MAX_PACKET_LEN,
	WriteBufferSize: MAX_PACKET_LEN,
}

func Handle(w http.ResponseWriter, r *http.Request) {
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

		utils.Log.Infof("%s sent: %s", conn.RemoteAddr(), string(msg))

		// write output message to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			utils.Log.Error(err)
			return
		}

		msg_split := strings.Split(string(msg), " ")

		handleCommand(msg_split[0], msg_split[1:])
	}
}
