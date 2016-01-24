package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/artemnikitin/talk_to_someone/ws"
)

var (
	port = flag.String("port", "8080", "Set port for server")
)

func main() {
	log.Println("Server started")
	go ws.Hub.Run()
	http.HandleFunc("/ws", ws.Handler)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
