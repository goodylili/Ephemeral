package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Message struct {
	MessageID  string
	ChatRoomID string
	Sender     string
	Content    string
	Timestamp  time.Time
}

func (c Client) AddMessage(ctx context.Context, chatRoomID, sender, content string) (*Message, error) {
	// create a new message object
	msg := &Message{
		MessageID:  fmt.Sprintf("%s:%v", chatRoomID, time.Now().UnixNano()), // A simple unique ID using timestamp
		ChatRoomID: chatRoomID,
		Sender:     sender,
		Content:    content,
		Timestamp:  time.Now(),
	}
	// Marshal the message i	nto JSON
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	// Store the message in a sorted set for the given chatRoomID with the timestamp as the score
	if _, err := c.Client.ZAdd(ctx, "messages:"+chatRoomID, &redis.Z{
		Score:  float64(msg.Timestamp.Unix()),
		Member: msgJSON,
	}).Result(); err != nil {
		return nil, err
	}
	// Store the message as a separate key with a 15-minute expiration to track its lifetime
	msgKey := fmt.Sprintf("message:%s", msg.MessageID)
	if _, err := c.Client.SetEX(ctx, msgKey, msgJSON, 15*time.Minute).Result(); err != nil {
		return nil, err
	}
	return msg, nil
}

func (c Client) FetchAllMessages(ctx context.Context, chatRoomID string) ([]Message, error) {
	// Fetch all messages from the sorted set
	results, err := c.Client.ZRange(ctx, "messages:"+chatRoomID, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	messages := make([]Message, 0, len(results))
	for _, result := range results {
		var msg Message
		if err := json.Unmarshal([]byte(result), &msg); err != nil {
			return nil, err
		}
		// Check if the message's individual key exists (i.e., it hasn't expired)
		msgKey := fmt.Sprintf("message:%s", msg.MessageID)
		exists, err := c.Client.Exists(ctx, msgKey).Result()
		if err != nil {
			return nil, err
		}
		// If the message hasn't expired, add it to the results
		if exists == 1 {
			messages = append(messages, msg)
		} else {
			// Remove the expired message from the sorted set
			if _, err := c.Client.ZRem(ctx, "messages:"+chatRoomID, result).Result(); err != nil {
				return nil, err
			}
		}
	}
	return messages, nil
}
