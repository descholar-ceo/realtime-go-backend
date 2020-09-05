package main

import (
	"fmt"
	"net/http"
	"time"

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

func main() {
	// http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

func addChannel(data interface{}) error {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.ID = "1"

	// fmt.Printf("the channel is : %#v\n", channel)
	fmt.Println("Added new channel")
	return nil
}

func subscribeChannel(socket *websocket.Conn) {
	// TODO: rethink Query / changefeed
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Added a new channel")
		message := Message{"channel add", Channel{"1", "Software support"}}
		socket.WriteJSON(message)
	}
}
