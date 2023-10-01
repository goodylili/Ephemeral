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
	User1     User
	User2     User
	Messages  []Message
	CreatedAt time.Time
}

func (c Client) createOrJoinChatRoom(ctx context.Context, room ChatRoom) (string, error) {
	// Check if the room ID is provided
	if room.ID == "" {
		room.ID = uuid.New().String()
	} else {
		// Check if a chatroom with the given ID already exists in Redis
		exists, err := c.Client.Exists(ctx, room.ID).Result()
		if err != nil {
			log.Printf("Failed to check existence of chat room: %v", err)
			return "", err
		}
		if exists > 0 {
			// The room already exists, so return the room ID for "joining"
			return room.ID, nil
		}
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
