package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize the router using gin default configurations
	router := gin.Default()
	// Initialize routes
	initializeRoutes(router)
	// Running the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
