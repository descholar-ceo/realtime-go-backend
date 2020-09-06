package main

import r "github.com/dancannon/gorethink"

func subscribe(session *r.Session) {
	var change r.ChangeResponse
	cursor, _ := r.Table("channel").Changes().Run(session)
}

func main() {
	session, _ := r.Connect(r.ConnectOpts{
		Address:  "172.17.0.2:28015",
		Database: "realtime_go_db",
	})
}
