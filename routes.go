// routes.go

package main

func initializRoutes() {

	// Handle the index router
	router.GET("/", showIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)

}
