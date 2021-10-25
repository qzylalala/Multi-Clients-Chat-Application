/**
 * @author qzylalala
 * @github qzylalala
 * @date 2021/10/21 20:42
 */

package main

import (
	"fmt"
	"github.com/qzylalala/AChat-application/pkg/websocket"
	"net/http"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}



func main() {
	fmt.Println("Multi-Clients Chat App v0.1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
