package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// /*Message is a struct*/
// type Message struct {
// 	Name string      `json:"name"`
// 	Data interface{} `json:"data"`
// }

// /*Channel is a struct type*/
// type Channel struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

func main() {
	recRawMsg := []byte(`{"name":"channel add", "data":{"name":"Hardware support"}}`)

	var recMsg Message
	err := json.Unmarshal(recRawMsg, &recMsg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", recMsg)
	if recMsg.Name == "channel add" {
		channel, err := addChannel(recMsg.Data)
		var sendMsg Message
		sendMsg.Name = "channel add"
		sendMsg.Data = channel
		sendRawMsg, err := json.Marshal(sendMsg)
		if err != nil {
			return
		}
		fmt.Println(string(sendRawMsg))
	}
}

func addChannel(data interface{}) (Channel, error) {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return channel, err
	}
	channel.ID = "1"

	fmt.Printf("the channel is : %#v\n", channel)
	return channel, nil
}
