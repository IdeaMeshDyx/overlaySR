package main

import (
	"flag"
	"log"
	"net/http"

	server "overlaysr/server/internal/app/server"
)

func main() {
	var hub server.WsServer
	hub.Init()
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", hub.Serving)
	log.Fatal(http.ListenAndServe(hub.GetAddr(), nil))
}
