package main

import (
	"fmt"
	"overlaysr/client/internal/app/agent"
	"overlaysr/client/internal/app/collector"
	"overlaysr/client/internal/pkg/data"
)

func main() {
	// instantiate a collector and a websocket agent
	ws_buffer := make(chan data.Message, 10)
	var hub agent.WsAgent
	hub.AddrUp()
	coll := collector.Collector{
		CollID: "collector1",
		Msg:    data.WsMsg{},
	}

	// start two goroutines to collect data
	go coll.Collect(ws_buffer)
	go hub.Read(ws_buffer)
	// TODO : now we are using a websocket conn per message. we should optimize this more efficiently
	// TODO2 : seperate the agent and the collector into several goroutines/threads/processes
	for {
		select {
		case <-ws_buffer:
			fmt.Printf("Collector: %s:\nReturns Data, tring to send through websocket\n", coll.CollID)
			go hub.Send()
		}
	}
	// TODO:
	// create a new websocket agent
	// send a message using the agent
}
