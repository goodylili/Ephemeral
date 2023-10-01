package database

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

type Message struct {
	Content    string
	User       User
	ChatRoomID string
	Time       time.Time
}

func (c Client) AddMessageToChatRoom(ctx context.Context, msg Message) error {
	// Fetch the chat room data from Redis using the chat room's ID.
	data, err := c.Client.Get(ctx, msg.ChatRoomID).Result()
	if err != nil {
		log.Printf("Failed to fetch chat room with ID %s from Redis: %v", msg.ChatRoomID, err)
		return err
	}

	// Deserialize the data back into a ChatRoom struct.
	var room ChatRoom
	err = json.Unmarshal([]byte(data), &room)
	if err != nil {
		log.Printf("Failed to deserialize chat room data: %v", err)
		return err
	}

	// Append the new message to the chat room's Messages slice.
	room.Messages = append(room.Messages, msg)

	// Serialize the updated ChatRoom struct.
	updatedData, err := json.Marshal(room)
	if err != nil {
		log.Printf("Failed to serialize updated chat room: %v", err)
		return err
	}

	// Save the updated data back to Redis. Note that this will overwrite the previous chat room data.
	// We also maintain the 15-minute expiration for consistency.
	err = c.Client.Set(ctx, msg.ChatRoomID, updatedData, 15*time.Minute).Err()
	if err != nil {
		log.Printf("Failed to update chat room in Redis: %v", err)
		return err
	}

	return nil
}

func (c Client) GetMessages(ctx context.Context, chatRoomID string) ([]Message, error) {
	// Fetch the chat room data from Redis using the chat room's ID.
	data, err := c.Client.Get(ctx, chatRoomID).Result()
	if err != nil {
		log.Printf("Failed to fetch chat room with ID %s from Redis: %v", chatRoomID, err)
		return nil, err
	}

	// Deserialize the data back into a ChatRoom struct.
	var room ChatRoom
	err = json.Unmarshal([]byte(data), &room)
	if err != nil {
		log.Printf("Failed to deserialize chat room data: %v", err)
		return nil, err
	}

	return room.Messages, nil
}
