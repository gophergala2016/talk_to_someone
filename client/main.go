package client

import (
	"log"
	"time"
	"strings"

	"github.com/gorilla/websocket"
)

var (
	connect *websocket.Conn
	active bool = false
	userID string
)

func CreateConnection(url, name, id string) bool {
	log.Println("Connecting to %s ...", url)
	userID = id
	var err error
	connect, _, err = websocket.DefaultDialer.Dial("ws://" + url, nil)
	if err != nil {
		log.Println("Can't connect to server because of :", err)
		return active
	}
	log.Println("Connection established")
	result := setup(name)
	if result {
		log.Println("Initialization finished")
		active = true
	}
	return active
}

func CloseConnection()  {
	active = false
	log.Println("Connection closed")
	connect.Close()
}

func setup(name string) bool {
	var result bool = false
	err := connect.WriteMessage(websocket.TextMessage, []byte("READY:" + userID + ":" + name))
	if err != nil {
		log.Println("Can't send a message because of:", err)
		return result
	}
	text := GetMessage()
	log.Println("Received message ||", text)
	if strings.Contains(text, "ACCEPT:" + userID + ":READY") {
		result = true
	}
	return result
}

func SendMessage(message string) bool {
	var result bool = false
	log.Println("Client active =", active)
	if active {
		str := "MESSAGE:" + userID + ":" + message
		connect.SetWriteDeadline(time.Now().Add(time.Second * 1))
		err := connect.WriteMessage(websocket.TextMessage, []byte(str))
		if err != nil {
			log.Println("Can't send a message because of:", err)
		}
		result = true
	}
	return result
}

func GetMessage() string {
	var result string
	if active {
		connect.SetReadDeadline(time.Now().Add(time.Second * 1))
		mtype, message, err := connect.ReadMessage()
		if err != nil {
			log.Println("Can't receive a message because of:", err)
		}
		if mtype == websocket.TextMessage {
			result = string(message)
		}
		if mtype == websocket.PingMessage {
			connect.WriteMessage(websocket.PongMessage, []byte(""))
		}
	}
	return result
}

func FinishSession() {
	connect.WriteMessage(websocket.TextMessage, []byte("FINISH:" + userID))
}

func IsActive() bool {
	return active
}
