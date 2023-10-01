package http

import (
	"context"
	"time"
)

type ChatRoom struct {
	ID        string    `json:"ID"`
	User1     User      `json:"User1"`
	User2     User      `json:"User2"`
	Messages  []Message `json:"Messages"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type ChatRoomRepo interface {
	CreateChatRoom(ctx context.Context, room ChatRoom) (string, error)
	DeleteChatRoom(ctx context.Context, chatRoomID string) error
}

// ChatRoomService is the blueprint for the chatroom logic
type ChatRoomService struct {
	Repo ChatRoomRepo
}

// NewChatRoomService creates a new service
func NewChatRoomService(repo ChatRoomRepo) ChatRoomService {
	return ChatRoomService{
		Repo: repo,
	}
}
