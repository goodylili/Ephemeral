package messages

import (
	"Ephemeral/internal/user"
	"context"
	"time"
)

type Message struct {
	Content    string    `json:"content"`
	User       user.User `json:"user"`
	ChatRoomID string    `json:"chatRoomID"`
	Time       time.Time `json:"time"`
}

type MessageRepo interface {
	AddMessageToChatRoom(ctx context.Context, msg Message) error
	GetMessages(ctx context.Context, chatRoomID string) ([]Message, error)
}

type MessageService struct {
	Repo MessageRepo
}

// NewMessageService creates a new service
func NewMessageService(repo MessageRepo) MessageService {
	return MessageService{
		Repo: repo,
	}
}
