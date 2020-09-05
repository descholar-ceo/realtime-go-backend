package main

import (
	"fmt"
)

/*Message is a struct*/
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

/*Client struct*/
type Client struct {
	send chan Message
}

func (client *Client) write() {
	for msg := range client.send {
		// TODO: call socket.sendJSON(msg)
		fmt.Println("%#v\n", msg)
	}
}

func (client *Client) subscribeChannel() {
	// TODO: changefeed Query RethinkDB
}

func main() {}
