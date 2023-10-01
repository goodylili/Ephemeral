package http

import "time"

type ChatRoom struct {
	ID        string
	Users     []User
	Messages  []Message
	CreatedAt time.Time
}
