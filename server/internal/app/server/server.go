/*
	package server

	This is Server, its main function is:
	- Init server info
	- serve http request for client

	Author: DYX, ZJX

	Date: 2023/07/22
*/

package server

import (
	"encoding/json"
	"log"
	"net/http"
	data "overlaysr/server/internal/pkg/data"

	"github.com/gorilla/websocket"
)

type Server interface {
	// Server upgrades HTTP connection for clients
	Serving(w http.ResponseWriter, r *http.Request)

	// server update chan
	update(update chan data.Message)
}

type WsServer struct {
	addr     string
	message  chan data.Message
	upgrader websocket.Upgrader
	clients  map[*websocket.Conn]bool // all clients
	// TODO : we may need more info in map
}

func (ws *WsServer) Init() {
	ws.addr = "localhost:8080"
	ws.message = make(chan data.Message, 100)
	ws.upgrader = websocket.Upgrader{}
}

func (ws *WsServer) GetAddr() string {
	return ws.addr
}

func (ws *WsServer) Serving(w http.ResponseWriter, r *http.Request) {
	c, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade err: %v", err)
		return
	}

	log.Printf("Connect from %s, put into clients map", c.RemoteAddr())
	ws.clients[c] = true

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Read error: ", err)
			delete(ws.clients, c)
			log.Printf("Disconnect from %s", c.RemoteAddr())
			break
		}
		log.Printf("Recv from client: %s", message)
		reply := "Ok"
		log.Printf("Send to client: %s", reply)
		msg, err := json.Marshal(reply)
		if err != nil {
			log.Println("Marshal error: ", err)
			break
		}
		err = c.WriteMessage(mt, msg)
		if err != nil {
			log.Println("Write error: ", err)
			break
		}
	}
}
