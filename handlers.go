package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func addChannel(client *Client, data interface{}) {
	var channel Channel
	var msg Message
	mapstructure.Decode(data, &channel)
	fmt.Printf("%#v\n", channel)
	// TODO: insert the new added channel in rethinkDB
	channel.ID = "ABC123"
	msg.Name = "channel add"
	msg.Data = channel
	client.send <- msg
}
