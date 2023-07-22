/*
package main

This is Server's main file, its main function is:
- Init server info
- Listen http request for client

Author: DYX, ZJX

Date: 2023/07/22
*/
package main

import (
	"fmt"
	"log"
	"net/http"

	server "overlaysr/server/internal/app/server"
)

func main() {
	var hub server.WsServer
	hub.Init()
	fmt.Printf("Listening on %s\n", hub.GetAddr())
	// err := http.ListenAndServe(hub.GetAddr(), nil)
	http.HandleFunc("/ws", hub.Serving)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
