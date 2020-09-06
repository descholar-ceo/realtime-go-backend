package main

import (
	"github.com/gorilla/websocket"
)

/*FindHandler function*/
type FindHandler func(string) (Handler, bool)

/*Message is a struct*/
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

/*Client struct*/
type Client struct {
	send        chan Message
	socket      *websocket.Conn
	findHandler FindHandler
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
func (client *Client) Read() {
	var msg Message
	for {
		if err := client.socket.ReadJSON(&msg); err != nil {
			break
		}
		if handler, found := client.findHandler(msg.Name); found {
			handler(client, msg.Data)
		}
	}
	client.socket.Close()
}

/*NewClient is the instation of the Client object */
func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client {
	return &Client{
		send:        make(chan Message),
		socket:      socket,
		findHandler: findHandler,
	}
}
