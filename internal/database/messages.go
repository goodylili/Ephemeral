package database

import "time"

type Message struct {
	Content  string
	Time     time.Time
	User     *User
	ChatRoom uint
}
