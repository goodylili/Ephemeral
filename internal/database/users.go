package database

import (
	"github.com/gorilla/websocket"
)

type User struct {
	ID         string
	Username   string
	ChatRoomID string
	WebSocket  *websocket.Conn
}
