Your idea of a time-based chat application where chat rooms auto-destruct is interesting and quite unique! Given the tech stack you've chosen (Golang and GoRedis), you have a strong foundation to start. Redis, with its support for auto-deleting keys (TTL), is particularly suited to your requirements.

Here's a basic outline of what you need to do:

1. **Backend Setup:**
    - Set up a Golang HTTP or WebSocket server.
    - Connect your Golang server to Redis using GoRedis.

2. **Creating a Chat Room:**
    - When a chat room is created, generate a unique chat room ID.
    - Store this chat room ID in Redis with a TTL of 900 seconds (15 minutes).

3. **Sending and Storing Messages:**
    - When a message is sent, append it to a list in Redis associated with the chat room ID.
    - Ensure the TTL of the chat room is reset or maintained with each message, unless you want individual messages to expire separately.

4. **Destroying Chat Rooms:**
    - Provide a mechanism to manually delete a chat room (remove the key from Redis).

5. **Fetching Messages:**
    - When a user joins or sends a message, fetch the messages from the Redis list associated with the chat room ID.

6. **Frontend:**
    - Create a basic frontend using your preferred framework or library.
    - Connect to your Golang server to send and receive messages.
    - Display messages in real-time and inform the user of the chat room's impending auto-destruction.

Here's a basic example for the backend using Golang and GoRedis:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

var redisClient *redis.Client
var ctx = context.Background()

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func createChatRoom() string {
	chatRoomID := uuid.New().String()
	err := redisClient.Set(ctx, chatRoomID, "", 15*time.Minute).Err()
	if err != nil {
		log.Fatalf("Failed to create chat room: %v", err)
	}
	return chatRoomID
}

func storeMessage(chatRoomID string, message string) {
	redisClient.RPush(ctx, chatRoomID+":messages", message)
	// Renew the TTL
	redisClient.Expire(ctx, chatRoomID, 15*time.Minute)
}

func getMessages(chatRoomID string) []string {
	messages, err := redisClient.LRange(ctx, chatRoomID+":messages", 0, -1).Result()
	if err != nil {
		log.Fatalf("Failed to fetch messages: %v", err)
	}
	return messages
}

func chatHandler(ws *websocket.Conn) {
	// Placeholder WebSocket handler
	// Implement your chat logic here
}

func main() {
	http.Handle("/chat", websocket.Handler(chatHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

This is a very basic structure and you'll need to add more functionality, error handling, and features. Also, for the sake of this example, I've used the `websocket` package for handling WebSockets. Depending on your requirements, you might choose to go with a more robust package or framework for WebSocket handling.

Remember to secure your application, especially when using WebSockets. Also, consider rate limiting, authentication, authorization, and other standard security practices.