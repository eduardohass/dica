package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize the router using gin default configurations
	router := gin.Default()
	// Setting up a route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// Running the API
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
