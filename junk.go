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
	// r.Table("user").Insert(user).Exec(session) //This doesn't return an entered data
	response, err := r.Table("user").Insert(user).RunWrite(session) //This return an entered data

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", response)

}
