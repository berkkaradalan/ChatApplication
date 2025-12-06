package websocket

import "sync"

type Message struct {
	RoomID string
	Data   []byte
}

type Hub struct {
	// Map of roomID to map of clients
	rooms map[string]map[*Client]bool

	// Channel for broadcasting messages to a specific room
	broadcast chan *Message

	// Channel for registering clients
	register chan *Client

	// Channel for unregistering clients
	unregister chan *Client

	// Mutex for thread-safe operations
	mu sync.RWMutex
}

var (
	hubInstance *Hub
	once        sync.Once
)

// GetHub returns the singleton hub instance
func GetHub() *Hub {
	once.Do(func() {
		hubInstance = &Hub{
			broadcast:  make(chan *Message, 256),
			register:   make(chan *Client),
			unregister: make(chan *Client),
			rooms:      make(map[string]map[*Client]bool),
		}
		go hubInstance.run()
	})
	return hubInstance
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			if h.rooms[client.roomID] == nil {
				h.rooms[client.roomID] = make(map[*Client]bool)
			}
			h.rooms[client.roomID][client] = true
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if clients, ok := h.rooms[client.roomID]; ok {
				if _, exists := clients[client]; exists {
					delete(clients, client)
					close(client.send)
					// Clean up empty rooms
					if len(clients) == 0 {
						delete(h.rooms, client.roomID)
					}
				}
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.RLock()
			clients := h.rooms[message.RoomID]
			h.mu.RUnlock()

			for client := range clients {
				select {
				case client.send <- message.Data:
				default:
					h.mu.Lock()
					close(client.send)
					delete(h.rooms[message.RoomID], client)
					h.mu.Unlock()
				}
			}
		}
	}
}

// BroadcastToRoom sends a message to all clients in a specific room
func (h *Hub) BroadcastToRoom(roomID string, data []byte) {
	h.broadcast <- &Message{
		RoomID: roomID,
		Data:   data,
	}
}