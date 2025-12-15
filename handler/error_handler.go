package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/omarsabri666/url_shorter/err"
)

func HandleError(c *gin.Context, e error) {
	if appErr, ok := e.(*errs.AppError); ok {
		c.JSON(appErr.Code, gin.H{
			"success": false,
			"message": appErr.Message,
			"details": appErr.Details,
		})
		return
	}
	c.JSON(500, gin.H{
		"success": false,
		"message": "internal server error",
	})

}
