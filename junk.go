package main

import (
	"fmt"

	r "github.com/dancannon/gorethink"
)

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
		fmt.Println(err)
		return
	}
	user := User{
		Name: "anonymous",
	}
}
