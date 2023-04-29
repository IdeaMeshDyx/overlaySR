package test

import (
	"fmt"
	"overlaysr/client/internal/app/agent"
	"overlaysr/client/internal/app/collector"
	"overlaysr/client/internal/pkg/data"
)

func main() {
	// create a new websocket agent
	buffer := make(chan data.Message, 10)
	var hub agent.WsAgent
	var coll collector.Collector
	go coll.Collect(buffer)
	go hub.Read(buffer)
	// msg := hub.Message
	fmt.Printf("%v", hub.Message)
	// send a message using the agent
	// fmt.Println(hub.Message)
}
