package main

import (
	"flag"
	agent "overlaysr/client/internal/app/agent"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	var hub agent.WsAgent
	hub.Send()
}
