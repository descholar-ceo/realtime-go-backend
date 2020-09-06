package main

import (
	"fmt"
	"time"

	r "github.com/dancannon/gorethink"
)

func subscribe(session *r.Session, stop <-chan bool) {
	result := make(chan r.ChangeResponse)
	cursor, _ := r.Table("channel").Changes().Run(session)

	go func() {
		var change r.ChangeResponse
		for cursor.Next(&change) {
			// fmt.Printf("%#v\n", change.NewValue)
			result <- change
		}
	}()
	for {
		select {
		case change := <-result:
			fmt.Printf("%#v\n", change.NewValue)
		case <-stop:
			cursor.Close()
			return
		}
	}
}

func main() {
	session, _ := r.Connect(r.ConnectOpts{
		Address:  "172.17.0.2:28015",
		Database: "realtime_go_db",
	})
	stop := make(chan bool)
	go subscribe(session, stop)
	time.Sleep(time.Second * 5)
	fmt.Println("Browser closes... \nWebsocket closes")
	time.Sleep(time.Second * 1000)
}
