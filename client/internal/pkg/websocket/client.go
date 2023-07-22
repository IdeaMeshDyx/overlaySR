package WsClient

import (
	"log"
	"net/url"
	data "overlaysr/client/internal/pkg/data"

	websocket "github.com/gorilla/websocket"
)

type client interface {
	// Init client info and connect to server
	Init()

	// Send message to server
	Send(data.Message)

	// main thread of client
	Runing()

	// Receive message from server
	// Receive()
	// NOTE: we are receiving message from server in main thread

}

type WsClient struct {
	TargetAddr string
	Wsconn     *websocket.Conn
	Exited     bool
}

func (c *WsClient) Init() {
	c.TargetAddr = "localhost:8080"
	u := url.URL{Scheme: "ws", Host: c.TargetAddr, Path: "/ws"}
	log.Printf("Trying to connect to %s", u.String())

	wsc, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial err: ", err)
	}
	c.Wsconn = wsc
	c.Exited = false
	log.Printf("Connected to %s successfully", u.String())
}

func (c *WsClient) Send(msg []byte) {
	err := c.Wsconn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("Write error: ", err)
	}
	log.Printf("Send success")
}

func (c *WsClient) Runing() {
	defer c.Wsconn.Close()
	// read thread
	for {
		if c.Exited {
			err := c.Wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Printf("close ws conn error: %v", err)
			}
			log.Println("Exited")
			return
		}
		_, message, err := c.Wsconn.ReadMessage()
		if err != nil {
			log.Println("Read error: ", err)
			return
		}
		log.Printf("Recv: %s", message)
	}
}
