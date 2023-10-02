package user

import (
	"Ephemeral/internal/chatroom"
	"context"
	"github.com/gorilla/websocket"
)

type User struct {
	Username string          `json:"username"`
	Conn     *websocket.Conn `json:"conn"`
}

type Repo interface {
	CreateChatRoom(ctx context.Context, room chatroom.ChatRoom) (string, error)
	DeleteChatRoom(ctx context.Context, chatRoomID string) error
}

// Service  is the blueprint for the chatroom logic
type Service struct {
	Repo Repo
}

// NewUserService creates a new service
func NewUserService(repo Repo) Service {
	return Service{
		Repo: repo,
	}
}
