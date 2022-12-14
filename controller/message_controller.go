package controller

import (
	"log"
	"net/http"

	"rescues/model"

	"github.com/gorilla/websocket"
)

// Configure the upgreader
var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan model.Message)     // broadcast channel
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type messageController struct {}

type MessageController interface {
	handelWSConnections(w http.ResponseWriter, r *http.Request)
}


func handelWSConnections(w http.ResponseWriter, r *http.Request) {
	//Upgrader initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg model.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleWSMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}


