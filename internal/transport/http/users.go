package http

import "github.com/gorilla/websocket"

type User struct {
	Username string          `json:"username"`
	Conn     *websocket.Conn `json:"conn"`
}
