//handles request and header data to get clients ip upon GET request

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIPIndexPage(c *gin.Context) {
	//fetch function for list from models file
	clientList := getClientList()
	//HTML method to render template
	c.HTML(
		//set http status as 200
		http.StatusOK,
		//template inheritance
		"index.html",
		//passed data for pages use
		gin.H{
			"title":    "InitialIpTest",
			"payload":  clientList,
			"payloads": c.ClientIP(),
		},
	)
}
