package main

import (
	"fmt"

	r "github.com/dancannon/gorethink"
)



func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "172.17.0.2:28015",
		Database: "realtime_go_db",
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	// user := User{
	// 	Name: "descholar",
	// }
	// r.Table("user").Insert(user).Exec(session) //This doesn't return an entered data
	// response, err := r.Table("user").Insert(user).RunWrite(session) //This return an entered data
	// response, err := r.Table("user").Get("701967cb-9159-4085-86a0-59b524f46166").Update(user).RunWrite(session)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%#v\n", response)
	cursor, _ := r.Table("user").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(session)

	var changeResp r.ChangeResponse
	for cursor.Next(&changeResp) {
		fmt.Printf("%#v\n", changeResp)
	}
}
