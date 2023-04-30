package main

import (
	"fmt"
	"overlaysr/client/internal/app/agent"
	"overlaysr/client/internal/app/collector"
	"overlaysr/client/internal/pkg/data"
)

func main() {
	// instantiate a collector and a websocket agent
	buffer := make(chan data.Message, 10)
	var hub agent.WsAgent
	coll := collector.Collector{
		CollID: "collector1",
		Msg:    data.WsMsg{},
	}

	// start two goroutines to collect data
	go coll.Collect(buffer)
	go hub.Read(buffer)
	for msg := range buffer {
		data, _ := msg.Byte()
		fmt.Printf("From: %s:\nData:%s\n", coll.CollID, data)
		go hub.Send()
	}
	// TODO:
	// create a new websocket agent
	// send a message using the agent
}
