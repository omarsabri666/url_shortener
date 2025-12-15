package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omarsabri666/url_shorter/handler"
	"github.com/omarsabri666/url_shorter/middleware"
)

func RegisterUserRoutes(rg *gin.RouterGroup, handler *handler.UserHandler) {
	url := rg.Group("/auth")
	{
		url.POST("/signup", handler.Signup)
		url.POST("/signin", handler.Login)
		url.POST("/signout", middleware.AuthMiddleware(), handler.Logout)
		url.POST(("/refresh"), handler.RefreshToken)
	}

}
