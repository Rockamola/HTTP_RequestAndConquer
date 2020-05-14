//handles all the routes for the templates

package main

func initalizeRoutes() {

	//index route handler
	router.GET("/", showIPIndexPage)
}
