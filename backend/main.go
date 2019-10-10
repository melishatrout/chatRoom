package main

import (
	"fmt"
	"github.com/melishatrout/chatRoom/backend/pkg/websocket"
	"net/http"
)

func serveWs(w http.ResponseWriter, r *http.Request) {

	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprint(w, "%+V\n", err)
	}

	go websocket.Writer(ws)
	websocket.Reader(ws)

}

func setUpRoutes() {
	http.HandleFunc("/ws", serveWs)

}

func main() {
	fmt.Println("Chat App")
	setUpRoutes()
	http.ListenAndServe(":8080", nil)

}
