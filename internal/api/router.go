package router

import (
	"github.com/gin-gonic/gin"
	"oms/internal/api/handlers"
	"oms/internal/api/middleware"
)

func SetupRoutes(r *gin.Engine, handlers *handlers.AppHandlers) {
	api := r.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/register", handlers.AuthHandler.Register)
	auth.POST("/login", handlers.AuthHandler.Login)

	protected := api.Group("", middleware.AuthMiddleware())

	user := protected.Group("/user")
	user.GET("/get", handlers.UserHandler.Get)
}
