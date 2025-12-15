package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omarsabri666/url_shorter/handler"
	// "github.com/omarsabri666/url_shorter/middleware"
)

func RegisterURLRoutes(rg *gin.RouterGroup, handler *handler.URLHandler) {

	// Protected routes
	protected := rg.Group("/")
	// protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/", handler.CreateURL)
		protected.GET("/:short_url", handler.GetURL)
	}

	// Public routes producation only
	// rg.GET("/:short_url", handler.GetURL)
}
