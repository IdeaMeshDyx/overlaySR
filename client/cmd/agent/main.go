/*
package main

This is agent. Its functions are as follows:
- Init agent and client
- Receive message from collector
- Send message to server through websocket client

Author: DYX, ZJX

Date: 2023/07/22
*/

package main

import (
	"log"
	"os"
	"os/signal"
	"overlaysr/client/internal/app/agent"
	"overlaysr/client/internal/app/collector"
	"overlaysr/client/internal/pkg/data"
	"time"
)

func main() {
	// instantiate a collector and a websocket agent
	ws_buffer := make(chan data.Message, 10)
	defer close(ws_buffer)
	var agent agent.Agent
	agent.Init()
	coll := collector.Collector{
		CollID: "collector1",
		Msg:    data.WsMsg{},
		Exited: false,
	}

	// collector write data into this chan
	go coll.Collect(ws_buffer)

	go agent.Runing()
	// main func

	// agent read and send collector's data in this chan
	go agent.ReadAndSend(ws_buffer)

	done := make(chan struct{})
	defer close(done)
	interrupt := make(chan os.Signal, 1)
	defer signal.Stop(interrupt)
	signal.Notify(interrupt, os.Interrupt)
	// main thread to exit
	for {
		select {
		case <-done:
			log.Println("agent done")
			return
		case <-interrupt:
			log.Println("keyboard interrupt")
			agent.WsClient.Exited = true
			coll.Exited = true
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}
