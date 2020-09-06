package main

import (
	"log"
	"net/http"

	r "github.com/dancannon/gorethink"
)

/*Channel is a struct type */
type Channel struct {
	ID   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
}

/*User struct*/
type User struct {
	ID   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "172.17.0.2:28015",
		Database: "realtime_go_db",
	})

	if err != nil {
		log.Panic(err.Error())
	}

	router := NewRouter(session)

	router.Handle("channel add", addChannel)
	router.Handle("channel subscribe", subscribeChannel)

	http.Handle("/", router)
	http.ListenAndServe(":8081", nil)
}
