package chatroom

import (
	"Ephemeral/internal/messages"
	"Ephemeral/internal/user"

	"context"
	"time"
)

type ChatRoom struct {
	ID        string           `json:"ID"`
	User1     user.User        `json:"User1"`
	User2     user.User        `json:"User2"`
	Message   messages.Message `json:"Message"`
	CreatedAt time.Time        `json:"CreatedAt"`
}

type Repo interface {
	CreateChatRoom(ctx context.Context, room ChatRoom) (string, error)
	DeleteChatRoom(ctx context.Context, chatRoomID string) error
	JoinChatRoom(ctx context.Context, roomID string) (string, error)
}

// Service is the blueprint for the chatroom logic
type Service struct {
	Repo Repo
}

// NewChatRoomService creates a new service
func NewChatRoomService(repo Repo) Service {
	return Service{
		Repo: repo,
	}
}

// CreateChatRoom creates a new chatroom
func (s Service) CreateChatRoom(ctx context.Context, room ChatRoom) (string, error) {
	return s.Repo.CreateChatRoom(ctx, room)
}

// DeleteChatRoom deletes a chatroom
func (s Service) DeleteChatRoom(ctx context.Context, chatRoomID string) error {
	return s.Repo.DeleteChatRoom(ctx, chatRoomID)
}

// JoinChatRoom joins a chatroom
func (s Service) JoinChatRoom(ctx context.Context, roomID string) (string, error) {
	return s.Repo.JoinChatRoom(ctx, roomID)
}
