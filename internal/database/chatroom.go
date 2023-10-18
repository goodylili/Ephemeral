package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type ChatRoom struct {
	ChatRoomName string
	ChatRoomID   string
	Users        []string
	CreatedAt    time.Time
}

func (c Client) CreateChatRoom(ctx context.Context, username string, roomname string) error {
	chatRoomID := fmt.Sprintf("%d", time.Now().UnixNano()) // or any other unique ID mechanism
	chatRoom := ChatRoom{
		ChatRoomName: roomname,
		ChatRoomID:   chatRoomID,
		Users:        []string{username},
		CreatedAt:    time.Now(),
	}

	// Serialize ChatRoom and store it
	data, err := json.Marshal(chatRoom)
	if err != nil {
		return err
	}

	// Set the chatroom with a 15-minute expiration
	err = c.Client.Set(ctx, "chatroom:"+chatRoomID, data, 15*time.Minute).Err()
	return err
}

func (c Client) JoinChatRoom(ctx context.Context, username string, chatRoomID string) error {
	data, err := c.Client.Get(ctx, "chatroom:"+chatRoomID).Result()
	if err != nil {
		return err
	}

	var chatRoom ChatRoom
	err = json.Unmarshal([]byte(data), &chatRoom)
	if err != nil {
		return err
	}

	// Add user to the chat room
	chatRoom.Users = append(chatRoom.Users, username)

	// Serialize back and store
	newData, err := json.Marshal(chatRoom)
	if err != nil {
		return err
	}

	// Update the chatroom (this resets the 15-minute expiration, which may or may not be what you want)
	err = c.Client.Set(ctx, "chatroom:"+chatRoomID, newData, 15*time.Minute).Err()
	return err
}
