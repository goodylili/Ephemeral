package http

import (
	"context"
	"github.com/gorilla/websocket"
)

type User struct {
	Username string          `json:"username"`
	Conn     *websocket.Conn `json:"conn"`
}

type UserRepo interface {
	CreateChatRoom(ctx context.Context, room ChatRoom) (string, error)
	DeleteChatRoom(ctx context.Context, chatRoomID string) error
}

// UserService is the blueprint for the chatroom logic
type UserService struct {
	Repo ChatRoomRepo
}

// NewUserService creates a new service
func NewUserService(repo ChatRoomRepo) ChatRoomService {
	return ChatRoomService{
		Repo: repo,
	}
}
