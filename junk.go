package main

import (
	"encoding/json"
	"fmt"
)

/*Message is a struct*/
type Message struct {
	Name string
	Data interface{}
}

func main() {
	recRawMsg := []byte(`{"name":"channel add",` +
		`"data":{"name":"Hardware support"}`)

	var recMsg Message
	err := json.Unmarshal(recRawMsg, &recMsg)
	if err != nil {
		fmt.Println(err)
		return
	}
}
