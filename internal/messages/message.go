package messages

import (
	"context"
	"time"
)

type Message struct {
	MessageID  string
	ChatRoomID string
	Sender     string
	Content    string
	Timestamp  time.Time
}

type Repo interface {
	AddMessage(ctx context.Context, chatRoomID, sender, content string) (*Message, error)
	FetchAllMessages(ctx context.Context, chatRoomID string) ([]Message, error)
}

// Service is the blueprint for the chatroom logic
type Service struct {
	Repo Repo
}

// NewMessageService creates a new service
func NewMessageService(repo Repo) Service {
	return Service{
		Repo: repo,
	}
}

// AddMessage adds a new message to the chatroom
func (s Service) AddMessage(ctx context.Context, chatRoomID, sender, content string) (*Message, error) {
	return s.Repo.AddMessage(ctx, chatRoomID, sender, content)
}

// FetchAllMessages fetches all messages from the chatroom
func (s Service) FetchAllMessages(ctx context.Context, chatRoomID string) ([]Message, error) {
	return s.Repo.FetchAllMessages(ctx, chatRoomID)
}
