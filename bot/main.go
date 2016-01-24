package main

import (
	"flag"
	"log"

	"github.com/artemnikitin/talk_to_someone_client"
)

var (
	server = flag.String("server", "localhost:8080/ws", "Server address")
)

func main() {
	flag.Parse()
	client.CreateConnection(*server, "BOT")
	client.Setup("BOT")
	log.Println("Client active =", client.IsActive())
	for client.IsActive() {
		message := client.GetMessage()
		log.Println("Message received :", message)
		if message == "" {
			break
		}
		client.SendMessage("Mmm, you're so cool )")
		log.Println("Message sended")
	}
	defer client.CloseConnection()
}
