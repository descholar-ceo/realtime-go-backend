package main

import (
	"fmt"
	"net/http"
)

func main() {
	// fmt.Println("Hello there, welcome to the new world!")
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there, I am coming from the go server")
}
