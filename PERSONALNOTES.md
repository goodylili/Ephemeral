If you're creating a chat app with Go using Redis and Gorilla WebSockets and want the chatroom and messages to expire in 15 minutes, there are some functionalities you might consider. First, let's start with the basic user struct:

```go
type User struct {
    ID         string // Unique ID for the user
    Username   string
    ChatRoomID string // This ensures a user is only in one chat room
    WebSocket  *websocket.Conn
}
```

Here are some potential functionalities and the corresponding functions:

1. **User Management**:
   - `CreateUser(username string) (*User, error)`: Create a new user.
   - `GetUserByID(id string) (*User, error)`: Retrieve a user by their ID.
   - `DeleteUser(id string) error`: Delete a user.

3. **Message Management**:
   - `SendMessage(userID, messageContent string) error`: Send a message in the chat room. The message should be stored in Redis and set to expire after 15 minutes.
   - `FetchMessages(chatRoomID string) ([]string, error)`: Fetch all non-expired messages for a chat room.

4. **WebSocket Management**:
   - `StartWebSocketConnection(userID string) error`: Start a WebSocket connection for a user.
   - `CloseWebSocketConnection(userID string) error`: Close a WebSocket connection for a user.
   - `BroadcastMessageToChatRoom(chatRoomID, messageContent string) error`: Send a message to all users in a chat room via WebSocket.

5. **Utility Functions**:
   - `IsUserInAnyChatRoom(userID string) (bool, error)`: Check if a user is already in a chat room.
   - `ExpireChatRoom(chatRoomID string) error`: This is mainly handled by Redis, but if you need to perform any application-level cleanup, this function can help.

6. **Error Handling**:
   - Make sure to handle potential errors appropriately, e.g., when a user tries to join a chat room that doesn’t exist or if they try to join a chat room while already in another one.

7. **Clean Up**:
   - Periodically, you might want to clean up stale chat rooms or messages (those beyond the 15-minute mark). Even though you've set them to expire in Redis, it's good to have a mechanism to ensure data consistency.

Remember to also handle concurrency appropriately, especially when users are trying to join or leave chat rooms. Go's channels, mutexes, or other synchronization primitives will be invaluable here. Given that Redis operations and WebSocket broadcasts are involved, you'll want to ensure that your operations are atomic where necessary.