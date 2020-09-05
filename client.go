package main

/*Message is a struct*/
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

/*Client struct*/
type Client struct {
	send chan Message
}

func (client *Client) write() {

}

func main() {}
