/**
 * @author qzylalala
 * @github qzylalala
 * @date 2021/10/24 15:34
 */

package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const BUFFER = 1024

var upgrader = websocket.Upgrader{
	ReadBufferSize: BUFFER,
	WriteBufferSize: BUFFER,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ws, nil
}