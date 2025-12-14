package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return  func (c *gin.Context){
		id := uuid.New().String()
		c.Set("requestID",id)
		c.Header("X-Request-ID",id)
		log.Println("Request ID:",id)
		c.Next()

	}

}