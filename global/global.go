package global

import (
	"net/http"
	"time"
)

type contextKey string

const UserIdKey contextKey = "userID"

// ACCESS_TOKEN_EXP := time

const (
	AccessTokenExp  = 1 * time.Hour
	RefreshTokenExp = 30 * 24 * time.Hour
)

type Middleware func(http.Handler) http.Handler

func Chain(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, x := range xs {
			next = x(next)
		}
		
		return next
	}

}
