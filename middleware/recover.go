package middleware

import (
	"log"
	"net/http"

	"github.com/omarsabri666/url_shorter/helpers"
	"github.com/omarsabri666/url_shorter/model/user"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// log the panic
				log.Printf("panic recovered: %v\n", rec)
				// respond with 500 Internal Server Error
				helpers.WriteJson(w, 500, user.UserResponse{Message: "interal server error", Success: false})

			}
		}()

		// call the next handler
		next.ServeHTTP(w, r)
	})
}
