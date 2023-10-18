package http

import (
	"Ephemeral/internal/chatroom"
	"Ephemeral/internal/messages"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Handler struct {
	Router   *mux.Router
	Server   *http.Server
	ChatRoom *chatroom.Service
	Message  *messages.Service
}

func (h *Handler) mapRoutes() {
	
}

func NewHandler(chatRoom *chatroom.Service, message *messages.Service) *Handler {
	log.Println("Setting Up Handlers...")
	h := &Handler{
		ChatRoom: chatRoom,
		Message:  message,
	}
	h.Router = mux.NewRouter()

	h.mapRoutes()

	h.Server = &http.Server{
		Addr:         "0.0.0.0:8080", // Good practice to set timeouts to avoid Slow-loris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.Router,
	}
	return h
}

// Serve - gracefully serves our newly set up handler function
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	<-c

	// CreateAccount a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := h.Server.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("shutting down gracefully")
	return nil
}
