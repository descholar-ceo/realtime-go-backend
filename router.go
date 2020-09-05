package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

/*Handler func type*/
type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

/*Router is a struct*/
type Router struct {
	rules map[string]Handler
}

/*NewRouter construct the router object*/
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

/*Handle func*/
func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

/*ServeHTTP func*/
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
}
