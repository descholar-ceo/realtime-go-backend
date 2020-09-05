package main

import (
	"net/http"
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
	router := &Router{}
	// http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
