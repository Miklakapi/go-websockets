package websocket

import (
	"fmt"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan Message
}

type Message struct {
	Sender *Client
	Data   []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.RegisterNewClient(client)
		case client := <-h.unregister:
			h.RemoveClient(client)
		case message := <-h.broadcast:
			h.HandleMessage(message)
		}
	}
}

func (h *Hub) RegisterNewClient(client *Client) {
	h.clients[client] = true
	fmt.Println("Size of clients: ", len(h.clients))
}

func (h *Hub) RemoveClient(client *Client) {
	delete(h.clients, client)
	client.Close()
	fmt.Println("Size of clients: ", len(h.clients))
}

func (h *Hub) HandleMessage(message Message) {
	for client := range h.clients {
		if client == message.Sender {
			continue
		}

		select {
		case client.send <- message:
		default:
			delete(h.clients, client)
			client.Close()
		}
	}
}
