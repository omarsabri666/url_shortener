package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/omarsabri666/url_shorter/helpers"
	"github.com/omarsabri666/url_shorter/model/user"
	"github.com/redis/go-redis/v9"
)

func RateLimiterMiddlewareHttp(
	next http.Handler,
	redis *redis.Client,
	window time.Duration,
	limit int,
) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := helpers.GetClientIP(r)
		key := fmt.Sprintf("rate:%s", ip)

		ctx := r.Context()

		count, err := redis.Incr(ctx, key).Result()
		if err != nil {
			log.Println("rate limiter redis error:", err)
			helpers.WriteJson(w, 500, user.UserResponse{
				Message: "Internal server error",
				Success: false,
			})
			return
		}

		if count == 1 {
			redis.Expire(ctx, key, window)
		}

		ttl, _ := redis.TTL(ctx, key).Result()

		remaining := limit - int(count)
		if remaining < 0 {
			remaining = 0
		}

		// ✅ Set headers BEFORE response
		w.Header().Set("X-Rate-Limit-Limit", strconv.Itoa(limit))
		w.Header().Set("X-Rate-Limit-Remaining", strconv.Itoa(remaining))
		w.Header().Set("X-Rate-Limit-Reset", strconv.Itoa(int(ttl.Seconds())))

		if count > int64(limit) {
			helpers.WriteJson(w, 429, user.UserResponse{
				Message: "Rate limit exceeded",
				Success: false,
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
