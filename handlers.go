package main

import (
	r "github.com/dancannon/gorethink"
	"github.com/mitchellh/mapstructure"
)

func addChannel(client *Client, data interface{}) {
	var channel Channel
	mapstructure.Decode(data, &channel)
	// TODO: insert the new added channel in rethinkDB
	err := r.Table("channel").
		Insert(channel).
		Exec(client.session)
	if err != nil {
		client.send <- Message{"error", err.Error()}
	}
}
