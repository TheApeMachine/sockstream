package sockstream

import (
	"log"

	"gocv.io/x/gocv"
)

// Hub is a wrapper around active clients and message broadcasting.
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// NewHub returns a reference to a new instance of Hub.
func NewHub(capCh chan gocv.Mat) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// Run the hub.
func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = true
			log.Println("new client registered")
		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)
			}
			log.Println("client unregistered")
		case message := <-hub.broadcast:
			for client := range hub.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(hub.clients, client)
				}
			}
		}
	}
}
