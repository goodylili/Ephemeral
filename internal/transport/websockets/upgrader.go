package websockets

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func Upgrader() websocket.Upgrader {
	upgradeInstance := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return upgradeInstance

}
