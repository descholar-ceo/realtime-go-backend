package main

import (
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
)

/*FindHandler function*/
type FindHandler func(string) (Handler, bool)

/*Message is a struct */
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

/*Client struct */
type Client struct {
	send         chan Message
	socket       *websocket.Conn
	findHandler  FindHandler
	session      *r.Session
	stopChannels map[int]chan bool
}

/*NewStopChannel is a func which is in charge of stopping a goroutine*/
func (client *Client) NewStopChannel(stopKey int) chan bool {
	stop := make(chan bool)
	client.stopChannels[stopKey] = stop
	return stop
}

/*StopForKey is a method incharge of stopping */
func (client *Client) StopForKey(key int) {
	if mChannel, found := client.stopChannels[key]; found {
		mChannel <- true
		delete(client.stopChannels, key)
	}
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

/*Read method is to read the inMessage */
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

/*Close method cleses the client connection*/
func (client *Client) Close() {
	for _, mChannel := range client.stopChannels {
		mChannel <- true
	}
	close(client.send)
}

/*NewClient is the instation of the Client object */
func NewClient(socket *websocket.Conn, findHandler FindHandler, session *r.Session) *Client {
	return &Client{
		send:         make(chan Message),
		socket:       socket,
		findHandler:  findHandler,
		session:      session,
		stopChannels: make(map[int]chan bool),
	}
}
