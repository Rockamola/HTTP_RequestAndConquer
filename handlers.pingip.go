//handles request and header data to get clients ip upon GET request

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type client struct {
	Title   string "json:'title'"
	Content string "json:'content'"
}

var clientList = []client{
	client{Title: "Client's IP", Content: "content"},
}

func showIPIndexPage(c *gin.Context) {
	//HTML method to render template
	c.HTML(
		//set http status as 200
		http.StatusOK,
		//template inheritance
		"index.html",
		//passed data for pages use
		gin.H{
			"title":   "IPTest",
			"payload": clientList,
		},
	)
}

//try adding ClientIP func, calling the header info into it

func getIP(r *http.Request) *client {
	i := r.Header.Get("X-FORWARDED-FOR")

	return &client{
		Content: i,
	}
}
