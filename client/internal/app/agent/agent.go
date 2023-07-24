package agent

import (
	"log"
	data "overlaysr/client/internal/pkg/data"
	ws "overlaysr/client/internal/pkg/websocket"
)

type agent interface {
	// Init agent info
	Init()

	// main thread of agent
	Runing()

	// Read from collector
	ReadAndSend(update chan data.Message)
}

type Agent struct {
	Name     string
	Message  data.Message
	WsClient ws.WsClient
	// TODO: add more clients
}

func (agent *Agent) ReadAndSend(cilium chan data.Message) {
	for msg := range cilium {
		agent.Message = msg
		msg, err := agent.Message.Byte()
		if err != nil {
			log.Printf("Agent Byte error: %s\n", err)
			continue
		}
		go agent.WsClient.Send(msg)
	}
}

func (agent *Agent) Init() {
	agent.Name = "overlayAgent"
	agent.WsClient.Init()
}

func (agent *Agent) Runing() {
	// TODO: add more clients to run as go routine
	go agent.WsClient.Runing()

	// for {
	// 	time.Sleep(time.Second)
	// }
}
