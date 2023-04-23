package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	clien :=
		flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", client.echo)
	http.HandleFunc("/", client.home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
