package websocket

import (
	"log"
	"os"
	"sync"

	error_utils "github.com/MauricioMilano/stock_app/utils/error"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Error      chan string
	Clients    map[*Client]bool
	Broadcast  chan Message
}

type RoomMap struct {
	Pools map[string]*Pool
	mu    sync.RWMutex
}

func (r *RoomMap) GetPool(name string) *Pool {
	r.mu.RLock()
	pool := r.Pools[name]
	r.mu.RUnlock()
	if pool == nil {
		r.Pools[name] = NewPool()
		r.mu.Lock()
		pool = r.Pools[name]
		r.mu.Unlock()
		go pool.Start()
		return pool
	}
	return pool
}

func NewRoomMap() *RoomMap {
	rooms := &RoomMap{
		Pools: make(map[string]*Pool),
	}
	return rooms
}
func (r *RoomMap) NewRoom(name string) {
	r.Pools[name] = NewPool()

}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (p *Pool) Start() {
	defer p.ReviveWebsocket()
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			log.Println("info:", "New client. Size of connection pool:", len(p.Clients))

		case client := <-p.Unregister:
			log.Println("info:", "disconnected a client. size of connection pool:", len(p.Clients))
			delete(p.Clients, client)

		case msg := <-p.Broadcast:
			log.Println("info", "broadcast message to clients in pool")
			for c := range p.Clients {
				err := c.Connection.WriteJSON(msg)
				error_utils.ErrorCheck(err)
			}
		case errorMessage := <-p.Error:
			log.Println("error", errorMessage)
		}

	}
}

func (p *Pool) ReviveWebsocket() {
	if err := recover(); err != nil {
		if os.Getenv("LOG_PANIC_TRACE") == "true" {
			log.Println(
				"level:", "error",
				"err: ", err,
				// "trace", string(debug.Stack()),
			)
		} else {
			log.Println(
				"level", "error",
				"err", err,
			)
		}
		go p.Start()
	}
}
