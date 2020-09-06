package main

import (
	"fmt"

	r "github.com/dancannon/gorethink"
	"github.com/mitchellh/mapstructure"
)

const (
	/*ChannelStop is an iota*/
	ChannelStop = iota
	/*UserStop is an iota*/
	UserStop
	/*MessageStop is an iota*/
	MessageStop
)

func addChannel(client *Client, data interface{}) {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	go func() {
		err = r.Table("channel").
			Insert(channel).
			Exec(client.session)
		if err != nil {
			client.send <- Message{"error", err.Error()}
			return
		}
	}()
}

func subscribeChannel(client *Client, data interface{}) {
	stop := client.NewStopChannel(ChannelStop)
	result := make(chan r.ChangeResponse)
	cursor, err := r.Table("channel").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(client.session)

	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	go func() {
		var change r.ChangeResponse
		for cursor.Next(&change) {
			result <- change

		}
	}()

	go func() {
		for {
			select {
			case <-stop:
				cursor.Close()
				return
			case change := <-result:
				if change.NewValue != nil && change.OldValue == nil {
					client.send <- Message{"channel add", change.NewValue}
					fmt.Println("Sent add channel")
				}
			}
		}
	}()
}

func unsubscribeChannel(client *Client, data interface{}) {}
