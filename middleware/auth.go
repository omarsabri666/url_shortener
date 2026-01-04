package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/omarsabri666/url_shorter/global"
	"github.com/omarsabri666/url_shorter/helpers"
	"github.com/omarsabri666/url_shorter/model/token"
)




func AuthMiddlewareHttp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helpers.WriteJson(w, 401, token.TokenResponse{Message: "missing token", Success: false})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			helpers.WriteJson(w, 401, token.TokenResponse{Message: "Invalid Authorization header format", Success: false})
			return
		}
		tokenString := parts[1]
		accessTokenSecret := os.Getenv("ACCESS_TOKEN")
		claims := jwt.MapClaims{}
		tokenStruct, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			return []byte(accessTokenSecret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil || !tokenStruct.Valid {

			if errors.Is(err, jwt.ErrTokenExpired) {
				// c.JSON(401, gin.H{"error": "Token expired"})
				helpers.WriteJson(w, 401, token.TokenResponse{Message: "Token expired", Success: false})
				return
			}

			// c.JSON(401, gin.H{"error": "Invalid token"})
			helpers.WriteJson(w, 401, token.TokenResponse{Message: "Invalid token", Success: false})
			return
		}
		println(claims)

		// Extract userID - expiration already validated by ParseWithClaims
		userID, ok := claims["sub"].(string)
		if !ok || userID == "" {
			// c.JSON(401, gin.H{"error": "Invalid token claims"})
			helpers.WriteJson(w, 401, token.TokenResponse{Message: "Invalid token claims", Success: false})
			return
		}

		// c.Set("userID", userID)
		ctx := context.WithValue(r.Context(), global.UserIdKey, userID)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
