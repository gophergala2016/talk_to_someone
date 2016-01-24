package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/gophergala2016/talk_to_someone/server/utils"
	"github.com/gorilla/websocket"
)

type User struct {
	id           string
	name         string
	counterparty string
	connect      *websocket.Conn
	messageBox   chan []byte
}

func (user *User) read() {
	defer func() {
		user.connect.Close()

	}()
	for {
		//user.connect.SetReadDeadline(time.Now().Add(time.Second * 30))
		mtype, message, err := user.connect.ReadMessage()
		if err != nil {
			log.Println("Can't receive a message because of:", err)
			break
		}
		if mtype == websocket.TextMessage {
			text := string(message)
			if utils.StringStartWith(text, "READY:") {
				id, name := utils.GetIdAndName(text)
				if user.id == "" || user.name == "" {
					user.id = id
					user.name = name
					Save(user.id, user)
				}
			}
			if utils.StringStartWith(text, "GETPERSON:") {
				//id := utils.GetId(text)
			}
			if utils.StringStartWith(text, "MESSAGE:") {
				sender := utils.GetId(text)
				user, exist := Get(sender)
				if exist {
					user2, exist := Get(user.counterparty)
					if exist {
						user2.messageBox <- []byte(text)
					}
				}
			}
		}

		Hub.broadcast <- message
	}
}

func (user *User) write() {
	ticker := time.NewTicker(time.Millisecond * 200)
	for {
		select {
		case message, ok := <-user.messageBox:
			if !ok {
				user.connect.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := user.connect.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := user.connect.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
	defer func() {
		ticker.Stop()
		user.connect.Close()
	}()
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Handler(resp http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Println("Can't upgrade websocket because of:", err)
	}
	conn.SetPongHandler(func(string) error {
		return nil
	})
	user := &User{
		connect:    conn,
		messageBox: make(chan []byte, 4096),
	}
	go user.write()
	go user.read()
}
