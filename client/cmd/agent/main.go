package main

import (
	"flag"
	agent "overlaysr/client/internal/app/agent"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

// main is the entry point of the program
func main() {
	// create a new websocket agent
	var hub agent.WsAgent

	// send a message using the agent
	hub.Send()
}
