package main

import (
	"flag"
	"log"

	"github.com/gophergala2016/talk_to_someone_client"
)

var (
	server = flag.String("server", "localhost:8080/ws", "Server address")
)

func main() {
	flag.Parse()
	client.CreateConnection(*server, "BOT", "BOT")
	log.Println("Client active =", client.IsActive())
	for {
		message := client.GetMessage()
		log.Println("Message received =", message)
		client.SendMessage("Mmm, you're so cool )")
		log.Println("Message sended")
	}
	client.CloseConnection()
}
