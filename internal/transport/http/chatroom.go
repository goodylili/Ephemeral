package http

import (
	"io/ioutil"
	"net/http"
)

func handleJoinChatRoomForm(writer http.ResponseWriter, request *http.Request) (string, string, error) {
	err := request.ParseForm()
	if err != nil {
		return "", "", err
	}
	username := request.FormValue("username")
	chatroomId := request.FormValue("chatroomId")

	_, err = ioutil.ReadFile("../../static/html/join.html") // Assuming data.txt is the file you want to read.
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return "", "", err
	}

	return username, chatroomId, nil

}

func handleCreateChatRoomForm(writer http.ResponseWriter, request *http.Request) (string, string, error) {
	err := request.ParseForm()
	if err != nil {
		return "", "", err
	}
	username := request.FormValue("username")
	roomname := request.FormValue("roomname")

	_, err = ioutil.ReadFile("../../static/html/login.html") // Assuming data.txt is the file you want to read.
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return "", "", err
	}
	return username, roomname, nil
}
