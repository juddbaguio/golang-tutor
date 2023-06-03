package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lesismal/nbio/nbhttp/websocket"
)

func (sh *SampleHandler) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.NewUpgrader()
	upgrader.OnMessage(func(c *websocket.Conn, mt websocket.MessageType, b []byte) {
		var personalInfo PersonWithJSONTag
		err := json.Unmarshal(b, &personalInfo)

		if err != nil {
			log.Println(err)
		}

		log.Println(personalInfo)

		c.WriteMessage(mt, []byte("You sent me this message: "+string(b)))
	})

	upgrader.OnOpen(func(c *websocket.Conn) {
		log.Println("websocket connected", c.RemoteAddr().String())
	})

	upgrader.OnClose(func(c *websocket.Conn, err error) {
		log.Println("I am closing this connection")
		c.Close()
	})

	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
	}

}
