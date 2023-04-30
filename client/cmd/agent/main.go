package main

import (
	"flag"
	agent "overlaysr/client/internal/app/agent"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	// create a new websocket agent
	// cil :=
	var hub agent.WsAgent
	// msg := hub.Message
	// fmt.Print(msg.Byte())
	// send a message using the agent
	hub.Send()
}
