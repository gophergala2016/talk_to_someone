package client

import (
	"log"
	"strings"
	//"time"

	"github.com/gorilla/websocket"
	"github.com/gophergala2016/talk_to_someone/server/utils"
)

var (
	connect *websocket.Conn
	active  = false
	userID  string
)

func CreateConnection(url, id string) bool {
	log.Println("Connecting to %s ...", url)
	userID = id
	var err error
	connect, _, err = websocket.DefaultDialer.Dial("ws://"+url, nil)
	if err != nil {
		log.Println("Can't connect to server because of :", err)
		return active
	}
	log.Println("Connection established")
	active = true
	return active
}

func CloseConnection() {
	active = false
	log.Println("Connection closed")
	connect.Close()
}

func Setup(name string) bool {
	result := false
	err := connect.WriteMessage(websocket.TextMessage, []byte("READY:"+userID+":"+name))
	if err != nil {
		log.Println("Can't send a message because of:", err)
		return result
	}
	text := GetMessage()
	log.Println("Received message :", text)
	if strings.Contains(text, "ACCEPT:"+userID+":READY") {
		result = true
	}
	return result
}

func SendMessage(message string) bool {
	result := false
	log.Println("Client active =", active)
	if active {
		str := "MESSAGE:" + userID + ":" + message
		//connect.SetWriteDeadline(time.Now().Add(time.Second * 30))
		err := connect.WriteMessage(websocket.TextMessage, []byte(str))
		if err != nil {
			log.Println("Can't send a message because of:", err)
			active = false
			CloseConnection()
		}
		result = true
	}
	return result
}

func GetMessage() string {
	var result string
	if active {
		//connect.SetReadDeadline(time.Now().Add(time.Second * 1))
		mtype, message, err := connect.ReadMessage()
		if err != nil {
			log.Println("Can't receive a message because of:", err)
			active = false
			CloseConnection()
		}
		switch mtype {
		case websocket.TextMessage:
			result = string(message)
			if utils.StringStartWith(result, "MESSAGE:") {
				pos := strings.LastIndex(result, ":")
				result = string(result[pos + 1:])
			}
		case websocket.PingMessage:
			connect.WriteMessage(websocket.PongMessage, []byte(""))
		case websocket.CloseMessage:
			CloseConnection()
		}
	}
	return result
}

func FinishSession() {
	//connect.SetWriteDeadline(time.Now().Add(time.Second * 1))
	connect.WriteMessage(websocket.TextMessage, []byte("FINISH:"+userID))
}

func IsActive() bool {
	return active
}
