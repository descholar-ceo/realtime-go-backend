package main

import r "github.com/dancannon/gorethink"

func subscribe(session *r.Session) {

}

func main() {
	session, _ := r.Connect(r.ConnectOpts{
		Address:  "172.17.0.2:28015",
		Database: "realtime_go_db",
	})
}
