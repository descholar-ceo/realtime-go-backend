package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

/*Message is a struct*/
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

/*Channel is a struct type*/
type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello there, I am coming from the go server")
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// msgType, msg, err := socket.ReadMessage()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		var inMsg, outMsg Message
		if err := socket.ReadJSON(&inMsg); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", inMsg)
		switch inMsg.Name {
		case "channel add":
			err := addChannel(inMsg.Data)
			if err != nil {
				outMsg = Message{"error", err}
				if err := socket.WriteJSON(outMsg); err != nil {
					fmt.Println(err)
					break
				}
			}
		case "channel subscribe":
			subscribeChannel()
		}
		// fmt.Printf("The messageType is: %v\nThe messsage is: %v\n", int(msgType), string(msg))
		// if err = socket.WriteMessage(msgType, msg); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}
}

func addChannel(data interface{}) error {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.ID = "1"

	fmt.Printf("the channel is : %#v\n", channel)
	return nil
}

func subscribeChannel() {
	// TODO: rethink Query / changefeed
}
