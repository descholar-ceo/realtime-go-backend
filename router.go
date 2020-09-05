package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Handler func(*Client,interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

/*Router is a struct*/
type Router struct{}

func (r *Router) Handle(msgName string, Handler){}

func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
}
