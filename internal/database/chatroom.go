package database

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"time"
)

type ChatRoom struct {
	ID        string
	Users     []*User
	Messages  []Message
	CreatedAt time.Time
}

// Modified function that takes an initialized ChatRoom instance.
func (c Client) createChatRoom(ctx context.Context, room ChatRoom) (string, error) {
	// Ensure the ChatRoom instance has a unique ID. If not, assign one.
	if room.ID == "" {
		room.ID = uuid.New().String()
	}

	// Serialize the ChatRoom instance to a JSON string
	data, err := json.Marshal(room)
	if err != nil {
		log.Printf("Failed to serialize chat room: %v", err)
		return "", err
	}

	// Store the serialized data in Redis with a 15-minute expiration
	err = c.Client.Set(ctx, room.ID, data, 15*time.Minute).Err()
	if err != nil {
		log.Printf("Failed to create chat room in Redis: %v", err)
		return "", err
	}

	return room.ID, nil
}

// A function to delete a chatroom immediately based on its ID
func (c Client) deleteChatRoom(ctx context.Context, chatRoomID string) error {
	_, err := c.Client.Del(ctx, chatRoomID).Result()
	if err != nil {
		log.Printf("Failed to delete chat room with ID %s: %v", chatRoomID, err)
		return err
	}
	return nil
}
