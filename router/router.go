package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	r := gin.Default()

	// Initialize Routes
	initializeUserRoutes(r)

	r.Run(":8080")

}
