//data models for pingip/index page

package main

type client struct {
	Title string `json:"title"`
}

var clientList = []client{
	client{Title: "Current Client IP:"},
}

//return clientList
func getClientList() []client {
	return clientList
}
