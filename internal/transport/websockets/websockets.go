package websockets

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Upgrades() websocket.Upgrader {
	upgrade := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return upgrade

}

func WebSocketEndpoint(writer http.ResponseWriter, request *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	upgrade := Upgrades()
	ws, err := upgrade.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected...")
	//err = ws.WriteMessage(1, []byte("Hi Client!"))
	//if err != nil {
	//	log.Println(err)
	//}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	ReadMessages(ws)
}

// ReadMessages define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func ReadMessages(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}
