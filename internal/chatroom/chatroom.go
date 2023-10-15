package chatroom

import (
	"context"
	"log"
)

type Repo interface {
	CreateChatRoom(ctx context.Context, username string, roomname string) error
	JoinChatRoom(ctx context.Context, username string, chatRoomID string) error
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

func (s *Service) CreateChatRoom(ctx context.Context, username string, roomname string) error {
	err := s.CreateChatRoom(ctx, username, roomname)
	if err != nil {
		log.Printf("Error Encountered Creating Chatroom: %v:", err)
		return err
	}
	return nil

}

func (s *Service) JoinChatRoom(ctx context.Context, username string, chatRoomID string) error {
	err := s.JoinChatRoom(ctx, username, chatRoomID)
	if err != nil {
		log.Printf("Error Encountered Joining Chatroom: %v:", err)
		return err
	}
	return nil
}
