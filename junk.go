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

/*Channel is a struct type*/
type Channel struct {
	ID   string
	Name string
}

func main() {
	recRawMsg := []byte(`{"name":"channel add", "data":{"name":"Hardware support"}}`)

	var recMsg Message
	err := json.Unmarshal(recRawMsg, &recMsg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", recMsg)
	if recMsg.Name == "channel add" {
		addChannel(recMsg.Data)
	}
}

func addChannel(data interface{}) (Channel, error) {

}
