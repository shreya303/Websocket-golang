package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v\n", err)
		return
	}
	defer ws.Close()

	log.Println("Client connected")

	for {
		var msg string
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			break
		}

		log.Printf("Received: %s\n", msg)

		err = ws.WriteJSON(fmt.Sprintf("Echo: %s", msg))
		if err != nil {
			log.Printf("Error writing message: %v\n", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	port := ":8080"
	log.Printf("WebSocket server starting on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed: %v\n", err)
	}
}
