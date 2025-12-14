package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	errs "github.com/omarsabri666/url_shorter/err"
	"github.com/omarsabri666/url_shorter/handler"
	"github.com/redis/go-redis/v9"
)

func RateLimiterMiddleware(redis *redis.Client,window time.Duration,limit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
        key := fmt.Sprintf("rate:%s", ip)
		count,err := redis.Incr(c,key).Result()
		if err != nil {
		 handler.HandleError(c, errs.RateLimitExceeded("rate limiter error"))
		 log.Println(err)

			c.Abort()
		
		}
		if count ==1 {
			redis.Expire(c,key,window)
		}
		remaining  := limit - int(count)
		if remaining  < 0 {
			remaining  = 0
		
		}
		   c.Header("X-Rate-Limit-Limit", fmt.Sprintf("%d", limit))
        c.Header("X-Rate-Limit-Remaining", fmt.Sprintf("%d", remaining))
        c.Header("X-Rate-Limit-Reset", fmt.Sprintf("%d", int(window.Seconds())))


		if count > int64(limit) {
			 handler.HandleError(c, errs.RateLimitExceeded(""))

			// c.JSON(429, gin.H{"error": "rate limit exceeded"})
			c.Abort()
			return
		}


		c.Next()
	}
}