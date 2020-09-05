package main

import (
	"github.com/gorilla/websocket"
)

/*Message is a struct*/
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

/*Client struct*/
type Client struct {
	send   chan Message
	socket *websocket.Conn
}

/*Write method*/
func (client *Client) Write() {
	for msg := range client.send {
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	client.socket.Close()
}

/*Read method is to read the inMessage*/
fucn (client *Client) Read(){}

/*NewClient is the instation of the Client object*/
func NewClient(socket *websocket.Conn) *Client {
	return &Client{
		send:   make(chan Message),
		socket: socket,
	}
}
