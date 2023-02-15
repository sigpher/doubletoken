package router

import (
	"doubletoken/handler"
	"doubletoken/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/auth", handler.AuthHandler)
	router.GET("/", middleware.JWTAuthMiddleware(), handler.HomeHandler)
	return router
}
