package main

import (
	"fmt"
	"net/http"

	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
)

/*Handler func type*/
type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

/*Router is a struct */
type Router struct {
	rules   map[string]Handler
	session *r.Session
}

/*NewRouter construct the router object*/
func NewRouter(session *r.Session) *Router {
	return &Router{
		rules:   make(map[string]Handler),
		session: session,
	}
}

/*Handle func*/
func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

/*FindHandler is a methd to help find the proper handler*/
func (r *Router) FindHandler(msgName string) (Handler, bool) {
	handler, found := r.rules[msgName]
	return handler, found
}

/*ServeHTTP func*/
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	client := NewClient(socket, r.FindHandler, r.session)
	defer client.Close()
	go client.Write()
	client.Read()
}
