package websocket

import (
	"encoding/json"
	"fmt"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan JSONMessage

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

// NewHub init new hub instance
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		Broadcast:  make(chan JSONMessage),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Run Run
func (h *Hub) Run() {
	for {
		select {

		case client := <-h.Register:
			h.clients[client] = true

		case client := <-h.Unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		case message := <-h.Broadcast:
			data := Message{}
			if err := json.Unmarshal(message, &data); err != nil {
				// log.Printf("error: %v \n", err)
				fmt.Printf("error: %v \n", err)
				break
			}

			for client := range h.clients {
				if client.ID.ID == data.To.ID ||
					hasIntersected(client.ID.Groups, data.To.Groups) {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}

// Send 发送消息
func (h *Hub) Send(msg Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	h.Broadcast <- JSONMessage(data)
	return nil
}
