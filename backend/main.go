/**
 * @author qzylalala
 * @github qzylalala
 * @date 2021/10/21 20:42
 */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const BUFFER = 1024

var upgrader = websocket.Upgrader{
		ReadBufferSize: BUFFER,
		WriteBufferSize: BUFFER,
		
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Panic(err)
			return
		}

		fmt.Println(string(p))

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Panic(err)
			return
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panic(err)
	}

	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
