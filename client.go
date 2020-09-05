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

/*Write method*/
func (client *Client) Write() {
	for msg := range client.send {
		// TODO: call socket.sendJSON(msg)
		fmt.Printf("%#v\n", msg)
	}
}

/*NewClient is the instation of the Client object*/
func NewClient() *Client {
	return &Client{
		send: make(chan Message),
	}
}
