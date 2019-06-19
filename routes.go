// routes.go

package main

func initializRoutes() {

	// Handle the index router
	router.GET("/", showIndexPage)
}
