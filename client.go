package main

import (
	"fmt"
	"math/rand"
	"time"
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

func (client *Client) subscribeChannels() {
	// TODO: changefeed Query RethinkDB

	for {
		time.Sleep(r())
		client.send <- Message{"channel add", ""}
	}
}
func (client *Client) subscribeMessages() {
	// TODO: changefeed Query RethinkDB

	for {
		time.Sleep(r())
		client.send <- Message{"message add", ""}
	}
}

func r() time.Duration {
	return time.Millisecond * time.Duration(rand.Intn(1000))
}

func main() {}
