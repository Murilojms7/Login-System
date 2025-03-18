package router

import (
	"github.com/Murilojms7/Login-System/handler"
	"github.com/gin-gonic/gin"
)

func initializeUserRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/user"
	v1 := router.Group(basePath)
	{
		v1.GET("/", handler.ListUserHandler)
		v1.GET("/:id", handler.ShowUserHandler)
		v1.POST("/register", handler.CreateUserHandler)
		v1.POST("/login", handler.LoginUserHandler)
		v1.PUT("/", handler.UpdateUserHandler)
		v1.DELETE("/", handler.DeleteUserHandler)
	}
}
