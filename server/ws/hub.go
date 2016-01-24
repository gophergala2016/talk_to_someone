package ws

import "github.com/gophergala2016/talk_to_someone/server/utils"

type hub struct {
	broadcast chan []byte
}

var Hub = hub{
	broadcast: make(chan []byte),
}

func (hub *hub) Run() {
	for {
		message := <-hub.broadcast
		processMessage(string(message))
	}
}

func processMessage(message string) {
	if utils.StringStartWith(message, "READY:") {
		id, _ := utils.GetIdAndName(message)
		user, exist := Get(id)
		if exist {
			user.messageBox <- []byte("ACCEPT:" + id + ":READY")
		}
	}
	if utils.StringStartWith(message, "GETPERSON:") {
		id := utils.GetId(message)
		user, exist := Get(id)
		if exist {
			user2, exist := Get(user.counterparty)
			if exist {
				user.messageBox <- []byte("FOUNDPERSON:" + user2.name)
			}
		}
	}
	if utils.StringStartWith(message, "MESSAGE:") {
		id := utils.GetId(message)
		user, exist := Get(id)
		if exist {
			target, exist := Get(user.counterparty)
			if exist {
				target.messageBox <- []byte(message)
			}
		}
	}
}
