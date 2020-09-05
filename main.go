package main

import "net/http"

func main() {
	// fmt.Println("Hello there, welcome to the new world!")
	http.HandleFunc("/", handler)
}
