package main

import (
	"github.com/gin-gonic/gin"
)

// Helper function to create a router
func getRouterWithInitialized(withTemplateLoaded bool) *gin.Engine {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	if withTemplateLoaded {
		// Process the templates at the start so that they don't have to be loaded
		// from the disk again. This makes serving HTML pages very fast.
		router.LoadHTMLGlob("templates/*")
	}
	return router
}

func main() {

	router := getRouterWithInitialized(true)

	// Handle index
	router.GET("/", showIndexPage)
	// Handle single article by its id
	router.GET("/article/view/:article_id", showArticle)

	// Start serving the application
	router.Run()
}