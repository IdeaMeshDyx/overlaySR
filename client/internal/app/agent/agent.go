package agent

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	data "overlaysr/client/internal/pkg/data"
	"time"

	"github.com/gorilla/websocket"
)

type agent interface {
	/**TODO:
	* Client 获取 Message 之后直接发送
	 */
	Send() error

	/**TODO:
	接收server 的更新信息并放入 update chan
	*/
	Update(update chan data.Message)

	/**TODO:
	* 接收collector 的更新信息并放入 update chan
	 */
	Read(update chan data.Message)
}

type WsAgent struct {
	Addr    string
	Message data.Message
}

func (agent *WsAgent) Read(cilium chan data.Message) {
	for msg := range cilium {
		agent.Message = msg
	}
}

func (agent *WsAgent) AddrUp() {
	agent.Addr = *flag.String("Addr", "localhost:8080", "http service Address")
}

func (agent *WsAgent) Send() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: agent.Addr, Path: "/echo"}
	log.Printf("Connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			log.Printf("Sending tick message : %s \n", t.String())
			wmsg, _ := agent.Message.Byte()
			err := c.WriteMessage(websocket.TextMessage, wmsg)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
