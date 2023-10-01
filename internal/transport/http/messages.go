package http

import "time"

type Message struct {
	Content    string    `json:"content"`
	User       User      `json:"user"`
	ChatRoomID string    `json:"chatRoomID"`
	Time       time.Time `json:"time"`
}
