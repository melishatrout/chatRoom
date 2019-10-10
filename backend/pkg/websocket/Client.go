package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID string
	Conn *websocket.Conn
	Pool *sync.Pool
}

type Message struct {
	Type int `json:"type"`
	Body string `json:"body"`
}


