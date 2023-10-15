package http

import (
	"Ephemeral/internal/chatroom"
	"io/ioutil"
	"net/http"
)

func handleJoinChatRoomForm(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "This form is invalid", http.StatusBadRequest)
	}
	username := request.FormValue("username")
	chatroomId := request.FormValue("chatroomId")

	if username == "" || chatroomId == "" {
		http.Error(writer, "This form is invalid", http.StatusBadRequest)
	}

	_, err = ioutil.ReadFile("../../static/html/join.html") // Assuming data.txt is the file you want to read.
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	err = chatroom.JoinChatRoom(request.Context(), username, chatroomId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func handleCreateChatRoomForm(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "This form is invalid", http.StatusBadRequest)
	}

	username := request.FormValue("username")
	roomname := request.FormValue("roomname")

	_, err = ioutil.ReadFile("../../static/html/login.html") // Assuming data.txt is the file you want to read.
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	if username == "" || roomname == "" {
		http.Error(writer, "This form is invalid", http.StatusBadRequest)
	}

	err = chatroom.CreateChatRoom(request.Context(), username, roomname)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
