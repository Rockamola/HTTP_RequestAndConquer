//a web app for testing parameters of http requests, meant to serve as
//platform for workarounds rather than security functions. But both go hand-in-hand
//so really, it's just purely educational for myself

package main

import (
	//gin library, for framework/web server engine
	"github.com/gin-gonic/gin"
)

//global gin engine, serving as the router for web app
var router *gin.Engine

func main() {

	//use default router construct provided by gin
	router = gin.Default()

	//load templates via gin package, makes for quick & easy processing
	router.LoadHTMLGlob("templates/*")

	//temporary route function for index page
	initalizeRoutes()

	//start web app
	router.Run()
}
