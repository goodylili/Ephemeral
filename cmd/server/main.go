package main

import (
	"Ephemeral/internal/chatroom"
	"Ephemeral/internal/database"
	"Ephemeral/internal/messages"
	"Ephemeral/internal/transport/http"
	"fmt"
	"log"
)

// Run - is going to be responsible for / the instantiation and startup of our / go application
func Run() error {
	fmt.Println("starting up the application...")

	redisClient, err := database.NewRedisInstance()
	if err != nil {
		log.Fatalln("Database Connection Failure")
		return err
	}

	chatRoomService := chatroom.NewChatRoomService(redisClient)

	messageService := messages.NewMessageService(redisClient)

	handler := http.NewHandler(&chatRoomService, &messageService)

	if err := handler.Serve(); err != nil {
		log.Println("failed to gracefully serve our application")
		return err
	}
	return nil

}

func main() {
	if err := Run(); err != nil {
		log.Println(err)
	}
}
