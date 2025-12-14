package middleware

import (
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	errs "github.com/omarsabri666/url_shorter/err"
	"github.com/omarsabri666/url_shorter/handler"
)
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
                    handler.HandleError(c,        errs.Unauthorized("Unauthorized"))

            c.Abort()
            return
        }

        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            // c.JSON(401, gin.H{"error": "Invalid Authorization header format"})
            handler.HandleError(c,errs.Unauthorized("Invalid Authorization header format"))
            c.Abort()
            return
        }

        tokenString := parts[1]
        accessTokenSecret := os.Getenv("ACCESS_TOKEN")

        claims := jwt.MapClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
            return []byte(accessTokenSecret), nil
        }, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
    

        if err != nil || !token.Valid {

          

          if errors.Is(err, jwt.ErrTokenExpired) {
        // c.JSON(401, gin.H{"error": "Token expired"})
        handler.HandleError(c,        errs.Unauthorized("Token expired"))
        c.Abort()
        return
    }


            // c.JSON(401, gin.H{"error": "Invalid token"})
            errs.Unauthorized("Invalid token")
            c.Abort()
            return
        }
        println(claims)

        // Extract userID - expiration already validated by ParseWithClaims
        userID, ok := claims["sub"].(string)
        if !ok || userID == "" {
            // c.JSON(401, gin.H{"error": "Invalid token claims"})
                    handler.HandleError(c,        errs.Unauthorized("Invalid token claims"))

            c.Abort()
            return
        }

        c.Set("userID", userID)
        c.Next()
    }
}
// func AuthMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         authHeader := c.GetHeader("Authorization")
//         if authHeader == "" {
//             c.JSON(401, gin.H{"error": "Unauthorized"})
//             c.Abort()
//             return
//         }

//         parts := strings.SplitN(authHeader, " ", 2)
//         if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
//             c.JSON(401, gin.H{"error": "Invalid Authorization header format"})
//             c.Abort()
//             return
//         }

//         tokenString := parts[1]
//         accessTokenSecret := os.Getenv("ACCESS_TOKEN")

//         token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
//             return []byte(accessTokenSecret), nil
//         }, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

//         if err != nil || !token.Valid {
//             c.JSON(401, gin.H{"error": "Invalid or expired token"})
//             c.Abort()
//             return
//         }

//         claims, ok := token.Claims.(jwt.MapClaims)
//         if !ok {
//             c.JSON(401, gin.H{"error": "Invalid token claims"})
//             c.Abort()
//             return
//         }
// 		    if exp, ok := claims["exp"].(float64); ok {
//             if time.Unix(int64(exp), 0).Before(time.Now()) {
//                 c.JSON(401, gin.H{"error": "Token expired"})
//                 c.Abort()
//                 return
//             }
//         }

//         // Add userID to context so handlers can access it
//         if userID, ok := claims["sub"].(string); ok {
//             c.Set("userID", userID)
//         } else {
//             c.JSON(401, gin.H{"error": "Invalid token claims"})
//             c.Abort()
//             return
//         }

//         c.Next()
//     }
// }